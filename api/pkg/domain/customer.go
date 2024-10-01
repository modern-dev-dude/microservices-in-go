package domain

import (
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/dto"
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppErr)
	GetCustomerById(string) (*Customer, *errs.AppErr)
}

func (c Customer) getStatusAsText() string {
	statusTxt := "inactive"
	if c.Status == "1" {
		statusTxt = "active"
	}

	return statusTxt
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.getStatusAsText(),
	}
}
