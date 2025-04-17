package usecase

import (
	"context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// List はプロジェクトの一覧を取得する
func (a *projectUsecase) List(ctx context.Context) ([]entities.Project, error) {
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

	// プロジェクトリストを取得する
	projects, err := a.projectRepo.List(ctx, tx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
