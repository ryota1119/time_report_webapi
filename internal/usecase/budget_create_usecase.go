package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ BudgetCreateUsecase = (*budgetCreateUsecase)(nil)

// BudgetCreateUsecase BudgetUsecaseのインターフェースを定義
type BudgetCreateUsecase interface {
	// Create は予算を新規作成する
	Create(ctx context.Context, input BudgetUsecaseCreateInput) (*entities.Budget, error)
}

// budgetCreateUsecase ユースケース
type budgetCreateUsecase struct {
	db         *sql.DB
	budgetRepo repository.BudgetRepository
}

func NewBudgetCreateUsecase(
	db *sql.DB,
	budgetRepo repository.BudgetRepository,
) BudgetCreateUsecase {
	return &budgetCreateUsecase{
		db:         db,
		budgetRepo: budgetRepo,
	}
}

// BudgetUsecaseCreateInput BudgetUsecase Createメソッド用input
type BudgetUsecaseCreateInput struct {
	ProjectID    uint
	BudgetAmount int64
	BudgetMemo   *string
	StartDate    string
	EndDate      string
}

// Create は予算を新規作成する
func (a *budgetCreateUsecase) Create(ctx context.Context, input BudgetUsecaseCreateInput) (*entities.Budget, error) {
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

	// 予算を作成する
	budget, err := entities.NewBudget(input.ProjectID, input.BudgetAmount, input.BudgetMemo, input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}
	budgetID, err := a.budgetRepo.Create(ctx, tx, budget)
	if err != nil {
		return nil, err
	}
	budget.ID = *budgetID

	return budget, nil
}
