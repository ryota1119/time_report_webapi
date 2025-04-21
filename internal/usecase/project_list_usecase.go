package usecase

import (
	"context"
	"database/sql"
	"github.com/ryota1119/time_resport_webapi/internal/usecase/dto"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"
)

var _ ProjectListUsecase = (*projectListUsecase)(nil)

// ProjectListUsecase は usecase.projectListUsecase のインターフェースを定義
type ProjectListUsecase interface {
	List(ctx context.Context) ([]*dto.ProjectWithCustomer, error)
}

// projectListUsecase ユースケース
type projectListUsecase struct {
	db           *sql.DB
	projectRepo  repository.ProjectRepository
	customerRepo repository.CustomerRepository
}

// NewProjectListUsecase は projectListUsecase を初期化する
func NewProjectListUsecase(
	db *sql.DB,
	projectRepo repository.ProjectRepository,
	customerRepo repository.CustomerRepository,
) ProjectListUsecase {
	return &projectListUsecase{
		db:           db,
		projectRepo:  projectRepo,
		customerRepo: customerRepo,
	}
}

// List はプロジェクトの一覧を取得する
func (a *projectListUsecase) List(ctx context.Context) ([]*dto.ProjectWithCustomer, error) {
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

	customerIDSet := map[entities.CustomerID]struct{}{}
	for _, project := range projects {
		customerIDSet[project.CustomerID] = struct{}{}
	}
	var customerIDs []entities.CustomerID
	for customerID := range customerIDSet {
		customerIDs = append(customerIDs, customerID)
	}

	customers, err := a.customerRepo.FindByIDs(ctx, tx, customerIDs)
	if err != nil {
		return nil, err
	}
	customerMap := toCustomerMap(customers)

	var results []*dto.ProjectWithCustomer
	for _, project := range projects {
		c := customerMap[project.CustomerID]
		results = append(results, &dto.ProjectWithCustomer{
			Project:  &project,
			Customer: c,
		})
	}

	return results, nil
}

// toCustomerMap は entities.Customer のスライスを受け取っって map に変換する
func toCustomerMap(list []entities.Customer) map[entities.CustomerID]*entities.Customer {
	m := make(map[entities.CustomerID]*entities.Customer)
	for i := range list {
		c := list[i]
		m[c.ID] = &c
	}
	return m
}
