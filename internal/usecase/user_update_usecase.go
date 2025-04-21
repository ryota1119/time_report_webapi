package usecase

import (
	"context"
	"database/sql"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"
)

var _ UserUpdateUsecase = (*userUpdateUsecase)(nil)

// UserUpdateUsecase は usecase.userUpdateUsecase のインターフェースを定義
type UserUpdateUsecase interface {
	Update(ctx context.Context, input UserUpdateUsecaseInput) (*entities.User, error)
}

// userUpdateUsecase ユースケース
type userUpdateUsecase struct {
	db       *sql.DB
	userRepo repository.UserRepository
}

// NewUserUpdateUsecase は userUpdateUsecase を初期化する
func NewUserUpdateUsecase(
	db *sql.DB,
	userRepo repository.UserRepository,
) UserUpdateUsecase {
	return &userUpdateUsecase{
		db:       db,
		userRepo: userRepo,
	}
}

// UserUpdateUsecaseInput はuserUsecase.Updateのインプット
type UserUpdateUsecaseInput struct {
	UserID int
	Name   string
	Email  string
	Role   string
}

// Update はユーザー情報を更新する
func (a *userUpdateUsecase) Update(ctx context.Context, input UserUpdateUsecaseInput) (*entities.User, error) {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	userID := entities.UserID(input.UserID)
	// 既存のユーザー情報を取得する
	user, err := a.userRepo.Find(ctx, tx, &userID)
	if err != nil {
		return nil, err
	}

	// admin権限でなかった場合
	if auth_context.ContextUserRole(ctx) != entities.AdminRole {
		// 他のユーザーの情報は更新できない
		if auth_context.ContextUserID(ctx) != user.ID {
			return nil, entities.ErrCannotUpdateOtherUsers
		}
		// roleの更新はできない
		if entities.Role(input.Role) != user.Role {
			return nil, entities.ErrCannotUpdateRole
		}
	}

	isUpdat := false

	// 何も更新がない場合は、エラーを返却し、handler層でno contentを返す
	if user.Name == entities.UserName(input.Name) {
		user.Name = entities.UserName(input.Name)
		isUpdat = true
	}
	if user.Email == entities.UserEmail(input.Email) {
		user.Email = entities.UserEmail(input.Email)
		isUpdat = true
	}
	if user.Role == entities.Role(input.Role) {
		user.Role = entities.Role(input.Role)
		isUpdat = true
	}
	if !isUpdat {
		return nil, entities.ErrNoContentUpdated
	}

	// ユーザー情報を更新する
	_, err = a.userRepo.Update(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
