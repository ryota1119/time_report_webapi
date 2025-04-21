package jwt_token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// ValidateJwtToken はトークンの検証
func (j *jwtToken) ValidateJwtToken(token string) (*entities.CustomClaims, error) {
	claims := &entities.CustomClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil || !tkn.Valid {
		return nil, entities.ErrInvalidToken
	}
	return claims, nil
}
