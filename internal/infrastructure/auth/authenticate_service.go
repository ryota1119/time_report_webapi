package auth

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// AuthenticateFromToken はトークンを受け取って、ユーザーと組織情報を返す
func (a *authService) AuthenticateFromToken(ctx context.Context, accessToken string) (*entities.User, *entities.Organization, error) {
	// JWT トークンの検証
	claims, err := a.jwtTokenService.ValidateJwtToken(accessToken)
	if err != nil {
		return nil, nil, entities.ErrInvalidOrExpiredToken
	}

	// jwtから組織コードを取得
	orgCode := entities.OrganizationCode(claims.OrganizationCode)

	// jwtからユーザーIDを取得
	jwtUserID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return nil, nil, err
	}
	userID := entities.UserID(jwtUserID)

	// jti取得
	jti := entities.Jti(claims.ID)
	// Redis でトークンが有効か確認
	redisUserID, err := a.authRepo.GetUserIDByAccessJti(ctx, &jti)
	if errors.Is(err, redis.Nil) {
		return nil, nil, entities.ErrUserNotFoundInRedis
	} else if err != nil {
		return nil, nil, err
	}

	// jwt_token tokenのクレームユーザーIDとRedisのユーザーIDが不一致の場合、エラーを返す
	if userID != *redisUserID {
		return nil, nil, entities.ErrUnauthorized
	}

	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var o entities.CachedOrganization
	authOrgKey := "auth:organization:" + orgCode.String()
	cachedOrg, err := a.redisClient.Get(ctx, authOrgKey).Result()
	if errors.Is(err, redis.Nil) {
		org, err := a.organizationRepo.FindByCode(ctx, tx, &orgCode)
		if err != nil {
			return nil, nil, err
		}
		o = entities.CachedOrganization{
			ID:   org.ID,
			Name: org.Name,
			Code: orgCode,
		}
		cachedOrgJson, err := json.Marshal(o)
		if err != nil {
			return nil, nil, err
		}
		err = a.redisClient.Set(ctx, authOrgKey, cachedOrgJson, 15*time.Hour).Err()
		if err != nil {
			return nil, nil, err
		}
	} else {
		err = json.Unmarshal([]byte(cachedOrg), &o)
		if err != nil {
			return nil, nil, err
		}
	}
	org := entities.Organization{
		ID:   o.ID,
		Name: o.Name,
		Code: o.Code,
	}

	var u entities.CachedUser
	authUserKey := "auth:user:" + userID.String()
	cachedUser, err := a.redisClient.Get(ctx, authUserKey).Result()
	if errors.Is(err, redis.Nil) {
		user, err := a.userRepo.FindWithOrganizationID(ctx, tx, &userID, &org.ID)
		if err != nil {
			return nil, nil, err
		}
		u = entities.CachedUser{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		}
		cachedUserJson, err := json.Marshal(u)
		if err != nil {
			return nil, nil, err
		}
		err = a.redisClient.Set(ctx, authUserKey, cachedUserJson, 15*time.Hour).Err()
		if err != nil {
			return nil, nil, err
		}
	} else {
		err = json.Unmarshal([]byte(cachedUser), &u)
		if err != nil {
			return nil, nil, err
		}
	}
	user := entities.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}

	return &user, &org, nil
}
