package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/domain/errors"
	"github.com/ryota1119/time_resport/internal/helper/datetime"
)

// UpdateProjectUsecaseInput ProjectUsecase Updateメソッド用input
type UpdateProjectUsecaseInput struct {
	ProjectID  uint
	CustomerID uint
	Name       string
	UnitPrice  *int64
	StartDate  *string
	EndDate    *string
}

// Update はプロジェクト情報を更新する
func (a *projectUsecase) Update(ctx context.Context, input UpdateProjectUsecaseInput) (*entities.Project, error) {
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

	// 既存のプロジェクト情報を取得
	projectID := entities.ProjectID(input.ProjectID)
	project, err := a.projectRepo.Find(ctx, tx, &projectID)
	if err != nil {
		return nil, err
	}

	// StartDateとEndDateがnilでない場合、StartDateはEndDateより前である必要がある
	// 開始日・終了日をパース
	startDate, endDate, err := datetime.ParseStartEndDate(input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}
	if startDate != nil && endDate != nil {
		if !startDate.Before(*endDate) {
			return nil, errors.ErrStartDateMustBeBefore
		}
	}

	// 何も更新がない場合は、エラーを返却し、handler層でno contentを返す
	if project.Name.String() == input.Name &&
		project.StartDate != startDate &&
		project.EndDate != endDate {
		return nil, errors.ErrNoContentUpdated
	}

	newProject := entities.NewProject(input.CustomerID, input.Name, input.UnitPrice, startDate, endDate)
	newProject.ID = projectID

	// プロジェクト情報を更新する
	_, err = a.projectRepo.Update(ctx, tx, project)
	if err != nil {
		return nil, err
	}

	return newProject, nil
}
