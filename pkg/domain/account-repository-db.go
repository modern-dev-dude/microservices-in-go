package domain

import (
	"fmt"
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

func (d AccountRepositoryDb) NewTransaction(t Transaction) (*Transaction, *errs.AppErr) {

	sqlAccount := `SELECT amount FROM accounts WHERE account_id = ?`
	row := d.client.QueryRow(sqlAccount, t.AccountId)
	var account Account
	err := row.Scan(&account.Amount)
	if err != nil {
		logger.CustomError("Error getting account when creating a transaction: " + err.Error())
		return nil, errs.NewInternalServerError("Error fetching account while creating new transaction")
	}

	logger.Debug(fmt.Sprintf("Account ID %v", account.Amount))

	sqlInsert := `INSERT INTO accounts (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)`
	result, err := d.client.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if err != nil {
		logger.CustomError("Error while creating new transaction: " + err.Error())
		return nil, errs.NewInternalServerError("Error while creating new transaction")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.CustomError("Error while getting new transaction id: " + err.Error())
		return nil, errs.NewInternalServerError("Error while getting new transaction id")
	}

	t.TransactionId = strconv.FormatInt(id, 10)

	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
