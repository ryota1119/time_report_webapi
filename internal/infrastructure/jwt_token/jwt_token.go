package jwt_token

import (
	"github.com/ryota1119/time_resport/internal/domain/service"
	_ "github.com/ryota1119/time_resport/internal/helper/auth_context"
)

type jwtToken struct {
	secretKey []byte
}

var _ service.JwtTokenService = (*jwtToken)(nil)

func NewJwtToken(secretKey []byte) service.JwtTokenService {
	return &jwtToken{secretKey: secretKey}
}
