package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// DeleteProjectUsecaseInput は projectUsecase.SoftDelete のインプット
type DeleteProjectUsecaseInput struct {
	ProjectID uint
}

// SoftDelete はプロジェクト情報を論理削除する
func (a *projectUsecase) SoftDelete(ctx context.Context, input DeleteProjectUsecaseInput) error {
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

	// 既存のプロジェクト情報を取得する
	projectID := entities.ProjectID(input.ProjectID)
	project, err := a.projectRepo.Find(ctx, tx, &projectID)
	if err != nil {
		return err
	}

	// プロジェクト情報を削除する
	err = a.projectRepo.Delete(ctx, tx, &project.ID)
	if err != nil {
		return err
	}
	return nil
}
