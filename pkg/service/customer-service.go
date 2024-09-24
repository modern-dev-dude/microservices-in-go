package service

import (
	"github.com/modern-dev-dude/microservices-in-go/pkg/domain"
	"github.com/modern-dev-dude/microservices-in-go/pkg/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppErr)
	GetCustomer(string) (*domain.Customer, *errs.AppErr)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppErr) {
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppErr) {
	return s.repo.GetCustomerById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
