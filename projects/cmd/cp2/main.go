package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"fullName" xml:"fullName"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipCode" xml:"zipCode"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
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

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := generateCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
