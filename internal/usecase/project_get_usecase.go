package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// GetProjectUsecaseInput は projectUsecase.Get のインプット
type GetProjectUsecaseInput struct {
	ProjectID uint
}

// Get はプロジェクト情報を取得する
func (a *projectUsecase) Get(ctx context.Context, input GetProjectUsecaseInput) (*entities.Project, error) {
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

	projectID := entities.ProjectID(input.ProjectID)
	project, err := a.projectRepo.Find(ctx, tx, &projectID)
	if err != nil {
		return nil, err
	}

	return project, nil
}
