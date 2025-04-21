package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ UserListUsecase = (*userListUsecase)(nil)

// UserListUsecase は usecase.userListUsecase のインターフェースを定義
type UserListUsecase interface {
	List(ctx context.Context) ([]entities.User, error)
}

// userListUsecase ユースケース
type userListUsecase struct {
	db       *sql.DB
	userRepo repository.UserRepository
}

// NewUserListUsecase は userListUsecase を初期化する
func NewUserListUsecase(
	db *sql.DB,
	userRepo repository.UserRepository,
) UserListUsecase {
	return &userListUsecase{
		db:       db,
		userRepo: userRepo,
	}
}

// List は組織のユーザーの一覧を取得する
func (a *userListUsecase) List(ctx context.Context) ([]entities.User, error) {
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

	// ユーザー一覧を取得する
	users, err := a.userRepo.List(ctx, tx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
