package usecase

import (
	"context"
	"database/sql"
	"github.com/ryota1119/time_resport_webapi/internal/domain/repository"

	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"
)

var _ UserSoftDeleteUsecase = (*userSoftDeleteUsecase)(nil)

// UserSoftDeleteUsecase は usecase.userSoftDeleteUsecase のインターフェースを定義
type UserSoftDeleteUsecase interface {
	SoftDelete(ctx context.Context, input UserSoftDeleteUsecaseInput) error
}

// userSoftDeleteUsecase ユースケース
type userSoftDeleteUsecase struct {
	db       *sql.DB
	userRepo repository.UserRepository
}

// NewUserSoftDeleteUsecase は userSoftDeleteUsecase を初期化する
func NewUserSoftDeleteUsecase(
	db *sql.DB,
	userRepo repository.UserRepository,
) UserSoftDeleteUsecase {
	return &userSoftDeleteUsecase{
		db:       db,
		userRepo: userRepo,
	}
}

// UserSoftDeleteUsecaseInput はuserUsecase.SoftDeleteのインプット
type UserSoftDeleteUsecaseInput struct {
	UserID int
}

// SoftDelete はユーザー情報を論理削除する
func (a *userSoftDeleteUsecase) SoftDelete(ctx context.Context, input UserSoftDeleteUsecaseInput) error {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	userID := entities.UserID(input.UserID)
	// 自身のユーザー情報は削除できない
	if userID == auth_context.ContextUserID(ctx) {
		return entities.ErrCannotDeleteMyself
	}

	// 既存のユーザー情報を取得する
	_, err = a.userRepo.Find(ctx, tx, &userID)
	if err != nil {
		return err
	}

	// ユーザー情報を削除する
	err = a.userRepo.SoftDelete(ctx, tx, &userID)
	if err != nil {
		return err
	}
	return nil
}
