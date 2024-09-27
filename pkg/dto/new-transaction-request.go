package dto

import "github.com/modern-dev-dude/microservices-in-go/pkg/errs"

type NewTransactionRequest struct {
	CustomerId      string  `json:"customer_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (req NewTransactionRequest) Validate() *errs.AppErr {
	if req.Amount <= 0 {
		return errs.NewValidationError("amount must be greater than zero")
	}

	if req.TransactionType != "withdraw" && req.TransactionType != "deposit" {
		return errs.NewValidationError("transaction type must be withdraw, deposit")
	}

	return nil
}
