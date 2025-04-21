package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ CustomerUpdateUsecase = (*customerUpdateUsecase)(nil)

// CustomerUpdateUsecase は usecase.customerUpdateUsecase のインターフェースを定義
type CustomerUpdateUsecase interface {
	// Update は顧客情報を更新する
	Update(ctx context.Context, input CustomerUpdateUsecaseInput) (*entities.Customer, error)
}

// customerUpdateUsecase ユースケース
type customerUpdateUsecase struct {
	db           *sql.DB
	customerRepo repository.CustomerRepository
}

func NewCustomerUpdateUsecase(
	db *sql.DB,
	customerRepo repository.CustomerRepository,
) CustomerUpdateUsecase {
	return &customerUpdateUsecase{
		db:           db,
		customerRepo: customerRepo,
	}
}

// CustomerUpdateUsecaseInput は customerUsecase.Updateのinput
type CustomerUpdateUsecaseInput struct {
	CustomerID uint
	Name       string
	UnitPrice  *int64
	StartDate  *string
	EndDate    *string
}

// Update は顧客情報を更新する
func (a *customerUpdateUsecase) Update(ctx context.Context, input CustomerUpdateUsecaseInput) (*entities.Customer, error) {
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
		return nil, entities.ErrCustomerNotFound
	}

	customerPeriod, err := entities.NewCustomerPeriod(input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}

	isUpdated := false

	if customer.Name != entities.CustomerName(input.Name) {
		customer.Name = entities.CustomerName(input.Name)
		isUpdated = true
	}
	if customer.UnitPrice != entities.NewCustomerUnitPrice(input.UnitPrice) {
		customer.UnitPrice = entities.NewCustomerUnitPrice(input.UnitPrice)
		isUpdated = true
	}
	if customer.Period.Start == nil && customerPeriod.Start != nil ||
		customer.Period.Start != nil && customerPeriod.Start == nil ||
		(customer.Period.Start != nil && customerPeriod.Start != nil && !customer.Period.Start.Equal(*customerPeriod.Start)) {
		customer.Period.Start = customerPeriod.Start
		isUpdated = true
	}
	if customer.Period.End == nil && customerPeriod.End != nil ||
		customer.Period.End != nil && customerPeriod.End == nil ||
		(customer.Period.End != nil && customerPeriod.End != nil && !customer.Period.End.Equal(*customerPeriod.End)) {
		customer.Period.End = customerPeriod.End
		isUpdated = true
	}

	if !isUpdated {
		return nil, entities.ErrNoContentUpdated
	}

	if err := a.customerRepo.Update(ctx, tx, customer); err != nil {
		return nil, err
	}

	return customer, nil
}
