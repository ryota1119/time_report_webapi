package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/domain/errors"
	"github.com/ryota1119/time_resport/internal/helper/datetime"
)

// UpdateCustomerUsecaseInput は customerUsecase.Updateのinput
type UpdateCustomerUsecaseInput struct {
	CustomerID uint
	Name       string
	UnitPrice  *int64
	StartDate  *string
	EndDate    *string
}

// Update は顧客情報を更新する
func (a *customerUsecase) Update(ctx context.Context, input UpdateCustomerUsecaseInput) (*entities.Customer, error) {
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

	// 既存の顧客情報を取得
	customerID := entities.CustomerID(input.CustomerID)
	customer, err := a.customerRepo.Find(ctx, tx, &customerID)
	if err != nil {
		return nil, errors.ErrCustomerNotFound
	}

	// StartDateとEndDateがnilでない場合、StartDateはEndDateより前である必要がある
	// 開始日・終了日をパース
	startDate, endDate, err := datetime.ParseStartEndDate(input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}
	if startDate != nil && endDate != nil {
		if !startDate.Before(*endDate) {
			return nil, errors.ErrStartDateMustBeBefore
		}
	}

	// 何も更新がない場合は、エラーを返却し、handler層でno contentを返す
	if customer.Name.String() == input.Name &&
		customer.UnitPrice.Int64() == input.UnitPrice &&
		customer.StartDate != startDate &&
		customer.EndDate != endDate {
		return nil, errors.ErrNoContentUpdated
	}

	newCustomer := entities.NewCustomer(input.Name, input.UnitPrice, startDate, endDate)
	newCustomer.ID = customerID

	// 顧客情報を更新する
	err = a.customerRepo.Update(ctx, tx, newCustomer)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}
