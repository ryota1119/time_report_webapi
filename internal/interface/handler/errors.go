package handler

import "net/http"

type HandlerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HandlerError) Error() string {
	return e.Message
}

var (
	ErrBadRequest = &HandlerError{
		Code:    http.StatusBadRequest,
		Message: "bad request",
	}
	ErrUnauthorized = &HandlerError{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
	}
	ErrForbidden = &HandlerError{
		Code:    http.StatusForbidden,
		Message: "forbidden",
	}
	ErrNotFound = &HandlerError{
		Code:    http.StatusNotFound,
		Message: "resource not found",
	}
	ErrConflict = &HandlerError{
		Code:    http.StatusConflict,
		Message: "conflict occurred",
	}
	ErrInternalServer = &HandlerError{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}
)
