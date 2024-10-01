package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/errs"
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/logger"
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
	// get current balance
	var account Account
	sqlAccount := `SELECT account_id, amount FROM accounts WHERE customer_id = ?`
	row := d.client.QueryRow(sqlAccount, t.CustomerId)

	err := row.Scan(&account.AccountId, &account.Amount)
	if err != nil {
		logger.CustomError("Error getting account when creating a transaction: " + err.Error())
		return nil, errs.NewInternalServerError("Error fetching account while creating new transaction")
	}

	// validate amount then update
	if t.Amount > account.Amount && t.TransactionType == "withdraw" {
		logger.CustomError("Insufficient funds: ")
		return nil, errs.NewInternalServerError("Insufficient funds")
	}

	updatedBalance := t.Amount
	if t.TransactionType == "withdraw" {
		updatedBalance = account.Amount - t.Amount
		logger.Info("updated balance " + strconv.FormatFloat(updatedBalance, 'f', -1, 64))
	}

	if t.TransactionType == "deposit" {
		updatedBalance = account.Amount + t.Amount
	}

	sqlUpdateBalance := `UPDATE accounts SET amount = ? WHERE account_id = ?`
	result, err := d.client.Exec(sqlUpdateBalance, updatedBalance, account.AccountId)
	if err != nil {
		logger.CustomError("Error while updating account: " + err.Error())
		return nil, errs.NewInternalServerError("Error while creating new transaction")
	}

	sqlInsert := `INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)`
	result, err = d.client.Exec(sqlInsert, account.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

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
	// return updated amount
	t.Amount = updatedBalance

	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
