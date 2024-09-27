package domain

import (
	"github.com/modern-dev-dude/microservices-in-go/pkg/dto"
	"github.com/modern-dev-dude/microservices-in-go/pkg/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type Transaction struct {
	TransactionId   string
	CustomerId      string
	Amount          float64
	TransactionType string
	TransactionDate string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppErr)
	NewTransaction(Transaction) (*Transaction, *errs.AppErr)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

func (t Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:  t.TransactionId,
		CurrentBalance: t.Amount,
	}
}
