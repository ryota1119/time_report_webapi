package usecase

import (
	"context"
	"database/sql"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

var _ ProjectCreateUsecase = (*projectCreateUsecase)(nil)

// ProjectCreateUsecase は usecase.projectCreateUsecase のインターフェースを定義
type ProjectCreateUsecase interface {
	Create(ctx context.Context, input ProjectCreateUsecaseInput) (*entities.Project, error)
}

// projectCreateUsecase ユースケース
type projectCreateUsecase struct {
	db           *sql.DB
	projectRepo  repository.ProjectRepository
	customerRepo repository.CustomerRepository
}

// NewProjectCreateUsecase は projectCreateUsecase を初期化する
func NewProjectCreateUsecase(
	db *sql.DB,
	projectRepo repository.ProjectRepository,
	customerRepo repository.CustomerRepository,
) ProjectCreateUsecase {
	return &projectCreateUsecase{
		db:           db,
		projectRepo:  projectRepo,
		customerRepo: customerRepo,
	}
}

// ProjectCreateUsecaseInput projectCreateUsecase.Createメソッド用input
type ProjectCreateUsecaseInput struct {
	CustomerID uint
	Name       string
	UnitPrice  *int64
	StartDate  *string
	EndDate    *string
}

// Create はプロジェクトを新規作成する
func (a *projectCreateUsecase) Create(ctx context.Context, input ProjectCreateUsecaseInput) (*entities.Project, error) {
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

	// 顧客が存在するか確認
	customerID := entities.CustomerID(input.CustomerID)
	_, err = a.customerRepo.Find(ctx, tx, &customerID)
	if err != nil {
		return nil, err
	}

	// プロジェクトを作成する
	project, err := entities.NewProject(input.CustomerID, input.Name, input.UnitPrice, input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}
	projectID, err := a.projectRepo.Create(ctx, tx, project)
	if err != nil {
		return nil, err
	}
	project.ID = *projectID

	return project, nil
}
