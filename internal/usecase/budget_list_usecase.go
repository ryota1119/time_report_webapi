package usecase

import (
	"context"
	"database/sql"
	"github.com/ryota1119/time_resport/internal/domain/repository"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

var _ BudgetListUsecase = (*budgetListUsecase)(nil)

// BudgetListUsecase BudgetUsecaseのインターフェースを定義
type BudgetListUsecase interface {
	// List は予算一覧を取得する
	List(ctx context.Context) ([]entities.Budget, error)
}

// budgetListUsecase ユースケース
type budgetListUsecase struct {
	db         *sql.DB
	budgetRepo repository.BudgetRepository
}

func NewBudgetListUsecase(
	db *sql.DB,
	budgetRepo repository.BudgetRepository,
) BudgetListUsecase {
	return &budgetListUsecase{
		db:         db,
		budgetRepo: budgetRepo,
	}
}

// List は予算の一覧を取得する
func (a *budgetListUsecase) List(ctx context.Context) ([]entities.Budget, error) {
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

	// 予算リストを取得する
	budgets, err := a.budgetRepo.List(ctx, tx)
	if err != nil {
		return nil, err
	}

	return budgets, nil
}
