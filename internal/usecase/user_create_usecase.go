package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ UserCreateUsecase = (*userCreateUsecase)(nil)

// UserCreateUsecase は usecase.userCreateUsecase のインターフェースを定義
type UserCreateUsecase interface {
	Create(ctx context.Context, input UserCreateUsecaseInput) (_ *entities.User, err error)
}

// userCreateUsecase ユースケース
type userCreateUsecase struct {
	db       *sql.DB
	userRepo repository.UserRepository
}

// NewUserCreateUsecase は userCreateUsecase を初期化する
func NewUserCreateUsecase(
	db *sql.DB,
	userRepo repository.UserRepository,
) UserCreateUsecase {
	return &userCreateUsecase{
		db:       db,
		userRepo: userRepo,
	}
}

// UserCreateUsecaseInput userCreateUsecase.Createのインプット
type UserCreateUsecaseInput struct {
	Name     string
	Email    string
	Password string
	Role     string
}

// Create はユーザーを新規作成する
func (a *userCreateUsecase) Create(ctx context.Context, input UserCreateUsecaseInput) (_ *entities.User, err error) {
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

	// ユーザーのドメインモデルを作成
	user, err := entities.NewUser(
		input.Name,
		input.Email,
		input.Password,
		input.Role,
	)
	if err != nil {
		return nil, err
	}

	// ユーザー情報を作成する
	userID, err := a.userRepo.Create(ctx, tx, user)
	if err != nil {
		return nil, err
	}
	user.ID = *userID

	return user, nil
}
