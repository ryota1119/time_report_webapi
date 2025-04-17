package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// UserUsecaseCreateInput userUsecase.Createのインプット
type UserUsecaseCreateInput struct {
	Name     string
	Email    string
	Password string
	Role     string
}

// Create はユーザーを新規作成する
func (a *userUsecase) Create(ctx context.Context, input UserUsecaseCreateInput) (_ *entities.User, err error) {
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
