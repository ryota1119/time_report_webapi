package usecase

import (
	"context"
	"errors"

	"github.com/ryota1119/time_resport/internal/domain/entities"
	"github.com/ryota1119/time_resport/internal/helper/auth_context"
)

// UserUsecaseUpdateInput はuserUsecase.Updateのインプット
type UserUsecaseUpdateInput struct {
	UserID int
	Name   string
	Email  string
	Role   string
}

// Update はユーザー情報を更新する
func (a *userUsecase) Update(ctx context.Context, input UserUsecaseUpdateInput) (*entities.User, error) {
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
			return nil, errors.New("user_id is not found")
		}
		// roleの更新はできない
		if entities.Role(input.Role) != user.Role {
			return nil, errors.New("role is not found")
		}
	}

	// 何も更新がない場合は、エラーを返却し、handler層でno contentを返す
	if user.Name == entities.UserName(input.Name) &&
		user.Email == entities.UserEmail(input.Email) &&
		user.Role == entities.Role(input.Role) {
		return nil, errors.New("name and email and role is the same")
	}

	// entities.User情報更新
	user.Name = entities.UserName(input.Name)
	user.Email = entities.UserEmail(input.Email)
	user.Role = entities.Role(input.Role)
	// ユーザー情報を更新する
	_, err = a.userRepo.Update(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
