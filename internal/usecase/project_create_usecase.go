package usecase

import (
	"context"
	"errors"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/helper/datetime"
)

// CreateProjectUsecaseInput ProjectUsecase Createメソッド用input
type CreateProjectUsecaseInput struct {
	CustomerID uint
	Name       string
	UnitPrice  *int64
	StartDate  *string
	EndDate    *string
}

// Create はプロジェクトを新規作成する
func (a *projectUsecase) Create(ctx context.Context, input CreateProjectUsecaseInput) (*entities.Project, error) {
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

	// StartDateとEndDateがnilでない場合、StartDateはEndDateより前である必要がある
	// 開始日・終了日をパース
	startDate, endDate, err := datetime.ParseStartEndDate(input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}
	if startDate != nil && endDate != nil {
		if !startDate.Before(*endDate) {
			return nil, errors.New("StartDate must be before EndDate")
		}
	}

	// プロジェクトを作成する
	project := entities.NewProject(input.CustomerID, input.Name, input.UnitPrice, startDate, endDate)
	_, err = a.projectRepo.Create(ctx, tx, project)
	if err != nil {
		return nil, err
	}

	return project, nil
}
