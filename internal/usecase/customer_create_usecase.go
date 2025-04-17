package usecase

import (
	"context"
	"errors"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/helper/datetime"
)

// CreateCustomerUsecaseInput は customerUsecase.Create のinput
type CreateCustomerUsecaseInput struct {
	Name      string
	UnitPrice *int64
	StartDate *string
	EndDate   *string
}

// Create は顧客を新規作成する
func (a *customerUsecase) Create(ctx context.Context, input CreateCustomerUsecaseInput) (*entities.Customer, error) {
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

	// StartDateとEndDateがnilでない場合、StartDateはEndDateより前である必要がある
	// 開始日・終了日をパース
	startDate, endDate, err := datetime.ParseStartEndDate(input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}
	if startDate != nil && endDate != nil {
		if !startDate.Before(*endDate) {
			return nil, errors.New("StartDate must be before EndDate")
		}
	}

	customer := entities.NewCustomer(input.Name, input.UnitPrice, startDate, endDate)
	customerID, err := a.customerRepo.Create(ctx, tx, customer)
	if err != nil {
		return nil, err
	}

	customer.ID = *customerID

	return customer, nil
}
