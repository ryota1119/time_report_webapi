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

type CustomerRepository struct{}

var _ repository.CustomerRepository = (*CustomerRepository)(nil)

func NewCustomerRepository() repository.CustomerRepository {
	return &CustomerRepository{}

}

func (r *CustomerRepository) Create(ctx context.Context, tx *sql.Tx, customer *entities.Customer) (*entities.CustomerID, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "INSERT INTO `customers` (`organization_id`, `name`, `unit_price`, `start_date`, `end_date`, `created_at`, `updated_at`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(
		ctx,
		query,
		organizationID,
		customer.Name,
		customer.UnitPrice,
		customer.StartDate,
		customer.EndDate,
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
	userID := entities.CustomerID(lastInsertID)

	return &userID, nil
}

func (r *CustomerRepository) List(ctx context.Context, tx *sql.Tx) ([]entities.Customer, error) {
	var customers []entities.Customer
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "SELECT `id``, `name`, `unit_price`, `start_date`, `end_date` " +
		"FROM `customers` " +
		"WHERE organization_id = ? " +
		"AND deleted_at IS NULL"
	rows, err := tx.QueryContext(
		ctx,
		query,
		&organizationID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer entities.Customer
		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.UnitPrice,
			&customer.StartDate,
			&customer.EndDate,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (r *CustomerRepository) Find(ctx context.Context, tx *sql.Tx, customerID *entities.CustomerID) (*entities.Customer, error) {
	var customer entities.Customer
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "SELECT `id`, `name`, `unit_price`, `start_date`, `end_date` " +
		"FROM `customers` " +
		"WHERE id = ? AND `organization_id` = ?" +
		"AND deleted_at IS NULL"
	result := tx.QueryRowContext(ctx, query, customerID, organizationID)
	err := result.Scan(
		&customer.ID,
		&customer.Name,
		&customer.UnitPrice,
		&customer.StartDate,
		&customer.EndDate,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrCustomerNotFound
		}
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepository) Update(ctx context.Context, tx *sql.Tx, customer *entities.Customer) error {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "UPDATE `customers` " +
		"SET `name` = ?, `unit_price` = ?, `start_date` = ?, `end_date` = ? ,  `updated_at` = ? " +
		"WHERE `id` = ? AND `organization_id` = ?" +
		"AND deleted_at IS NULL"
	_, err := tx.ExecContext(
		ctx,
		query,
		&customer.Name,
		&customer.UnitPrice,
		&customer.StartDate,
		&customer.EndDate,
		time.Now(),
		&customer.ID,
		&organizationID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) SoftDelete(ctx context.Context, tx *sql.Tx, userID *entities.CustomerID) error {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "DELETE FROM `customers` " +
		"WHERE `id` = ? AND `organization_id` = ? " +
		"AND deleted_at IS NULL"
	_, err := tx.ExecContext(
		ctx,
		query,
		userID,
		organizationID,
	)
	if err != nil {
		return err
	}

	return nil
}
