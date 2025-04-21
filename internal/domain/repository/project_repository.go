package repository

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// ProjectRepository ProjectRepositoryのインターフェースを定義
type ProjectRepository interface {
	// Create は顧客情報を作成する
	Create(ctx context.Context, tx *sql.Tx, project *entities.Project) (*entities.ProjectID, error)
	// List は顧客の一覧を取得する
	List(ctx context.Context, tx *sql.Tx) ([]entities.Project, error)
	// Find は顧客情報を取得する
	Find(ctx context.Context, tx *sql.Tx, projectID *entities.ProjectID) (*entities.Project, error)
	// Update は顧客情報を更新する
	Update(ctx context.Context, tx *sql.Tx, project *entities.Project) error
	// SoftDelete は顧客情報を論理削除する
	SoftDelete(ctx context.Context, tx *sql.Tx, projectID *entities.ProjectID) error
}
