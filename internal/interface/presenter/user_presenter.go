package presenter

import (
	"github.com/ryota1119/time_resport/internal/domain/entities"
)

// UserResponse 顧客APIの基本レスポンス構造体
// ユーザー情報を返すための共通フォーマット
type UserResponse struct {
	ID    entities.UserID    `json:"id" example:"1"`
	Name  entities.UserName  `json:"name" example:"山田太郎"`
	Email entities.UserEmail `json:"email" example:"password"`
	Role  entities.Role      `json:"role" example:"admin"`
}

// UserCreateResponse Createアクションのレスポンス
// UserResponseを継承し、新規作成時に利用
type UserCreateResponse UserResponse

// NewUserCreateResponse domain.UserからUserCreateResponseを生成する
// ユーザー作成時のレスポンスとして使用
func NewUserCreateResponse(user *entities.User) UserCreateResponse {
	return UserCreateResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

// UserListResponse 複数のユーザー情報を格納するレスポンス
// UserResponseのスライス型
type UserListResponse []UserResponse

// NewUserListResponse domain.UserのスライスからUserListResponseを生成する
// ユーザーの一覧取得時に使用
func NewUserListResponse(users []entities.User) UserListResponse {
	output := make([]UserResponse, len(users))
	for i, user := range users {
		output[i] = UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		}
	}
	return output
}

// UserGetResponse 特定のユーザー情報を返すレスポンス
// UserResponseを継承し、ユーザー取得時に使用
type UserGetResponse UserResponse

// NewUserGetResponse domain.UserからUserGetResponseを生成する
// ユーザー情報取得時のレスポンスとして使用
func NewUserGetResponse(user *entities.User) UserGetResponse {
	return UserGetResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

// UserUpdateResponse ユーザー情報更新時のレスポンス
// UserResponseを継承し、更新後のデータを返す
type UserUpdateResponse UserResponse

// NewUserUpdateResponse domain.UserからUserUpdateResponseを生成する
// ユーザー更新時のレスポンスとして使用
func NewUserUpdateResponse(user *entities.User) UserUpdateResponse {
	return UserUpdateResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
