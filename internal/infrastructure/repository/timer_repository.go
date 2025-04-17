package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/domain/repository"
	"github.com/ryota1119/time_resport/internal/helper/auth_context"
)

type TimerRepository struct{}

var _ repository.TimerRepository = (*TimerRepository)(nil)

func NewTimerRepository() repository.TimerRepository {
	return &TimerRepository{}
}

func (r *TimerRepository) Create(ctx context.Context, tx *sql.Tx, timer *entities.Timer) (*entities.TimerID, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)
	userID := auth_context.ContextUserID(ctx)

	var endAt driver.Value
	if timer.EndAt != nil {
		endAt = timer.EndAt.Value()
	} else {
		endAt = nil
	}

	query := "INSERT INTO `timers` (`organization_id`, `user_id`, `project_id`, `title`, `memo`, `start_at`, `end_at`, `created_at`, `updated_at`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(
		ctx,
		query,
		organizationID,
		userID,
		timer.ProjectID,
		timer.Title,
		timer.Memo,
		timer.StartAt.Value(),
		endAt,
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
	timerID := entities.TimerID(lastInsertID)

	return &timerID, nil
}

func (r *TimerRepository) ExistsRunningTimer(ctx context.Context, tx *sql.Tx) (bool, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)
	userID := auth_context.ContextUserID(ctx)

	query := "SELECT COUNT(*) FROM `timers` WHERE `organization_id` = ? AND `user_id` = ? AND `end_at` IS NULL"
	var count int
	if err := tx.QueryRowContext(ctx, query, organizationID, userID).Scan(&count); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return count > 0, nil
}
