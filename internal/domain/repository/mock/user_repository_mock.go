package mock

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/domain/repository"
)

var _ repository.UserRepository = (*UserRepositoryMock)(nil)

type UserRepositoryMock struct {
	CreateFunc      func(ctx context.Context, tx *sql.Tx, user *entities.User) (*entities.UserID, error)
	ListFunc        func(ctx context.Context, tx *sql.Tx) ([]entities.User, error)
	FindFunc        func(ctx context.Context, tx *sql.Tx, userID *entities.UserID) (*entities.User, error)
	FindByEmailFunc func(ctx context.Context, tx *sql.Tx, email *entities.UserEmail) (*entities.User, error)
	UpdateFunc      func(ctx context.Context, tx *sql.Tx, user *entities.User) (*entities.UserID, error)
	SoftDeleteFunc  func(ctx context.Context, tx *sql.Tx, userID *entities.UserID) error
}

func (m *UserRepositoryMock) Create(ctx context.Context, tx *sql.Tx, user *entities.User) (*entities.UserID, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, tx, user)
	}
	return nil, nil
}

func (m *UserRepositoryMock) List(ctx context.Context, tx *sql.Tx) ([]entities.User, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx, tx)
	}
	return nil, nil
}

func (m *UserRepositoryMock) Find(ctx context.Context, tx *sql.Tx, userID *entities.UserID) (*entities.User, error) {
	if m.FindFunc != nil {
		return m.FindFunc(ctx, tx, userID)
	}
	return nil, nil
}

func (m *UserRepositoryMock) FindByEmail(ctx context.Context, tx *sql.Tx, email *entities.UserEmail) (*entities.User, error) {
	if m.FindByEmailFunc != nil {
		return m.FindByEmailFunc(ctx, tx, email)
	}
	return nil, nil
}

func (m *UserRepositoryMock) Update(ctx context.Context, tx *sql.Tx, user *entities.User) (*entities.UserID, error) {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, tx, user)
	}
	return nil, nil
}

func (m *UserRepositoryMock) SoftDelete(ctx context.Context, tx *sql.Tx, userID *entities.UserID) error {
	if m.SoftDeleteFunc != nil {
		return m.SoftDeleteFunc(ctx, tx, userID)
	}
	return nil
}
