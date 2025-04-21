package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"
	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"
)

type UserRepository struct{}

var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

// Create はユーザー情報をDBに保存する
func (r *UserRepository) Create(ctx context.Context, tx *sql.Tx, user *entities.User) (*entities.UserID, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "INSERT INTO `users` (`organization_id`, `name`, `email`, `password`, `role`, `created_at`, `updated_at`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(
		ctx,
		query,
		organizationID,
		user.Name,
		user.Email,
		user.HashedPassword,
		user.Role,
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
	userID := entities.UserID(lastInsertID)

	return &userID, nil
}

// List はユーザー情報一覧をDBから取得する
func (r *UserRepository) List(ctx context.Context, tx *sql.Tx) ([]entities.User, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	var users []entities.User

	query := "SELECT `id`, `name`, `email`, `role` " +
		"FROM `users` " +
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
		var user entities.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Role,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Find はユーザー情報をDBから取得する
func (r *UserRepository) Find(ctx context.Context, tx *sql.Tx, userID *entities.UserID) (*entities.User, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	var user entities.User

	query := "SELECT `id`, `name`, `email`, `role` " +
		"FROM `users` " +
		"WHERE `id` = ? AND `organization_id` = ? " +
		"AND deleted_at IS NULL"
	result := tx.QueryRowContext(ctx, query, userID, organizationID)
	err := result.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// FindWithOrganizationID はユーザーIDと組織IDからユーザー情報を取得する
func (r *UserRepository) FindWithOrganizationID(ctx context.Context, tx *sql.Tx, userID *entities.UserID, organizationID *entities.OrganizationID) (*entities.User, error) {
	var user entities.User

	query := "SELECT `id`, `name`, `email`, `role` " +
		"FROM `users` " +
		"WHERE `id` = ? AND `organization_id` = ? " +
		"AND deleted_at IS NULL"
	result := tx.QueryRowContext(ctx, query, userID, organizationID)
	err := result.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// FindByEmail はEmailに紐づくユーザー情報をDBから取得する
func (r *UserRepository) FindByEmail(ctx context.Context, tx *sql.Tx, email *entities.UserEmail, organizationID *entities.OrganizationID) (*entities.User, error) {
	var user entities.User

	query := "SELECT `id`, `name`, `email`, `password`, `role` " +
		"FROM `users` " +
		"WHERE `email` = ? AND `organization_id` = ? " +
		"AND deleted_at IS NULL"
	result := tx.QueryRowContext(ctx, query, email, organizationID)
	err := result.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// Update はDBのユーザー情報を更新を行う
func (r *UserRepository) Update(ctx context.Context, tx *sql.Tx, user *entities.User) (*entities.UserID, error) {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "UPDATE `users` SET `name` = ?, `email` = ?, `role` = ? " +
		"WHERE `id` = ? AND `organization_id` = ? " +
		"AND deleted_at IS NULL"
	_, err := tx.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Role,
		user.ID,
		organizationID,
	)
	if err != nil {
		return nil, err
	}

	return &user.ID, nil
}

// SoftDelete はDBのユーザー情報を論理削除を行う
func (r *UserRepository) SoftDelete(ctx context.Context, tx *sql.Tx, userID *entities.UserID) error {
	organizationID := auth_context.ContextOrganizationID(ctx)

	query := "UPDATE `users` SET `deleted_at` = ? " +
		"WHERE `id` = ? AND `organization_id` = ? "
	_, err := tx.ExecContext(
		ctx,
		query,
		time.Now(),
		userID,
		organizationID,
	)
	if err != nil {
		return err
	}

	return nil
}
