package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	domainerrors "github.com/ryota1119/time_resport/internal/domain/errors"
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

// Create はタイマー情報を作成する
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

// Find はタイマー情報を取得する
func (r *TimerRepository) Find(ctx context.Context, tx *sql.Tx, timerID *entities.TimerID) (*entities.Timer, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)
	userID := auth_context.ContextUserID(ctx)

	var timer entities.Timer

	query := "SELECT `id`, `user_id`, `project_id`, `title`, `memo`, `start_at`, `end_at` " +
		"FROM `timers` " +
		"WHERE `id` = ? AND `organization_id` = ? AND `user_id` = ? "
	result := tx.QueryRowContext(ctx, query, timerID, organizationID, userID)
	err := result.Scan(
		&timer.ID,
		&timer.UserID,
		&timer.ProjectID,
		&timer.Title,
		&timer.Memo,
		&timer.StartAt,
		&timer.EndAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainerrors.ErrTimerNotFound
		}
		return nil, err
	}
	return &timer, nil
}

// Update はタイマー情報を更新する
func (r *TimerRepository) Update(ctx context.Context, tx *sql.Tx, timer *entities.Timer) (*entities.TimerID, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)
	userID := auth_context.ContextUserID(ctx)

	var endAt driver.Value
	if timer.EndAt != nil {
		endAt = timer.EndAt.Value()
	} else {
		endAt = nil
	}

	query := "UPDATE `timers` " +
		"SET `project_id` = ?, `title` = ?, `memo` = ?, `start_at` = ?, `end_at` = ?, `updated_at` = ? " +
		"WHERE `id` = ? AND `organization_id` = ? AND `user_id` = ? "
	result, err := tx.ExecContext(
		ctx,
		query,
		timer.ProjectID,
		timer.Title,
		timer.Memo,
		timer.StartAt.Value(),
		endAt,
		time.Now(),
		timer.ID,
		organizationID,
		userID,
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

// ExistsRunningTimer はユーザーの稼働中のタイマーがあるかどうかチェックする
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
