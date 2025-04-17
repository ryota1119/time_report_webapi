package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/domain/repository"
)

// ProjectUsecase ProjectUsecaseのインターフェースを定義
type ProjectUsecase interface {
	// Create はプロジェクトを新規作成する
	Create(ctx context.Context, input CreateProjectUsecaseInput) (*entities.Project, error)
	// List はプロジェクトの一覧を取得する
	List(ctx context.Context) ([]entities.Project, error)
	// Get はプロジェクト情報を取得する
	Get(ctx context.Context, input GetProjectUsecaseInput) (*entities.Project, error)
	// Update はプロジェクト情報を更新する
	Update(ctx context.Context, input UpdateProjectUsecaseInput) (*entities.Project, error)
	// SoftDelete はプロジェクト情報を削除する
	SoftDelete(ctx context.Context, input DeleteProjectUsecaseInput) error
}

type projectUsecase struct {
	db          *sql.DB
	projectRepo repository.ProjectRepository
}

var _ ProjectUsecase = (*projectUsecase)(nil)

func NewProjectUsecase(
	db *sql.DB,
	projectRepo repository.ProjectRepository,
) ProjectUsecase {
	return &projectUsecase{
		db:          db,
		projectRepo: projectRepo,
	}
}
