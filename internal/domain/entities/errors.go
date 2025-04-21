package entities

import "errors"

var (
	// 共通エラー
	ErrNotFound              = errors.New("not found")
	ErrInvalidRequest        = errors.New("invalid request")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrForbidden             = errors.New("forbidden")
	ErrInternalServer        = errors.New("internal server error")
	ErrStartDateMustBeBefore = errors.New("start date must be before end date")
	ErrNoContentUpdated      = errors.New("no content updated")

	// 認証関連
	ErrMissingAuthorizationHeader = errors.New("missing Authorization header")
	ErrInvalidTokenFormat         = errors.New("invalid token format")
	ErrInvalidToken               = errors.New("invalid token")
	ErrInvalidOrExpiredToken      = errors.New("invalid or expired token")
	ErrUserNotFoundInRedis        = errors.New("user not found")
	ErrTokenNotFoundInRedis       = errors.New("token not found")
	ErrPasswordNotMatch           = errors.New("password not match")
	ErrCouldNotGenerateToken      = errors.New("could not generate token")

	// ヘルパー関連
	// コンテキストヘルパー
	ErrUserIdNotFoundInContext         = errors.New("user id not found in context")
	ErrOrganizationIdNotFoundInContext = errors.New("organization id not found in context")
	ErrRoleNotFoundInContext           = errors.New("role not found in context")

	// 組織関連
	ErrOrganizationNotFound      = errors.New("organization not found")
	ErrOrganizationAlreadyExists = errors.New("organization already exists")

	// ユーザー関連
	ErrUserNotFound           = errors.New("user not found")
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrCannotUpdateOtherUsers = errors.New("only administrators can update other users")
	ErrCannotUpdateRole       = errors.New("only administrators can update role")
	ErrCannotDeleteMyself     = errors.New("cannot delete myself")

	// 顧客関連
	ErrCustomerNotFound = errors.New("customer not found")
	ErrCustomerExists   = errors.New("customer already exists")

	// プロジェクト関連
	ErrProjectNotFound               = errors.New("project not found")
	ErrProjectAlreadyExists          = errors.New("project already exists")
	ErrCouldNotParseProjectStartDate = errors.New("could not parse project start date")
	ErrCouldNotParseProjectEndDate   = errors.New("could not parse project end date")

	// 予算関連
	ErrBudgetNotFound      = errors.New("budget not found")
	ErrBudgetAlreadyExists = errors.New("budget already exists")

	// タイマー関連
	ErrTimerNotFound       = errors.New("timer not found")
	ErrTimerAlreadyRunning = errors.New("timer already running")
	ErrTimerAlreadyStopped = errors.New("timer already Stopped")
)
