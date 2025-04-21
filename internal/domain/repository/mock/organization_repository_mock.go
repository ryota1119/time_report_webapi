package mock

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
)

type OrganizationRepositoryMock struct {
	// CreateFunc は組織情報を作成するのモック
	CreateFunc func(ctx context.Context, tx *sql.Tx, organization *entities.Organization) (*entities.OrganizationID, error)
	// FindFunc は組織情報を取得するのモック
	FindFunc func(ctx context.Context, tx *sql.Tx, organizationID *entities.OrganizationID) (*entities.Organization, error)
	// FindByCodeFunc はorganizationCodeから組織情報を取得するのモック
	FindByCodeFunc func(ctx context.Context, tx *sql.Tx, organizationCode *entities.OrganizationCode) (*entities.Organization, error)
}

func (m *OrganizationRepositoryMock) Create(ctx context.Context, tx *sql.Tx, organization *entities.Organization) (*entities.OrganizationID, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, tx, organization)
	}
	return nil, nil
}

func (m *OrganizationRepositoryMock) Find(ctx context.Context, tx *sql.Tx, organizationID *entities.OrganizationID) (*entities.Organization, error) {
	if m.FindFunc != nil {
		return m.FindFunc(ctx, tx, organizationID)
	}
	return nil, nil
}

func (m *OrganizationRepositoryMock) FindByCode(ctx context.Context, tx *sql.Tx, organizationCode *entities.OrganizationCode) (*entities.Organization, error) {
	if m.FindFunc != nil {
		return m.FindByCodeFunc(ctx, tx, organizationCode)
	}
	return nil, nil
}
