package usecase

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"
	"github.com/ryota1119/time_resport_webapi/internal/domain/service"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ AuthRefreshTokenUsecase = (*authRefreshTokenUsecase)(nil)

// AuthRefreshTokenUsecase は authRefreshTokenUsecase の抽象
type AuthRefreshTokenUsecase interface {
	// RefreshToken はログインを行う
	RefreshToken(ctx context.Context, input AuthUsecaseRefreshTokenInput) (*entities.AuthToken, error)
}

// authRefreshTokenUsecase ユースケース
type authRefreshTokenUsecase struct {
	db               *sql.DB
	jwtTokenService  service.JwtTokenService
	authRepo         repository.AuthRepository
	organizationRepo repository.OrganizationRepository
	userRepo         repository.UserRepository
}

// NewAuthRefreshTokenUsecase は authRefreshTokenUsecase を初期化する
func NewAuthRefreshTokenUsecase(
	db *sql.DB,
	jwtTokenService service.JwtTokenService,
	authRepo repository.AuthRepository,
	organizationRepo repository.OrganizationRepository,
	userRepo repository.UserRepository,
) AuthRefreshTokenUsecase {
	return &authRefreshTokenUsecase{
		db:               db,
		jwtTokenService:  jwtTokenService,
		authRepo:         authRepo,
		organizationRepo: organizationRepo,
		userRepo:         userRepo,
	}
}

// AuthUsecaseRefreshTokenInput は authUsecase.RefreshToken のインプット
type AuthUsecaseRefreshTokenInput struct {
	RefreshToken string
}

// RefreshToken は新しい accessToken を発行する
func (a *authRefreshTokenUsecase) RefreshToken(ctx context.Context, input AuthUsecaseRefreshTokenInput) (*entities.AuthToken, error) {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// JWT トークンの検証
	claims, err := a.jwtTokenService.ValidateJwtToken(input.RefreshToken)
	if err != nil {
		return nil, err
	}

	organizationCode := entities.OrganizationCode(claims.OrganizationCode)
	organization, err := a.organizationRepo.FindByCode(ctx, tx, &organizationCode)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, "organization_id", organization.ID)

	// jwtからユーザーIDを取得
	jwtUserID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return nil, err
	}
	userID := entities.UserID(jwtUserID)

	// jti 取得
	jti := entities.Jti(claims.ID)
	// Redis でトークンが有効か確認
	redisUserID, err := a.authRepo.GetUserIDByRefreshJti(ctx, &jti)
	if errors.Is(err, redis.Nil) {
		return nil, entities.ErrUserNotFoundInRedis
	} else if err != nil {
		return nil, err
	}

	// リクエストで受け取ったユーザーIDとRedisに保存してあるユーザーIDの比較
	if userID != *redisUserID {
		return nil, entities.ErrUnauthorized
	}

	// データベースにユーザーが存在するか確認
	user, err := a.userRepo.Find(ctx, tx, &userID)
	if err != nil {
		return nil, err
	}
	org, err := a.organizationRepo.FindByCode(ctx, tx, &organizationCode)
	if err != nil {
		return nil, err
	}

	// アクセストークン生成
	newAccessToken, accessJti, err := a.jwtTokenService.GenerateJwtToken(user, org, 15*time.Minute)
	if err != nil {
		return nil, err
	}
	// アクセストークンをredisに保存
	err = a.authRepo.SaveAccessToken(ctx, userID, accessJti, 15*time.Hour)
	if err != nil {
		return nil, err
	}

	// リフレッシュトークン生成
	newRefreshToken, refreshJti, err := a.jwtTokenService.GenerateJwtToken(user, org, 24*time.Hour)
	if err != nil {
		return nil, err
	}
	// リフレッシュトークンをredisに保存
	err = a.authRepo.SaveRefreshToken(ctx, userID, refreshJti, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return entities.NewAuthToken(
		*newAccessToken,
		*newRefreshToken,
	), nil
}
