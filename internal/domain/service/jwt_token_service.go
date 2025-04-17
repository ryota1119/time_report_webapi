package service

import (
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

type JwtTokenService interface {
	GenerateJwtToken(user *entities.User, organization *entities.Organization, expires time.Duration) (*string, *entities.Jti, error)
	ValidateJwtToken(token string) (*entities.CustomClaims, error)
}
