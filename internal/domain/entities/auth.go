package entities

import "time"

// SignedToken 署名トークン
type SignedToken string

// String は SignedToken を string 型にキャストする
func (t SignedToken) String() string {
	return string(t)
}

// Jti jtiトークン
type Jti string

// String は Jti を string 型にキャストする
func (j Jti) String() string {
	return string(j)
}

// ExpiresAt トークン有効期限
type ExpiresAt time.Time

// AuthToken ユーザーに発行された認証トークン
type AuthToken struct {
	AccessToken  SignedToken `json:"access_token"`
	RefreshToken SignedToken `json:"refresh_token"`
	ExpiresAt    time.Time   `json:"expires_at"`
}

// NewAuthToken は指定されたアクセストークンとリフレッシュトークン文字列から AuthToken を生成する
func NewAuthToken(accessToken, refreshToken string) *AuthToken {
	return &AuthToken{
		AccessToken:  SignedToken(accessToken),
		RefreshToken: SignedToken(refreshToken),
		ExpiresAt:    time.Now().Add(15 * time.Minute),
	}
}
