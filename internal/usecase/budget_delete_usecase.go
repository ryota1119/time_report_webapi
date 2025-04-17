package usecase

import (
	"context"
	"database/sql"
	"github.com/ryota1119/time_resport/internal/domain/repository"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

var _ BudgetDeleteUsecase = (*budgetDeleteUsecase)(nil)

// BudgetDeleteUsecase BudgetUsecaseのインターフェースを定義
type BudgetDeleteUsecase interface {
	// Delete は予算を新規作成する
	Delete(ctx context.Context, input DeleteBudgetUsecaseInput) error
}

// budgetDeleteUsecase ユースケース
type budgetDeleteUsecase struct {
	db         *sql.DB
	budgetRepo repository.BudgetRepository
}

func NewBudgetDeleteUsecase(
	db *sql.DB,
	budgetRepo repository.BudgetRepository,
) BudgetDeleteUsecase {
	return &budgetDeleteUsecase{
		db:         db,
		budgetRepo: budgetRepo,
	}
}

// DeleteBudgetUsecaseInput は budgetUsecase.SoftDelete のインプット
type DeleteBudgetUsecaseInput struct {
	BudgetID uint
}

// Delete は予算情報を削除する
func (a *budgetDeleteUsecase) Delete(ctx context.Context, input DeleteBudgetUsecaseInput) error {
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

	// 既存の予算情報を取得する
	budgetID := entities.BudgetID(input.BudgetID)
	budget, err := a.budgetRepo.Find(ctx, tx, &budgetID)
	if err != nil {
		return err
	}

	// 予算情報を削除する
	err = a.budgetRepo.Delete(ctx, tx, &budget.ID)
	if err != nil {
		return err
	}
	return nil
}
