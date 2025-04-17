package repository

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// CustomerRepository CustomerRepositoryのインターフェースを定義
type CustomerRepository interface {
	// Create は顧客情報を作成する
	Create(ctx context.Context, tx *sql.Tx, customer *entities.Customer) (*entities.CustomerID, error)
	// List は顧客の一覧を取得する
	List(ctx context.Context, tx *sql.Tx) ([]entities.Customer, error)
	// Find は顧客情報を取得する
	Find(ctx context.Context, tx *sql.Tx, customerID *entities.CustomerID) (*entities.Customer, error)
	// Update は顧客情報を更新する
	Update(ctx context.Context, tx *sql.Tx, customer *entities.Customer) error
	// SoftDelete は顧客情報を論理削除する
	SoftDelete(ctx context.Context, tx *sql.Tx, userID *entities.CustomerID) error
}
