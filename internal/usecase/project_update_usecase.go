package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ ProjectUpdateUsecase = (*projectUpdateUsecase)(nil)

// ProjectUpdateUsecase は usecase.projectUpdateUsecase のインターフェースを定義
type ProjectUpdateUsecase interface {
	Update(ctx context.Context, input ProjectUpdateUsecaseInput) (*entities.Project, error)
}

// projectUpdateUsecase ユースケース
type projectUpdateUsecase struct {
	db          *sql.DB
	projectRepo repository.ProjectRepository
}

// NewProjectUpdateUsecase は projectUpdateUsecase を初期化する
func NewProjectUpdateUsecase(
	db *sql.DB,
	projectRepo repository.ProjectRepository,
) ProjectUpdateUsecase {
	return &projectUpdateUsecase{
		db:          db,
		projectRepo: projectRepo,
	}
}

// ProjectUpdateUsecaseInput ProjectUsecase Updateメソッド用input
type ProjectUpdateUsecaseInput struct {
	ProjectID  uint
	CustomerID uint
	Name       string
	UnitPrice  *int64
	StartDate  *string
	EndDate    *string
}

// Update はプロジェクト情報を更新する
func (a *projectUpdateUsecase) Update(ctx context.Context, input ProjectUpdateUsecaseInput) (*entities.Project, error) {
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

	period, err := entities.NewProjectPeriod(input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}
	unitPrice := entities.NewProjectUnitPrice(input.UnitPrice)

	isUpdated := false

	if project.Name != entities.ProjectName(input.Name) {
		project.Name = entities.ProjectName(input.Name)
		isUpdated = true
	}
	if project.UnitPrice != nil && unitPrice == nil ||
		project.UnitPrice == nil && unitPrice != nil ||
		project.UnitPrice != nil && unitPrice != nil {
		project.UnitPrice = unitPrice
		isUpdated = true
	}
	if project.Period.Start == nil && period.Start != nil ||
		project.Period.Start != nil && period.Start == nil ||
		(project.Period.Start != nil && period.Start != nil && !project.Period.Start.Equal(*period.Start)) {
		project.Period.Start = period.Start
		isUpdated = true
	}
	if project.Period.End == nil && period.End != nil ||
		project.Period.End != nil && period.End == nil ||
		(project.Period.End != nil && period.End != nil && !project.Period.End.Equal(*period.End)) {
		project.Period.End = period.End
		isUpdated = true
	}

	if !isUpdated {
		return nil, entities.ErrNoContentUpdated
	}

	// プロジェクト情報を更新する
	if err = a.projectRepo.Update(ctx, tx, project); err != nil {
		return nil, err
	}

	return project, nil
}
