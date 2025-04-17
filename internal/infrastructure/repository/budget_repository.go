package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	domainerrors "github.com/ryota1119/time_resport/internal/domain/errors"
	"github.com/ryota1119/time_resport/internal/domain/repository"
	"github.com/ryota1119/time_resport/internal/helper/auth_context"
)

type BudgetRepository struct{}

var _ repository.BudgetRepository = (*BudgetRepository)(nil)

func NewBudgetRepository() repository.BudgetRepository {
	return &BudgetRepository{}

}

func (r *BudgetRepository) Create(ctx context.Context, tx *sql.Tx, budget *entities.Budget) (*entities.BudgetID, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "INSERT INTO `budgets` (`organization_id`, `project_id`, `amount`, `memo`, `start_date`, `end_date`, `created_at`, `updated_at`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(
		ctx,
		query,
		organizationID,
		budget.ProjectID,
		budget.BudgetAmount,
		budget.BudgetMemo,
		budget.BudgetPeriod.Start.Value(),
		budget.BudgetPeriod.End.Value(),
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
	userID := entities.BudgetID(lastInsertID)

	return &userID, nil
}

func (r *BudgetRepository) List(ctx context.Context, tx *sql.Tx) ([]entities.Budget, error) {
	var budgets []entities.Budget
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "SELECT `id`, `project_id`, `amount`, `memo`, `start_date`, `end_date` " +
		"FROM `budgets` " +
		"WHERE `organization_id` = ? "
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
		var budget entities.Budget

		err := rows.Scan(
			&budget.ID,
			&budget.ProjectID,
			&budget.BudgetAmount,
			&budget.BudgetMemo,
			&budget.BudgetPeriod.Start,
			&budget.BudgetPeriod.End,
		)
		if err != nil {
			return nil, err
		}
		budgets = append(budgets, budget)
	}
	return budgets, nil
}

func (r *BudgetRepository) Find(ctx context.Context, tx *sql.Tx, budgetID *entities.BudgetID) (*entities.Budget, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	var budget entities.Budget

	query := "SELECT `id`, `project_id`, `amount`, `memo`, `start_date`, `end_date` " +
		"FROM `budgets` " +
		"WHERE `id` = ? AND `organization_id` = ? "
	result := tx.QueryRowContext(ctx, query, budgetID, organizationID)
	err := result.Scan(
		&budget.ID,
		&budget.ProjectID,
		&budget.BudgetAmount,
		&budget.BudgetMemo,
		&budget.BudgetPeriod.Start,
		&budget.BudgetPeriod.End,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainerrors.ErrBudgetNotFound
		}
		return nil, err
	}

	return &budget, nil
}

func (r *BudgetRepository) FindWithProject(ctx context.Context, tx *sql.Tx, budgetID *entities.BudgetID) (*entities.BudgetWithProject, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	var budgetWithProject entities.BudgetWithProject
	var projectUnitPrice sql.NullInt64
	var budgetMemo sql.NullString

	query := "SELECT `b`.`id`, `p`.`name`, `p`.`unit_price`, `b`.`amount`, `b`.`memo`, `b`.`start_date`, `b`.`end_date` " +
		"FROM `budgets` AS `b` " +
		"LEFT JOIN `projects` AS `p` ON `p`.`id` = `b`.`project_id` " +
		"WHERE `b`.`id` = ? AND `b`.`organization_id` = ? "
	result := tx.QueryRowContext(ctx, query, budgetID, organizationID)
	err := result.Scan(
		&budgetWithProject.BudgetID,
		&budgetWithProject.ProjectName,
		&projectUnitPrice,
		&budgetWithProject.BudgetAmount,
		&budgetMemo,
		&budgetWithProject.BudgetPeriod.Start,
		&budgetWithProject.BudgetPeriod.End,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainerrors.ErrBudgetNotFound
		}
		return nil, err
	}

	if projectUnitPrice.Valid {
		pUnitPrice := entities.ProjectUnitPrice(projectUnitPrice.Int64)
		budgetWithProject.ProjectUnitPrice = &pUnitPrice
	}
	if budgetMemo.Valid {
		bMemo := entities.BudgetMemo(budgetMemo.String)
		budgetWithProject.BudgetMemo = &bMemo
	}

	return &budgetWithProject, nil
}

func (r *BudgetRepository) Update(ctx context.Context, tx *sql.Tx, budget *entities.Budget) (*entities.BudgetID, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "UPDATE `budgets` " +
		"SET `amount` = ?, `memo` = ?, `start_date` = ?, `end_date` = ? ,  `updated_at` = ? " +
		"WHERE `id` = ? AND `organization_id` = ?"
	result, err := tx.ExecContext(
		ctx,
		query,
		budget.BudgetAmount,
		budget.BudgetMemo,
		budget.BudgetPeriod.Start.Value(),
		budget.BudgetPeriod.End.Value(),
		time.Now(),
		budget.ID,
		organizationID,
	)
	if err != nil {
		return nil, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	budgetID := entities.BudgetID(lastInsertID)

	return &budgetID, nil
}

func (r *BudgetRepository) Delete(ctx context.Context, tx *sql.Tx, budgetID *entities.BudgetID) error {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "DELETE FROM `budgets` " +
		"WHERE `id` = ? AND `organization_id` = ?"
	_, err := tx.ExecContext(
		ctx,
		query,
		budgetID,
		organizationID,
	)
	if err != nil {
		return err
	}

	return nil
}
