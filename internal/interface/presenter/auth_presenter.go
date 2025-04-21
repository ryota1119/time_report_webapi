package presenter

import (
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// AuthResponse 認証APIの基本レスポンス構造体
// トークン情報を返すための共通フォーマット
type AuthResponse struct {
	AccessToken  entities.SignedToken `json:"access_token" example:"access_token"`
	RefreshToken entities.SignedToken `json:"refresh_token" example:"refresh_token"`
	ExpiresAt    time.Time            `json:"expires_at" example:"2020-09-29T23:59:59Z"`
}

// AuthLoginResponse Loginアクションのレスポンス
// AuthResponseを継承し、新規作成時に利用
type AuthLoginResponse AuthResponse

// NewAuthLoginResponse entities.AuthToken から AuthLoginResponse を生成する
// ユーザー作成時のレスポンスとして使用
func NewAuthLoginResponse(authToken *entities.AuthToken) AuthLoginResponse {
	return AuthLoginResponse{
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
		ExpiresAt:    authToken.ExpiresAt,
	}
}

// AuthRefreshTokenResponse RefreshTokenアクションのレスポンス
// AuthResponseを継承し、新規作成時に利用
type AuthRefreshTokenResponse AuthResponse

// NewAuthRefreshTokenResponse entities.AuthToken から AuthLoginResponse を生成する
// ユーザー作成時のレスポンスとして使用
func NewAuthRefreshTokenResponse(authToken *entities.AuthToken) AuthRefreshTokenResponse {
	return AuthRefreshTokenResponse{
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
		ExpiresAt:    authToken.ExpiresAt,
	}
}
