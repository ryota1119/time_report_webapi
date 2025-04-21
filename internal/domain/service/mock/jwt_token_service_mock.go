package mock

import (
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

type JwtTokenServiceMock struct {
	GenerateJwtTokenFunc func(user *entities.User, organizationID *entities.OrganizationID, expires time.Duration) (*string, *entities.Jti, error)
	ValidateJwtTokenFunc func(token string) (*entities.CustomClaims, error)
}

func (m *JwtTokenServiceMock) GenerateJwtToken(user *entities.User, organizationID *entities.OrganizationID, expires time.Duration) (*string, *entities.Jti, error) {
	if m.GenerateJwtTokenFunc != nil {
		return m.GenerateJwtTokenFunc(user, organizationID, expires)
	}
	return nil, nil, nil
}

func (m *JwtTokenServiceMock) ValidateJwtToken(token string) (*entities.CustomClaims, error) {
	if m.ValidateJwtTokenFunc != nil {
		return m.ValidateJwtTokenFunc(token)
	}
	return nil, nil
}
