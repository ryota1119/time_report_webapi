package usecase

import (
	"context"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"
)

// AuthLogoutUsecase は authLogoutUsecase の抽象
type AuthLogoutUsecase interface {
	// Logout はログアウトを行う
	Logout(ctx context.Context) error
}

// authLogoutUsecase ユースケース
type authLogoutUsecase struct {
	authRepo repository.AuthRepository
}

// NewAuthLogoutUsecase は authLogoutUsecase を初期化する
func NewAuthLogoutUsecase(
	authRepo repository.AuthRepository,
) AuthLogoutUsecase {
	return &authLogoutUsecase{
		authRepo: authRepo,
	}
}

// Logout はログアウトを行う
func (a *authLogoutUsecase) Logout(ctx context.Context) error {
	err := a.authRepo.DeleteToken(ctx)
	if err != nil {
		return err
	}
	return nil
}
