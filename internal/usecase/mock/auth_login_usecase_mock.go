package mock

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/usecase"
)

var _ usecase.AuthLoginUsecase = (*AuthLoginUsecaseMock)(nil)

type AuthLoginUsecaseMock struct {
	LoginFunc func(ctx context.Context, input usecase.AuthUsecaseLoginInput) (*entities.AuthToken, error)
}

func (m *AuthLoginUsecaseMock) Login(ctx context.Context, input usecase.AuthUsecaseLoginInput) (*entities.AuthToken, error) {
	if m.LoginFunc != nil {
		return m.LoginFunc(ctx, input)
	}
	return nil, nil
}
