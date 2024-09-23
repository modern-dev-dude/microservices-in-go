package domain

import "github.com/modern-dev-dude/microservices-in-go/pkg/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	GetCustomerById(string) (*Customer, *errs.AppErr)
}
