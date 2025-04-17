package mock

import (
	"context"
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// AuthRepositoryMock は AuthRepository のモック構造体
type AuthRepositoryMock struct {
	// SaveAccessTokenFunc はアクセストークンを保存するモック
	SaveAccessTokenFunc func(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error
	// SaveRefreshTokenFunc はリフレッシュトークンを保存するモック
	SaveRefreshTokenFunc func(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error
	// GetUserIDByAccessJtiFunc はアクセストークンから取得したjtiを使用してユーザーIDを取得するモック
	GetUserIDByAccessJtiFunc func(ctx context.Context, jti *entities.Jti) (*entities.UserID, error)
	// GetUserIDByRefreshTokenFunc はリフレッシュトークンをから取得したjtiを使用してユーザーIDを取得するモック
	GetUserIDByRefreshTokenFunc func(ctx context.Context, jti *entities.Jti) (*entities.UserID, error)
	// DeleteTokenFunc はトークンを削除するモック
	DeleteTokenFunc func(ctx context.Context) error
}

func (m *AuthRepositoryMock) SaveAccessToken(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error {
	if m.SaveAccessTokenFunc != nil {
		return m.SaveAccessTokenFunc(ctx, userID, jti, duration)
	}
	return nil
}

func (m *AuthRepositoryMock) SaveRefreshToken(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error {
	if m.SaveRefreshTokenFunc != nil {
		return m.SaveRefreshTokenFunc(ctx, userID, jti, duration)
	}
	return nil
}

func (m *AuthRepositoryMock) GetUserIDByAccessJti(ctx context.Context, jti *entities.Jti) (*entities.UserID, error) {
	if m.GetUserIDByAccessJtiFunc != nil {
		return m.GetUserIDByAccessJtiFunc(ctx, jti)
	}
	return nil, nil
}

func (m *AuthRepositoryMock) GetUserIDByRefreshToken(ctx context.Context, jti *entities.Jti) (*entities.UserID, error) {
	if m.GetUserIDByRefreshTokenFunc != nil {
		return m.GetUserIDByRefreshTokenFunc(ctx, jti)
	}
	return nil, nil
}

func (m *AuthRepositoryMock) DeleteToken(ctx context.Context) error {
	if m.DeleteTokenFunc != nil {
		return m.DeleteTokenFunc(ctx)
	}
	return nil
}
