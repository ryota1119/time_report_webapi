package usecase

import (
	"context"
	"database/sql"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/domain/repository"
)

// UserUsecase UserUsecaseのインターフェースを定義
type UserUsecase interface {
	// Create はユーザーを新規作成する
	Create(ctx context.Context, input UserUsecaseCreateInput) (*entities.User, error)
	// List はユーザーの一覧を取得する
	List(ctx context.Context) ([]entities.User, error)
	// Get はユーザー情報を取得する
	Get(ctx context.Context, input UserUsecaseGetInput) (*entities.User, error)
	// Update はユーザー情報を更新する
	Update(ctx context.Context, input UserUsecaseUpdateInput) (*entities.User, error)
	// SoftDelete はユーザー情報を論理削除する
	SoftDelete(ctx context.Context, input UserUsecaseSoftDeleteInput) error
}

type userUsecase struct {
	db       *sql.DB
	userRepo repository.UserRepository
}

var _ UserUsecase = (*userUsecase)(nil)

func NewUserUsecase(
	db *sql.DB,
	userRepo repository.UserRepository,
) UserUsecase {
	return &userUsecase{
		db:       db,
		userRepo: userRepo,
	}
}
