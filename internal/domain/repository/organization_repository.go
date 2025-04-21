package repository

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

// OrganizationRepository OrganizationRepositoryのインターフェースを定義
type OrganizationRepository interface {
	// Create は組織情報を作成する
	Create(ctx context.Context, tx *sql.Tx, organization *entities.Organization) (*entities.OrganizationID, error)
	// Find は組織情報を取得する
	Find(ctx context.Context, tx *sql.Tx, organizationID *entities.OrganizationID) (*entities.Organization, error)
	// FindByCode はorganizationCodeから組織情報を取得する
	FindByCode(ctx context.Context, tx *sql.Tx, organizationCode *entities.OrganizationCode) (*entities.Organization, error)
}
