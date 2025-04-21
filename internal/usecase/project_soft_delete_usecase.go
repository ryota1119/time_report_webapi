package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ ProjectSoftDeleteUsecase = (*projectSoftDeleteUsecase)(nil)

// ProjectSoftDeleteUsecase は usecase.projectSoftDeleteUsecase のインターフェースを定義
type ProjectSoftDeleteUsecase interface {
	SoftDelete(ctx context.Context, input ProjectSoftDeleteUsecaseInput) error
}

// projectSoftDeleteUsecase ユースケース
type projectSoftDeleteUsecase struct {
	db          *sql.DB
	projectRepo repository.ProjectRepository
}

// NewProjectSoftDeleteUsecase は projectSoftDeleteUsecase を初期化する
func NewProjectSoftDeleteUsecase(
	db *sql.DB,
	projectRepo repository.ProjectRepository,
) ProjectSoftDeleteUsecase {
	return &projectSoftDeleteUsecase{
		db:          db,
		projectRepo: projectRepo,
	}
}

// ProjectSoftDeleteUsecaseInput は projectUsecase.SoftDelete のインプット
type ProjectSoftDeleteUsecaseInput struct {
	ProjectID uint
}

// SoftDelete はプロジェクト情報を論理削除する
func (a *projectSoftDeleteUsecase) SoftDelete(ctx context.Context, input ProjectSoftDeleteUsecaseInput) error {
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
	err = a.projectRepo.SoftDelete(ctx, tx, &project.ID)
	if err != nil {
		return err
	}
	return nil
}
