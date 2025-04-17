package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// GetCustomerUsecaseInput は customerUsecase.Get のinput
type GetCustomerUsecaseInput struct {
	CustomerID uint
}

// Get は顧客情報を取得する
func (a *customerUsecase) Get(ctx context.Context, input GetCustomerUsecaseInput) (*entities.Customer, error) {
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

	customerID := entities.CustomerID(input.CustomerID)
	customer, err := a.customerRepo.Find(ctx, tx, &customerID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
