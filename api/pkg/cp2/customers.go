package cp2

type Customer struct {
	Id      int    `json:"id"`
	Name    string `json:"fullName" xml:"fullName"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipCode" xml:"zipCode"`
}

// to mimic a connection to a DB
func generateCustomers() []Customer {
	return []Customer{
		{
			Id:      1,
			Name:    "Steve",
			City:    "Los Angeles",
			Zipcode: "91505",
		},
		{
			Id:      2,
			Name:    "Roland",
			City:    "Boarderlands",
			Zipcode: "00000",
		},
		{
			Id:      3,
			Name:    "Firehawk",
			City:    "badlands",
			Zipcode: "00001",
		},
	}

}
