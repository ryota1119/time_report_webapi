package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ryota1119/time_resport/internal/helper/auth_context"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/domain/repository"
)

type ProjectRepository struct{}

var _ repository.ProjectRepository = (*ProjectRepository)(nil)

func NewProjectRepository() repository.ProjectRepository {
	return &ProjectRepository{}

}

func (r *ProjectRepository) Create(ctx context.Context, tx *sql.Tx, project *entities.Project) (*entities.ProjectID, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "INSERT INTO `projects` (`organization_id`, `customer_id`, `name`, `unit_price`, `start_date`, `end_date`, `created_at`, `updated_at`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(
		ctx,
		query,
		organizationID,
		project.CustomerID,
		project.Name,
		project.UnitPrice,
		project.StartDate,
		project.EndDate,
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
	userID := entities.ProjectID(lastInsertID)

	return &userID, nil
}

func (r *ProjectRepository) List(ctx context.Context, tx *sql.Tx) ([]entities.Project, error) {
	var projects []entities.Project
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "SELECT `p`.`id`, `p`.`customer_id`, `p`.`name`, `p`.`unit_price`, `p`.`start_date`, `p`.`end_date` " +
		"FROM `projects` AS `p` " +
		"LEFT JOIN `customers` AS `c` ON `p`.`customer_id` = `c`.`id` " +
		"WHERE `c`.`organization_id` = ? " +
		"`deleted_at` IS NULL"
	rows, err := tx.QueryContext(
		ctx,
		query,
		organizationID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var project entities.Project

		err := rows.Scan(
			&project.ID,
			&project.CustomerID,
			&project.Name,
			&project.UnitPrice,
			&project.StartDate,
			&project.EndDate,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (r *ProjectRepository) Find(ctx context.Context, tx *sql.Tx, projectID *entities.ProjectID) (*entities.Project, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	var project entities.Project

	query := "SELECT `p`.`id`, `p`.`customer_id`, `p`.`name`, `p`.`unit_price`, `p`.`start_date`, `p`.`end_date` " +
		"FROM `projects` AS `p` " +
		"LEFT JOIN `customers` AS `c` ON `p`.`customer_id` = `c`.`id` " +
		"WHERE `p`.`id` = ? AND `c`.`organization_id` = ? " +
		"AND `deleted_at` IS NULL"
	result := tx.QueryRowContext(ctx, query, projectID, organizationID)
	err := result.Scan(
		&project.ID,
		&project.CustomerID,
		&project.Name,
		&project.UnitPrice,
		&project.StartDate,
		&project.EndDate,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrProjectNotFound
		}
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) Update(ctx context.Context, tx *sql.Tx, project *entities.Project) (*entities.ProjectID, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "UPDATE `projects` AS `p` " +
		"SET `p`.`name` = ?, `p`.`unit_price` = ?, `p`.`start_date` = ?, `p`.`end_date` = ? ,  `p`.`updated_at` = ? " +
		"LEFT JOIN `customers` AS `c` ON `p`.`customer_id` = `c`.`id` " +
		"WHERE `p`.`id` = ? AND `c`.`organization_id` = ?"
	result, err := tx.ExecContext(
		ctx,
		query,
		&project.Name,
		&project.UnitPrice,
		&project.StartDate,
		&project.EndDate,
		time.Now(),
		&project.ID,
		&organizationID,
	)
	if err != nil {
		return nil, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	projectID := entities.ProjectID(lastInsertID)

	return &projectID, nil
}

func (r *ProjectRepository) Delete(ctx context.Context, tx *sql.Tx, projectID *entities.ProjectID) error {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "DELETE FROM `projects` AS `p` " +
		"LEFT JOIN `customers` AS `c` ON `p`.`customer_id` = `c`.`id` " +
		"WHERE `p`.`id` = ? AND `c`.`organization_id` = ?"
	_, err := tx.ExecContext(
		ctx,
		query,
		projectID,
		organizationID,
	)
	if err != nil {
		return err
	}

	return nil
}
