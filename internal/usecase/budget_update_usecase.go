package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ BudgetUpdateUsecase = (*budgetUpdateUsecase)(nil)

// BudgetUpdateUsecase BudgetUsecaseのインターフェースを定義
type BudgetUpdateUsecase interface {
	// Update は予算を新規作成する
	Update(ctx context.Context, input BudgetUsecaseUpdateInput) (*entities.Budget, error)
}

// budgetUpdateUsecase ユースケース
type budgetUpdateUsecase struct {
	db         *sql.DB
	budgetRepo repository.BudgetRepository
}

func NewBudgetUpdateUsecase(
	db *sql.DB,
	budgetRepo repository.BudgetRepository,
) BudgetUpdateUsecase {
	return &budgetUpdateUsecase{
		db:         db,
		budgetRepo: budgetRepo,
	}
}

// BudgetUsecaseUpdateInput BudgetUsecase Updateメソッド用input
type BudgetUsecaseUpdateInput struct {
	BudgetID     uint
	ProjectID    uint
	BudgetAmount int64
	BudgetMemo   *string
	StartDate    string
	EndDate      string
}

// Update は予算情報を更新する
func (a *budgetUpdateUsecase) Update(ctx context.Context, input BudgetUsecaseUpdateInput) (*entities.Budget, error) {
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

	// 既存の予算情報を取得
	budgetID := entities.BudgetID(input.BudgetID)
	budget, err := a.budgetRepo.Find(ctx, tx, &budgetID)
	if err != nil {
		return nil, err
	}

	newBudget, err := entities.NewBudget(input.ProjectID, input.BudgetAmount, input.BudgetMemo, input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}
	newBudget.ID = budgetID

	// 予算情報を更新する
	_, err = a.budgetRepo.Update(ctx, tx, budget)
	if err != nil {
		return nil, err
	}

	return newBudget, nil
}
