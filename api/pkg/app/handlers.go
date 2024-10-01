package app

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/errs"
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/logger"
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/service"
	"log"
	"net/http"
	"strconv"
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
	reqId := GenerateReqId()
	logger.WriteLogToConsole(r, reqId)

	isCorrectMethod := IsCorrectMethod(w, r, "GET")
	//  return early logs and response written
	if isCorrectMethod == false {
		return
	}

	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeResposne(w, http.StatusNotFound, err.AsMessage(), _json)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		writeResposne(w, http.StatusNotFound, customers, _xml)
		return
	}

	writeResposne(w, http.StatusOK, customers, _json)

}

func (ch *CustomerHandlers) getCustomerHandler(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("id")
	// check if id is an int
	_, err := strconv.Atoi(customerId)
	if err != nil {
		errMsg := "customer id is not of type int customer id: " + customerId
		writeResposne(
			w,
			http.StatusNotFound,
			errs.NewNotFoundError(errMsg).AsMessage(),
			_json)
		return
	}

	customer, appErr := ch.service.GetCustomer(customerId)
	if appErr != nil {
		writeResposne(w, http.StatusNotFound, appErr.AsMessage(), _json)
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
