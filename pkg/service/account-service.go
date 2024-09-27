package service

import (
	"github.com/modern-dev-dude/microservices-in-go/pkg/domain"
	"github.com/modern-dev-dude/microservices-in-go/pkg/dto"
	"github.com/modern-dev-dude/microservices-in-go/pkg/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppErr)
	NewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppErr)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppErr) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:00:00"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	res := newAccount.ToNewAccountResponseDto()
	return &res, nil
}

func (s DefaultAccountService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppErr) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	t := domain.Transaction{
		TransactionId:   "",
		Amount:          req.Amount,
		CustomerId:      req.CustomerId,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:00:00"),
	}

	transaction, err := s.repo.NewTransaction(t)
	if err != nil {
		return nil, err
	}

	res := transaction.ToNewTransactionResponseDto()
	return &res, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
