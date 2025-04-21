package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"
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
		customer.Period.Start.Value(),
		customer.Period.End.Value(),
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

	query := "SELECT `id`, `name`, `unit_price`, `start_date`, `end_date` " +
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
			&customer.Period.Start,
			&customer.Period.End,
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
		"WHERE id = ? AND `organization_id` = ? " +
		"AND `deleted_at` IS NULL"
	result := tx.QueryRowContext(ctx, query, customerID, organizationID)
	err := result.Scan(
		&customer.ID,
		&customer.Name,
		&customer.UnitPrice,
		&customer.Period.Start,
		&customer.Period.End,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrCustomerNotFound
		}
		return nil, err
	}

	return &customer, nil
}

// FindByIDs は entities.CustomerID スライスから顧客情報のスライスを取得する
func (r *CustomerRepository) FindByIDs(ctx context.Context, tx *sql.Tx, customerIDs []entities.CustomerID) ([]entities.Customer, error) {
	if len(customerIDs) == 0 {
		return []entities.Customer{}, nil
	}

	organizationID := auth_context.ContextOrganizationID(ctx)

	placeholders := make([]string, len(customerIDs))
	args := make([]interface{}, 0, len(customerIDs)+1)
	for i, id := range customerIDs {
		placeholders[i] = "?"
		args = append(args, id)
	}
	args = append(args, organizationID)

	query := fmt.Sprintf("SELECT `id`, `name`, `unit_price`, `start_date`, `end_date` "+
		"FROM `customers` "+
		"WHERE id IN (%s) AND `organization_id` = ? "+
		"AND `deleted_at` IS NULL", strings.Join(placeholders, ","))
	rows, err := tx.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		var (
			id        uint
			name      string
			unitPrice *int64
			startDate *time.Time
			endDate   *time.Time
		)
		if err := rows.Scan(
			&id,
			&name,
			&unitPrice,
			&startDate,
			&endDate,
		); err != nil {
			return nil, err
		}
		customer := entities.Customer{
			ID:        entities.CustomerID(id),
			Name:      entities.CustomerName(name),
			UnitPrice: entities.NewCustomerUnitPrice(unitPrice),
			Period: entities.CustomerPeriod{
				Start: (*entities.CustomerStartDate)(startDate),
				End:   (*entities.CustomerEndDate)(endDate),
			},
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (r *CustomerRepository) Update(ctx context.Context, tx *sql.Tx, customer *entities.Customer) error {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "UPDATE `customers` " +
		"SET `name` = ?, `unit_price` = ?, `start_date` = ?, `end_date` = ? ,  `updated_at` = ? " +
		"WHERE `id` = ? AND `organization_id` = ? " +
		"AND deleted_at IS NULL"
	_, err := tx.ExecContext(
		ctx,
		query,
		customer.Name,
		customer.UnitPrice,
		customer.Period.Start.Value(),
		customer.Period.End.Value(),
		time.Now(),
		&customer.ID,
		&organizationID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) SoftDelete(ctx context.Context, tx *sql.Tx, customerID *entities.CustomerID) error {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "UPDATE `customers` " +
		"SET `deleted_at` = ? ,  `updated_at` = ? " +
		"WHERE `id` = ? " +
		"AND `organization_id` = ? " +
		"AND `deleted_at` IS NULL"
	_, err := tx.ExecContext(
		ctx,
		query,
		time.Now(),
		time.Now(),
		customerID,
		organizationID,
	)
	if err != nil {
		return err
	}

	return nil
}
