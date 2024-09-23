package cp2

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"
)

// place customers in memory for testing
var customers = generateCustomers()

func Greet(w http.ResponseWriter, r *http.Request) {
	reqId := generateReqId()
	logger.WriteLogToConsole(r, reqId)
	fmt.Fprint(w, "hello")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	reqId := generateReqId()
	logger.WriteLogToConsole(r, reqId)

	err := isNotCorrectMethod(w, r, "GET")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	reqId := generateReqId()
	logger.WriteLogToConsole(r, reqId)

	err := isNotCorrectMethod(w, r, "GET")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

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

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	reqId := generateReqId()
	logger.WriteLogToConsole(r, reqId)

	err := isNotCorrectMethod(w, r, "POST")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	data := Customer{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	customers = append(customers, Customer{
		// Today me is happy, but a later version will forget this when working with concurrent connections and will be sad
		Id:      customers[len(customers)-1].Id + 1,
		Name:    data.Name,
		City:    data.City,
		Zipcode: data.Zipcode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func isNotCorrectMethod(w http.ResponseWriter, r *http.Request, allowedMethod string) error {
	if r.Method != allowedMethod {
		w.Header().Set("Allow", allowedMethod)
		http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
		return errors.New("this method is not allowed")
	}
	return nil
}

func generateReqId() string {
	return uuid.New().String()
}
