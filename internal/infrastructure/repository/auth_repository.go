package repository

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/redis/go-redis/v9"
)

const MaxTokens = 5

type AuthRepository struct {
	redisClient *redis.Client
}

func NewAuthRepository(redisClient *redis.Client) repository.AuthRepository {
	return &AuthRepository{redisClient}
}

func (r *AuthRepository) SaveAccessToken(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error {
	accessKey := "login_users:access:" + jti.String()
	userAccessTokensKey := "tokens:user_access_tokens:" + userID.String()

	err := r.redisClient.Set(ctx, accessKey, userID.String(), duration).Err()
	if err != nil {
		return err
	}

	err = r.redisClient.LPush(ctx, userAccessTokensKey, accessKey).Err()
	if err != nil {
		return err
	}

	count, err := r.redisClient.LLen(ctx, userAccessTokensKey).Result()
	if err != nil {
		return err
	}
	if count > MaxTokens {
		oldestJTI, err := r.redisClient.RPop(ctx, userAccessTokensKey).Result()
		if err != nil {
			return err
		}

		err = r.redisClient.Del(ctx, oldestJTI).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *AuthRepository) SaveRefreshToken(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error {
	refreshKey := "login_users:refresh:" + jti.String()
	userRefreshTokensKey := "tokens:user_refresh_tokens:" + userID.String()

	err := r.redisClient.Set(ctx, refreshKey, userID.String(), duration).Err()
	if err != nil {
		return err
	}

	err = r.redisClient.LPush(ctx, userRefreshTokensKey, refreshKey).Err()
	if err != nil {
		return err
	}

	count, err := r.redisClient.LLen(ctx, userRefreshTokensKey).Result()
	if err != nil {
		return err
	}
	if count > MaxTokens {
		oldestJTI, err := r.redisClient.RPop(ctx, userRefreshTokensKey).Result()
		if err != nil {
			return err
		}

		err = r.redisClient.Del(ctx, oldestJTI).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *AuthRepository) GetUserIDByAccessJti(ctx context.Context, jti *entities.Jti) (*entities.UserID, error) {
	userID, err := r.redisClient.Get(ctx, "login_users:access:"+jti.String()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, entities.ErrTokenNotFoundInRedis
		}
		return nil, err
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}

	returnUserID := entities.UserID(intUserID)

	return &returnUserID, nil
}

func (r *AuthRepository) GetUserIDByRefreshJti(ctx context.Context, jti *entities.Jti) (*entities.UserID, error) {
	userID, err := r.redisClient.Get(ctx, "login_users:refresh:"+jti.String()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, entities.ErrTokenNotFoundInRedis
		}
		return nil, err
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}

	returnUserID := entities.UserID(intUserID)

	return &returnUserID, nil
}

func (r *AuthRepository) DeleteToken(ctx context.Context) error {
	userID := auth_context.ContextOrganizationID(ctx)

	accessTokensKey := "tokens:user_access_tokens:" + userID.String()
	accessTokens, err := r.redisClient.LRange(ctx, accessTokensKey, 0, -1).Result()
	if err != nil {
		return err
	}
	for _, accessToken := range accessTokens {
		err = r.redisClient.Del(ctx, accessToken).Err()
		if err != nil {
			return err
		}
	}
	err = r.redisClient.Del(ctx, accessTokensKey).Err()
	if err != nil {
		return err
	}

	secretTokensKey := "tokens:user_refresh_tokens:" + userID.String()
	secretTokens, err := r.redisClient.LRange(ctx, secretTokensKey, 0, -1).Result()
	if err != nil {
		return err
	}
	for _, accessToken := range secretTokens {
		err := r.redisClient.Del(ctx, accessToken).Err()
		if err != nil {
			return err
		}
	}
	err = r.redisClient.Del(ctx, secretTokensKey).Err()
	if err != nil {
		return err
	}

	return nil
}
