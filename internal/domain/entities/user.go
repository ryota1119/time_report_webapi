package entities

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// UserID ユーザーID
type UserID uint

// String はUserIDをstring型にキャストする
func (uid UserID) String() string {
	return strconv.FormatUint(uint64(uid), 10)
}

// UserName ユーザーネーム
type UserName string

// UserEmail ユーザーEメールアドレス
type UserEmail string

// HashedPassword ハッシュパスワード
type HashedPassword string

// Role ユーザー権限
type Role string

const (
	// AdminRole 管理者権限
	AdminRole Role = "admin"
	// MemberRole メンバー権限
	MemberRole Role = "member"
	// UnknownRole 該当なし
	UnknownRole Role = "unknown"
)

// String はRoleをstring型にキャストする
func (r Role) String() string {
	return string(r)
}

// User ユーザー情報
type User struct {
	ID             UserID
	Name           UserName
	Email          UserEmail
	HashedPassword HashedPassword
	Role           Role
}

// NewUser は指定されたユーザー名とメールアドレス、パスワード、権限ロールから User を生成する
// パスワードは文字列型で受け取り、ハッシュ化した値を生成する
func NewUser(name, email, password, role string) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:           UserName(name),
		Email:          UserEmail(email),
		HashedPassword: *hashedPassword,
		Role:           Role(role),
	}, nil
}

// HashPassword は指定された文字列をハッシュ化し Password を生成する
func hashPassword(password string) (*HashedPassword, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashedPassword := HashedPassword(hashed)
	return &hashedPassword, nil
}

// CheckHashedPassword は二つのパスワードを比較する
func (p HashedPassword) CheckHashedPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
}

// CachedUser Redis用のキャッシュユーザー
type CachedUser struct {
	ID    UserID    `json:"id"`
	Name  UserName  `json:"name"`
	Email UserEmail `json:"email"`
	Role  Role      `json:"role"`
}
