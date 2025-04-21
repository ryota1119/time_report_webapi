package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ CustomerSoftDeleteUsecase = (*customerSoftDeleteUsecase)(nil)

// CustomerSoftDeleteUsecase は usecase.customerSoftDeleteUsecase のインターフェースを定義
type CustomerSoftDeleteUsecase interface {
	// SoftDelete は顧客情報を更新する
	SoftDelete(ctx context.Context, input CustomerSoftDeleteUsecaseInput) error
}

// customerSoftDeleteUsecase ユースケース
type customerSoftDeleteUsecase struct {
	db           *sql.DB
	customerRepo repository.CustomerRepository
}

func NewCustomerSoftDeleteUsecase(
	db *sql.DB,
	customerRepo repository.CustomerRepository,
) CustomerSoftDeleteUsecase {
	return &customerSoftDeleteUsecase{
		db:           db,
		customerRepo: customerRepo,
	}
}

// CustomerSoftDeleteUsecaseInput は customerUsecase.Create のinput
type CustomerSoftDeleteUsecaseInput struct {
	CustomerID uint
}

// SoftDelete は顧客情報を論理削除する
func (a *customerSoftDeleteUsecase) SoftDelete(ctx context.Context, input CustomerSoftDeleteUsecaseInput) error {
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
