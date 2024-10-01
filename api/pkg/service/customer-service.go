package service

import (
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/domain"
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/dto"
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppErr)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppErr) {
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppErr) {
	c, err := s.repo.GetCustomerById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
