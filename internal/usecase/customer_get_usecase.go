package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ CustomerGetUsecase = (*customerGetUsecase)(nil)

// CustomerGetUsecase は usecase.customerGetUsecase のインターフェースを定義
type CustomerGetUsecase interface {
	// Get は顧客情報を更新する
	Get(ctx context.Context, input CustomerGetUsecaseInput) (*entities.Customer, error)
}

// customerGetUsecase ユースケース
type customerGetUsecase struct {
	db           *sql.DB
	customerRepo repository.CustomerRepository
}

func NewCustomerGetUsecase(
	db *sql.DB,
	customerRepo repository.CustomerRepository,
) CustomerGetUsecase {
	return &customerGetUsecase{
		db:           db,
		customerRepo: customerRepo,
	}
}

// CustomerGetUsecaseInput は customerUsecase.Get のinput
type CustomerGetUsecaseInput struct {
	CustomerID uint
}

// Get は顧客情報を取得する
func (a *customerGetUsecase) Get(ctx context.Context, input CustomerGetUsecaseInput) (*entities.Customer, error) {
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
