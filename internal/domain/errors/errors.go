package errors

type DomainError struct {
	Message string `json:"message"`
}

func (e *DomainError) Error() string {
	return e.Message
}

// 共通ドメインロジック関連エラー
var (
	ErrStartDateMustBeBefore = &DomainError{Message: "start date must be before end date"}
	ErrNoContentUpdated      = &DomainError{Message: "no content updated"}
)

// 認証のドメインロジック関連エラー
var (
	ErrPasswordNotMatch = &DomainError{Message: "password not match"}
)

// 組織のドメインロジック関連エラー
var (
	ErrOrganizationAlreadyExists = &DomainError{Message: "user already exists"}
	ErrOrganizationNotFound      = &DomainError{Message: "user not found"}
)

// ユーザーのドメインロジック関連エラー
var (
	ErrUserAlreadyExists = &DomainError{Message: "user already exists"}
	ErrUserNotFound      = &DomainError{Message: "user not found"}
)

// 顧客のドメインロジック関連エラー
var (
	ErrCustomerAlreadyExists = &DomainError{Message: "customer already exists"}
	ErrCustomerNotFound      = &DomainError{Message: "customer not found"}
)

// 予算のドメインロジック関連エラー
var (
	ErrBudgetAlreadyExists = &DomainError{Message: "budget already exists"}
	ErrBudgetNotFound      = &DomainError{Message: "budget not found"}
)
