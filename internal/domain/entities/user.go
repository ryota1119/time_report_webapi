package entities

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// UserID ユーザーID
type UserID uint

// String は UserID を string 型にキャストする
func (uid UserID) String() string {
	return strconv.FormatUint(uint64(uid), 10)
}

// Int は UserID を int 型にキャストする
func (uid UserID) Int() int {
	return int(uid)
}

// UserName ユーザーネーム
type UserName string

// String は UserName を string 型にキャストする
func (n UserName) String() string {
	return string(n)
}

// UserEmail ユーザーEメールアドレス
type UserEmail string

// String は UserEmail を string 型にキャストする
func (e UserEmail) String() string {
	return string(e)
}

// HashedPassword ハッシュパスワード
type HashedPassword string

// CheckHashedPassword は二つのパスワードを比較する
func (p HashedPassword) CheckHashedPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
}

// hashPassword は指定された文字列をハッシュ化し HashedPassword を生成する
func hashPassword(password string) (*HashedPassword, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashedPassword := HashedPassword(hashed)
	return &hashedPassword, nil
}

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

// String は Role を string 型にキャストする
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

// CachedUser Redis用のキャッシュユーザー
type CachedUser struct {
	ID    UserID    `json:"id"`
	Name  UserName  `json:"name"`
	Email UserEmail `json:"email"`
	Role  Role      `json:"role"`
}
