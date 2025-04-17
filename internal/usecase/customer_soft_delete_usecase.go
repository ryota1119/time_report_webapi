package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// SoftDeleteCustomerUsecaseInput は customerUsecase.Create のinput
type SoftDeleteCustomerUsecaseInput struct {
	CustomerID uint
}

// SoftDelete は顧客情報を論理削除する
func (a *customerUsecase) SoftDelete(ctx context.Context, input SoftDeleteCustomerUsecaseInput) error {
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

	// 既存の顧客情報を取得する
	customerID := entities.CustomerID(input.CustomerID)
	_, err = a.customerRepo.Find(ctx, tx, &customerID)
	if err != nil {
		return err
	}

	// 顧客情報を削除する
	err = a.customerRepo.SoftDelete(ctx, tx, &customerID)
	if err != nil {
		return err
	}
	return nil
}
