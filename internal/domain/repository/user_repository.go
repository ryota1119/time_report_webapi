package repository

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// UserRepository UserRepositoryのインターフェースを定義
type UserRepository interface {
	// Create はユーザー情報を作成する
	Create(ctx context.Context, tx *sql.Tx, user *entities.User) (*entities.UserID, error)
	// List はユーザーの一覧を取得する
	List(ctx context.Context, tx *sql.Tx) ([]entities.User, error)
	// Find はユーザーIDからユーザー情報を取得する
	Find(ctx context.Context, tx *sql.Tx, userID *entities.UserID) (*entities.User, error)
	// FindWithOrganizationID はユーザーIDと組織IDからユーザー情報を取得する
	FindWithOrganizationID(ctx context.Context, tx *sql.Tx, userID *entities.UserID, organizationID *entities.OrganizationID) (*entities.User, error)
	// FindByEmail はEmailからユーザー情報を取得する
	FindByEmail(ctx context.Context, tx *sql.Tx, email *entities.UserEmail, organizationID *entities.OrganizationID) (*entities.User, error)
	// Update はユーザーIDからユーザー情報を取得する
	Update(ctx context.Context, tx *sql.Tx, user *entities.User) (*entities.UserID, error)
	// SoftDelete はユーザー情報を論理削除する
	SoftDelete(ctx context.Context, tx *sql.Tx, userID *entities.UserID) error
}
