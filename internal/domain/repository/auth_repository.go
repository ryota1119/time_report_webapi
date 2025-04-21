package repository

import (
	"context"
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// AuthRepository AuthRepositoryのインターフェースを定義
type AuthRepository interface {
	// SaveAccessToken はアクセストークンを保存する
	SaveAccessToken(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error
	// SaveRefreshToken はリフレッシュトークンを保存する
	SaveRefreshToken(ctx context.Context, userID entities.UserID, jti *entities.Jti, duration time.Duration) error
	// GetUserIDByAccessJti はアクセストークンから取得したjtiを使用してユーザーIDを取得する
	GetUserIDByAccessJti(ctx context.Context, jti *entities.Jti) (*entities.UserID, error)
	// GetUserIDByRefreshJti はリフレッシュトークンをから取得したjtiを使用してユーザーIDを取得する
	GetUserIDByRefreshJti(ctx context.Context, jti *entities.Jti) (*entities.UserID, error)
	// DeleteToken はトークンを削除する
	DeleteToken(ctx context.Context) error
}
