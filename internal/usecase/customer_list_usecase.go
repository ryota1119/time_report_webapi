package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ CustomerListUsecase = (*customerListUsecase)(nil)

// CustomerListUsecase は usecase.customerListUsecase のインターフェースを定義
type CustomerListUsecase interface {
	// List は顧客情報を更新する
	List(ctx context.Context) ([]entities.Customer, error)
}

// customerListUsecase ユースケース
type customerListUsecase struct {
	db           *sql.DB
	customerRepo repository.CustomerRepository
}

func NewCustomerListUsecase(
	db *sql.DB,
	customerRepo repository.CustomerRepository,
) CustomerListUsecase {
	return &customerListUsecase{
		db:           db,
		customerRepo: customerRepo,
	}
}

// List は顧客の一覧を取得する
func (a *customerListUsecase) List(ctx context.Context) ([]entities.Customer, error) {
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
