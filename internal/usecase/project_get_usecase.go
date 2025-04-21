package usecase

import (
	"context"
	"database/sql"
	"github.com/ryota1119/time_resport_webapi/internal/usecase/dto"

	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ ProjectGetUsecase = (*projectGetUsecase)(nil)

// ProjectGetUsecase は usecase.projectGetUsecase のインターフェースを定義
type ProjectGetUsecase interface {
	Get(ctx context.Context, input ProjectGetUsecaseInput) (*dto.ProjectWithCustomer, error)
}

// projectGetUsecase ユースケース
type projectGetUsecase struct {
	db           *sql.DB
	projectRepo  repository.ProjectRepository
	customerRepo repository.CustomerRepository
}

// NewProjectGetUsecase は projectGetUsecase を初期化する
func NewProjectGetUsecase(
	db *sql.DB,
	projectRepo repository.ProjectRepository,
	customerRepo repository.CustomerRepository,
) ProjectGetUsecase {
	return &projectGetUsecase{
		db:           db,
		projectRepo:  projectRepo,
		customerRepo: customerRepo,
	}
}

// ProjectGetUsecaseInput は projectUsecase.Get のインプット
type ProjectGetUsecaseInput struct {
	ProjectID uint
}

// Get はプロジェクト情報を取得する
func (a *projectGetUsecase) Get(ctx context.Context, input ProjectGetUsecaseInput) (*dto.ProjectWithCustomer, error) {
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

	customer, err := a.customerRepo.Find(ctx, tx, &project.CustomerID)
	if err != nil {
		return nil, err
	}

	results := dto.ProjectWithCustomer{
		Project:  project,
		Customer: customer,
	}

	return &results, nil
}
