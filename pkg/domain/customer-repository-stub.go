package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "1",
			Name:        "Steve",
			City:        "Los Angeles",
			Zipcode:     "91505",
			DateOfBirth: "2000-01-01",
			Status:      "1",
		},
		{
			Id:          "2",
			Name:        "Roland",
			City:        "Boarderlands",
			Zipcode:     "00000",
			DateOfBirth: "2000-01-01",
			Status:      "1",
		},
		{
			Id:          "3",
			Name:        "Firehawk",
			City:        "badlands",
			Zipcode:     "00001",
			DateOfBirth: "2000-01-01",
			Status:      "2",
		},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
