package usecase

import (
	"context"
	"errors"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/helper/auth_context"
)

// UserUsecaseSoftDeleteInput はuserUsecase.SoftDeleteのインプット
type UserUsecaseSoftDeleteInput struct {
	UserID int
}

// SoftDelete はユーザー情報を論理削除する
func (a *userUsecase) SoftDelete(ctx context.Context, input UserUsecaseSoftDeleteInput) error {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	userID := entities.UserID(input.UserID)
	// 自身のユーザー情報は削除できない
	if auth_context.ContextUserID(ctx) == userID {
		return errors.New("can't be delete myself")
	}

	// 既存のユーザー情報を取得する
	_, err = a.userRepo.Find(ctx, tx, &userID)
	if err != nil {
		return err
	}

	// ユーザー情報を削除する
	err = a.userRepo.SoftDelete(ctx, tx, &userID)
	if err != nil {
		return err
	}
	return nil
}
