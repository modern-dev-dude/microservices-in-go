package app

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"
	"github.com/modern-dev-dude/microservices-in-go/pkg/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomersHandler(w http.ResponseWriter, r *http.Request) {
	reqId := generateReqId()
	logger.WriteLogToConsole(r, reqId)

	err := isNotCorrectMethod(w, r, "GET")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	customers, err := ch.service.GetAllCustomers()
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

func (ch *CustomerHandlers) getCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// if customerId, err := strconv.Atoi(r.PathValue("cusomter_id")); err == nil
	customerId := r.PathValue("id")
	// check if id is an int
	_, err := strconv.Atoi(customerId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not found")
		// dump error to a server log if this was in prod
		log.Printf("customer id is not of type int customer id: %v\n", customerId)
		return
	}

	customer, errs := ch.service.GetCustomer(customerId)
	if errs != nil {
		w.WriteHeader(errs.Code)
		fmt.Fprintf(w, errs.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
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
