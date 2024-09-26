package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/modern-dev-dude/microservices-in-go/pkg/errs"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppErr) {
	sqlInsert := `INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)`

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.CustomError("Error while creating new account: " + err.Error())
		return nil, errs.NewInternalServerError("Error while creating new account")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.CustomError("Error while getting new account id: " + err.Error())
		return nil, errs.NewInternalServerError("Error while retrieving account id")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
