package usecase

import (
	"context"
	"database/sql"
	"github.com/ryota1119/time_resport/internal/domain/repository"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

var _ BudgetGetUsecase = (*budgetGetUsecase)(nil)

// BudgetGetUsecase BudgetUsecaseのインターフェースを定義
type BudgetGetUsecase interface {
	// Get は予算を新規作成する
	Get(ctx context.Context, input BudgetUsecaseGetInput) (*entities.BudgetWithProject, error)
}

// budgetGetUsecase ユースケース
type budgetGetUsecase struct {
	db         *sql.DB
	budgetRepo repository.BudgetRepository
}

// NewBudgetGetUsecase は budgetGetUsecase を初期化する
func NewBudgetGetUsecase(
	db *sql.DB,
	budgetRepo repository.BudgetRepository,
) BudgetGetUsecase {
	return &budgetGetUsecase{
		db:         db,
		budgetRepo: budgetRepo,
	}
}

// BudgetUsecaseGetInput は budgetUsecase.Get のインプット
type BudgetUsecaseGetInput struct {
	BudgetID uint
}

// Get は予算情報を取得する
func (a *budgetGetUsecase) Get(ctx context.Context, input BudgetUsecaseGetInput) (*entities.BudgetWithProject, error) {
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

	budgetID := entities.BudgetID(input.BudgetID)
	budget, err := a.budgetRepo.FindWithProject(ctx, tx, &budgetID)
	if err != nil {
		return nil, err
	}

	return budget, nil
}
