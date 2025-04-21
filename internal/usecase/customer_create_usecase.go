package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ CustomerCreateUsecase = (*customerCreateUsecase)(nil)

// CustomerCreateUsecase は usecase.customerCreateUsecase のインターフェースを定義
type CustomerCreateUsecase interface {
	// Create は顧客情報を新規登録する
	Create(ctx context.Context, input CustomerCreateUsecaseInput) (*entities.Customer, error)
}

// customerCreateUsecase ユースケース
type customerCreateUsecase struct {
	db           *sql.DB
	customerRepo repository.CustomerRepository
}

func NewCustomerCreateUsecase(
	db *sql.DB,
	customerRepo repository.CustomerRepository,
) CustomerCreateUsecase {
	return &customerCreateUsecase{
		db:           db,
		customerRepo: customerRepo,
	}
}

// CustomerCreateUsecaseInput は customerUsecase.Create のinput
type CustomerCreateUsecaseInput struct {
	Name      string
	UnitPrice *int64
	StartDate *string
	EndDate   *string
}

// Create は顧客を新規作成する
func (a *customerCreateUsecase) Create(ctx context.Context, input CustomerCreateUsecaseInput) (*entities.Customer, error) {
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

	customer, err := entities.NewCustomer(input.Name, input.UnitPrice, input.StartDate, input.EndDate)
	customerID, err := a.customerRepo.Create(ctx, tx, customer)
	if err != nil {
		return nil, err
	}

	customer.ID = *customerID

	return customer, nil
}
