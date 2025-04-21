package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"
	"github.com/ryota1119/time_resport_webapi/internal/domain/service"
)

var _ AuthLoginUsecase = (*authLoginUsecase)(nil)

// AuthLoginUsecase は authLoginUsecase の抽象
type AuthLoginUsecase interface {
	// Login はログインを行う
	Login(ctx context.Context, input AuthUsecaseLoginInput) (*entities.AuthToken, error)
}

// authLoginUsecase ユースケース
type authLoginUsecase struct {
	db               *sql.DB
	jwtTokenService  service.JwtTokenService
	authRepo         repository.AuthRepository
	organizationRepo repository.OrganizationRepository
	userRepo         repository.UserRepository
}

// NewAuthLoginUsecase は authLoginUsecase を初期化する
func NewAuthLoginUsecase(
	db *sql.DB,
	jwtTokenService service.JwtTokenService,
	authRepo repository.AuthRepository,
	organizationRepo repository.OrganizationRepository,
	userRepo repository.UserRepository,
) AuthLoginUsecase {
	return &authLoginUsecase{
		db:               db,
		jwtTokenService:  jwtTokenService,
		authRepo:         authRepo,
		organizationRepo: organizationRepo,
		userRepo:         userRepo,
	}
}

// AuthUsecaseLoginInput は authUsecase.Login のインプット
type AuthUsecaseLoginInput struct {
	OrganizationCode string
	Email            string
	Password         string
}

// Login ログイン
func (a *authLoginUsecase) Login(ctx context.Context, input AuthUsecaseLoginInput) (*entities.AuthToken, error) {
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

	organizationCode := entities.OrganizationCode(input.OrganizationCode)
	organization, err := a.organizationRepo.FindByCode(ctx, tx, &organizationCode)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "organization_id", organization.ID)

	userEmail := entities.UserEmail(input.Email)
	// メールアドレスからユーザー情報を取得
	user, err := a.userRepo.FindByEmail(ctx, tx, &userEmail, &organization.ID)
	if err != nil {
		return nil, err
	}

	// パスワードの検証
	if err = user.HashedPassword.CheckHashedPassword(input.Password); err != nil {
		return nil, entities.ErrPasswordNotMatch
	}

	// アクセストークン生成
	accessToken, accessJti, err := a.jwtTokenService.GenerateJwtToken(user, organization, 15*time.Minute)
	if err != nil {
		return nil, err
	}
	// リフレッシュトークン生成
	refreshToken, refreshJti, err := a.jwtTokenService.GenerateJwtToken(user, organization, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	// アクセストークンをredisに保存
	err = a.authRepo.SaveAccessToken(ctx, user.ID, accessJti, 15*time.Hour)
	if err != nil {
		return nil, err
	}
	// リフレッシュトークンをredisに保存
	err = a.authRepo.SaveRefreshToken(ctx, user.ID, refreshJti, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return entities.NewAuthToken(
		*accessToken,
		*refreshToken,
	), nil
}
