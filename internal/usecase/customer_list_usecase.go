package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// List は顧客の一覧を取得する
func (a *customerUsecase) List(ctx context.Context) ([]entities.Customer, error) {
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

	// ログインユーザー組織IDを取得する
	customers, err := a.customerRepo.List(ctx, tx)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
