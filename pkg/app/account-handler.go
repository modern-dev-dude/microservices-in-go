package app

import (
	"encoding/json"
	"github.com/modern-dev-dude/microservices-in-go/pkg/dto"
	"github.com/modern-dev-dude/microservices-in-go/pkg/errs"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"
	"github.com/modern-dev-dude/microservices-in-go/pkg/service"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) newAccount(w http.ResponseWriter, r *http.Request) {
	reqId := GenerateReqId()
	logger.WriteLogToConsole(r, reqId)

	isCorrectMethod := IsCorrectMethod(w, r, "POST")
	//  return early logs and response written
	if isCorrectMethod == false {
		return
	}

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

	var req dto.NewAccountRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResposne(w, http.StatusBadRequest, err.Error(), _json)
	} else {
		req.CustomerId = customerId
		account, appErr := a.service.NewAccount(req)
		if appErr != nil {
			writeResposne(w, appErr.Code, appErr.Message, _json)
		} else {
			writeResposne(w, http.StatusCreated, account, _json)
		}
	}
}

func (a AccountHandler) createTransaction(w http.ResponseWriter, r *http.Request) {
	reqId := GenerateReqId()
	logger.WriteLogToConsole(r, reqId)

	isCorrectMethod := IsCorrectMethod(w, r, "POST")
	//  return early logs and response written
	if isCorrectMethod == false {
		return
	}

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

	var req dto.NewTransactionRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	req.CustomerId = customerId
	if err != nil {
		writeResposne(w, http.StatusBadRequest, err.Error(), _json)
	} else {
		transaction, appErr := a.service.NewTransaction(req)
		if appErr != nil {
			writeResposne(w, appErr.Code, appErr.Message, _json)
		} else {
			writeResposne(w, http.StatusCreated, transaction, _json)
		}
	}
}
