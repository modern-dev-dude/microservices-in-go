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
	"github.com/modern-dev-dude/microservices-in-go/pkg/errs"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"
	"github.com/modern-dev-dude/microservices-in-go/pkg/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

// enum for  content types
const (
	_xml = iota
	_json
)

func getContentType(ct int) (string, error) {
	switch ct {
	case 0:
		return "application/xml", nil

	case 1:
		return "application/json", nil

	default:
		return "", errors.New("unsupported content type")
	}
}

func (ch *CustomerHandlers) getAllCustomersHandler(w http.ResponseWriter, r *http.Request) {
	reqId := generateReqId()
	logger.WriteLogToConsole(r, reqId)

	err := isNotCorrectMethod(w, r, "GET")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	customers, errs := ch.service.GetAllCustomers()
	if errs != nil {
		writeResposne(w, http.StatusNotFound, errs.AsMessage(), _json)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		writeResposne(w, http.StatusNotFound, customers, _xml)
		return
	}

	writeResposne(w, http.StatusOK, customers, _json)

}

func (ch *CustomerHandlers) getCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// if customerId, err := strconv.Atoi(r.PathValue("cusomter_id")); err == nil
	customerId := r.PathValue("id")
	// check if id is an int
	_, err := strconv.Atoi(customerId)
	if err != nil {
		writeResposne(w, http.StatusNotFound, errs.NewNotFoundError("not found").AsMessage(), _json)

		// dump error to a server
		log.Printf("customer id is not of type int customer id: %v\n", customerId)
		return
	}

	customer, errs := ch.service.GetCustomer(customerId)
	if errs != nil {
		writeResposne(w, http.StatusNotFound, errs.AsMessage(), _json)
		// dump error to server
		log.Printf("code: %v\nmessage: %v", errs.Code, errs.Message)
	} else {
		if r.Header.Get("Content-Type") == "application/xml" {
			writeResposne(w, http.StatusNotFound, customer, _xml)
			return
		}
		writeResposne(w, http.StatusOK, customer, _json)
	}
}

func writeResposne(w http.ResponseWriter, code int, data interface{}, ct int) {
	contentType, err := getContentType(ct)
	if err != nil {
		log.Println(err.Error())
	}

	w.Header().Add("Content-Type", contentType)
	w.WriteHeader(code)

	if ct == _xml {
		if err := xml.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	}

	if ct == _json {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
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
