package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ UserGetUsecase = (*userGetUsecase)(nil)

// UserGetUsecase は usecase.userGetUsecase のインターフェースを定義
type UserGetUsecase interface {
	Get(ctx context.Context, input UserGetUsecaseInput) (*entities.User, error)
}

// userGetUsecase ユースケース
type userGetUsecase struct {
	db       *sql.DB
	userRepo repository.UserRepository
}

// NewUserGetUsecase は userGetUsecase を初期化する
func NewUserGetUsecase(
	db *sql.DB,
	userRepo repository.UserRepository,
) UserGetUsecase {
	return &userGetUsecase{
		db:       db,
		userRepo: userRepo,
	}
}

// UserGetUsecaseInput はuserUsecase.Getのインプット
type UserGetUsecaseInput struct {
	UserID int
}

// Get はユーザー情報を取得する
func (a *userGetUsecase) Get(ctx context.Context, input UserGetUsecaseInput) (*entities.User, error) {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	userID := entities.UserID(input.UserID)

	// ユーザー情報を取得する
	user, err := a.userRepo.Find(ctx, tx, &userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
