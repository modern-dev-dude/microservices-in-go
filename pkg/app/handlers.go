package app

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/modern-dev-dude/microservices-in-go/pkg/Logger"
	"github.com/modern-dev-dude/microservices-in-go/pkg/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomersHandler(w http.ResponseWriter, r *http.Request) {
	reqId := generateReqId()
	Logger.WriteLogToConsole(r, reqId)

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
