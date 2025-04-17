package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// List は組織のユーザーの一覧を取得する
func (a *userUsecase) List(ctx context.Context) ([]entities.User, error) {
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
