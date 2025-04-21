package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"
)

type OrganizationRepository struct{}

var _ repository.OrganizationRepository = (*OrganizationRepository)(nil)

func NewOrganizationRepository() repository.OrganizationRepository {
	return &OrganizationRepository{}
}

func (r *OrganizationRepository) Create(ctx context.Context, tx *sql.Tx, organization *entities.Organization) (*entities.OrganizationID, error) {
	query := "INSERT INTO `organizations` (`organization_name`, `organization_code`, `created_at`, `updated_at`) " +
		"VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(
		ctx,
		query,
		organization.Name,
		organization.Code,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	organizationID := entities.OrganizationID(lastInsertID)

	return &organizationID, nil
}

func (r *OrganizationRepository) Find(ctx context.Context, tx *sql.Tx, organizationID *entities.OrganizationID) (*entities.Organization, error) {
	var organization entities.Organization

	query := "SELECT `id`, `organization_name`, `organization_code` FROM `organizations` WHERE `id` = ?"
	result := tx.QueryRowContext(ctx, query, organizationID)
	err := result.Scan(
		&organization.ID,
		&organization.Name,
		&organization.Code,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrOrganizationNotFound
		}
		return nil, err
	}

	return &organization, nil
}

func (r *OrganizationRepository) FindByCode(ctx context.Context, tx *sql.Tx, organizationCode *entities.OrganizationCode) (*entities.Organization, error) {
	var organization entities.Organization

	query := "SELECT `id`, `organization_name`, `organization_code` FROM `organizations` WHERE `organization_code` = ?"
	result := tx.QueryRowContext(ctx, query, organizationCode)
	err := result.Scan(
		&organization.ID,
		&organization.Name,
		&organization.Code,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrOrganizationNotFound
		}
		return nil, err
	}

	return &organization, nil
}
