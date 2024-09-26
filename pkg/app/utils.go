package app

import (
	"github.com/google/uuid"
	"github.com/modern-dev-dude/microservices-in-go/pkg/errs"
	"net/http"
)

func IsCorrectMethod(w http.ResponseWriter, r *http.Request, allowedMethod string) bool {
	if r.Method != allowedMethod {
		err := errs.AppErr{Message: "method not supported", Code: http.StatusMethodNotAllowed}
		writeResposne(w, err.Code, err.AsMessage(), _json)
		return false
	}
	return true
}

func GenerateReqId() string {
	return uuid.New().String()
}
