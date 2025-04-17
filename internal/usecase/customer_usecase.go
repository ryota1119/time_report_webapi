package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/domain/repository"
)

// CustomerUsecase CustomerUsecaseのインターフェースを定義
type CustomerUsecase interface {
	// Create は顧客を新規作成する
	Create(ctx context.Context, input CreateCustomerUsecaseInput) (*entities.Customer, error)
	// List は顧客の一覧を取得する
	List(ctx context.Context) ([]entities.Customer, error)
	// Get は顧客情報を取得する
	Get(ctx context.Context, input GetCustomerUsecaseInput) (*entities.Customer, error)
	// Update は顧客情報を更新する
	Update(ctx context.Context, input UpdateCustomerUsecaseInput) (*entities.Customer, error)
	// SoftDelete は顧客情報を論理削除する
	SoftDelete(ctx context.Context, input SoftDeleteCustomerUsecaseInput) error
}

type customerUsecase struct {
	db           *sql.DB
	customerRepo repository.CustomerRepository
}

var _ CustomerUsecase = (*customerUsecase)(nil)

func NewCustomerUsecase(
	db *sql.DB,
	customerRepo repository.CustomerRepository,
) CustomerUsecase {
	return &customerUsecase{
		db:           db,
		customerRepo: customerRepo,
	}
}
