package cp2

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
)

func Greet(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "hello")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := generateCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	customers := generateCustomers()

	if customerId, err := strconv.Atoi(r.PathValue("cusomter_id")); err == nil {
		fmt.Printf("Customer id %v\n", customerId)
		for _, customer := range customers {
			if customer.Id == customerId {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(customer)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - customer doesn't exist"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - an error occured"))
	}

}
