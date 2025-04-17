package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// UserUsecaseGetInput はuserUsecase.Getのインプット
type UserUsecaseGetInput struct {
	UserID int
}

// Get はユーザー情報を取得する
func (a *userUsecase) Get(ctx context.Context, input UserUsecaseGetInput) (*entities.User, error) {
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
