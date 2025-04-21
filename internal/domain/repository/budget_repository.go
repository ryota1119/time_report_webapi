package repository

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// BudgetRepository BudgetRepositoryのインターフェースを定義
type BudgetRepository interface {
	// Create は予算情報を作成する
	Create(ctx context.Context, tx *sql.Tx, budget *entities.Budget) (*entities.BudgetID, error)
	// List は予算の一覧を取得する
	List(ctx context.Context, tx *sql.Tx) ([]entities.Budget, error)
	// Find は予算情報を取得する
	Find(ctx context.Context, tx *sql.Tx, BudgetID *entities.BudgetID) (*entities.Budget, error)
	// FindWithProject は予算情報を取得する
	FindWithProject(ctx context.Context, tx *sql.Tx, budgetID *entities.BudgetID) (*entities.BudgetWithProject, error)
	// Update は予算情報を更新する
	Update(ctx context.Context, tx *sql.Tx, budget *entities.Budget) (*entities.BudgetID, error)
	// Delete は予算情報を削除する
	Delete(ctx context.Context, tx *sql.Tx, budgetID *entities.BudgetID) error
}
