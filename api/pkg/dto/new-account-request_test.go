package dto

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func setupNewAccountRequest(accountType string, amount float64) NewAccountRequest {
	return NewAccountRequest{
		CustomerId:  "123",
		AccountType: accountType,
		Amount:      amount,
	}
}

func TestValidNewAccountRequest_Validate_WithValidRequest(t *testing.T) {
	data := setupNewAccountRequest("checking", 5000)

	err := NewAccountRequest.Validate(data)

	assert.Nil(t, err)
}

func TestValidNewAccountRequest_Validate_InvalidAccountType(t *testing.T) {
	data := setupNewAccountRequest("not valid", 5000.25)

	err := NewAccountRequest.Validate(data)

	assert.Equal(t, err.Message, "account type must be checking or saving")
	assert.Equal(t, err.Code, http.StatusUnprocessableEntity)
}

func TestValidNewAccountRequest_Validate_LackOfRequiredFunds(t *testing.T) {
	data := setupNewAccountRequest("checking", 654.52)

	err := NewAccountRequest.Validate(data)

	assert.Equal(t, err.Message, "you must have more than 5,000 to open a new account")
	assert.Equal(t, err.Code, http.StatusUnprocessableEntity)
}
