package cp2

type Customer struct {
	Name    string `json:"fullName" xml:"fullName"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipCode" xml:"zipCode"`
}

// to mimic a connection to a DB
func generateCustomers() []Customer {
	return []Customer{
		{
			Name:    "Steve",
			City:    "Los Angeles",
			Zipcode: "91505",
		},
		{
			Name:    "Roland",
			City:    "Boarderlands",
			Zipcode: "00000",
		},
		{
			Name:    "Firehawk",
			City:    "badlands",
			Zipcode: "00001",
		},
	}

}
