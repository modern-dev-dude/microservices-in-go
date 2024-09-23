package errs

import "net/http"

type AppErr struct {
	Code    int
	Message string
}

func NewNotFoundError(msg string) *AppErr {
	return &AppErr{
		Message: msg,
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(msg string) *AppErr {
	return &AppErr{
		Message: msg,
		Code:    http.StatusInternalServerError,
	}
}
