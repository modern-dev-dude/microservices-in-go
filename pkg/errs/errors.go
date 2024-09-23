package errs

import "net/http"

type AppErr struct {
	Code    int `json:",omitempty"`
	Message string
}

func (e AppErr) AsMessage() *AppErr {
	return &AppErr{
		Message: e.Message,
	}
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
