package entities

import "errors"

var (
	// 共通エラー
	ErrNotFound       = errors.New("not found")
	ErrInvalidRequest = errors.New("invalid request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrInternalServer = errors.New("internal server error")

	// 認証関連
	ErrMissingAuthorizationHeader = errors.New("missing Authorization header")
	ErrInvalidTokenFormat         = errors.New("invalid token format")
	ErrInvalidToken               = errors.New("invalid token")
	ErrInvalidOrExpiredToken      = errors.New("invalid or expired token")
	ErrUserNotFoundInRedis        = errors.New("user not found")
	ErrTokenNotFoundInRedis       = errors.New("token not found")

	// ヘルパー関連
	// コンテキストヘルパー
	ErrUserIdNotFoundInContext         = errors.New("user id not found in context")
	ErrOrganizationIdNotFoundInContext = errors.New("organization id not found in context")
	ErrRoleNotFoundInContext           = errors.New("role not found in context")

	// 組織関連
	ErrOrganizationNotFound = errors.New("user not found")
	ErrOrganizationExists   = errors.New("user already exists")

	// ユーザー関連
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")

	// 顧客関連
	ErrCustomerNotFound = errors.New("customer not found")
	ErrCustomerExists   = errors.New("customer already exists")

	// プロジェクト関連
	ErrProjectNotFound = errors.New("project not found")
	ErrProjectExists   = errors.New("project already exists")

	// タイマー関連
	ErrTimerAlreadyRunning = errors.New("timer already running")
)
