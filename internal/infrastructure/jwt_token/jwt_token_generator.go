package jwt_token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// GenerateJwtToken はしてされたユーザー情報と組織情報、有効きかんからトークン文字列を生成して返す
func (j *jwtToken) GenerateJwtToken(user *entities.User, organization *entities.Organization, expires time.Duration) (*string, *entities.Jti, error) {
	jti := entities.Jti(uuid.New().String())
	expirationTime := time.Now().Add(expires)

	claims := entities.CustomClaims{
		Role:             user.Role.String(),
		OrganizationCode: organization.Code.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti.String(),
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "gin_webapi",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(j.secretKey)
	if err != nil {
		return nil, nil, entities.ErrCouldNotGenerateToken
	}

	return &signedToken, &jti, nil
}
