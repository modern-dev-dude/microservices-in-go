package dto

import (
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

const CHECKING = "checking"
const SAVING = "checking"

func (req NewAccountRequest) IsChecking() bool {
	return req.AccountType == CHECKING
}

func (req NewAccountRequest) IsSavings() bool {
	return req.AccountType == SAVING
}

func (req NewAccountRequest) Validate() *errs.AppErr {
	if req.Amount < 5000 {
		return errs.NewValidationError("you must have more than 5,000 to open a new account")
	}

	if !req.IsSavings() && !req.IsChecking() {
		return errs.NewValidationError("account type must be checking or saving")
	}

	return nil
}
