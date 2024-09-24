package domain

import (
	"bytes"
	"database/sql"
	"text/template"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/modern-dev-dude/microservices-in-go/pkg/errs"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"

	_ "github.com/mattn/go-sqlite3"
)

var connectionString string = "./microservice-in-go.db"

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDb() (CustomerRepositoryDb, error) {
	db, err := sqlx.Open("sqlite3", connectionString)
	if err != nil {
		return CustomerRepositoryDb{}, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDb{db}, nil
}

func (reciever CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppErr) {
	if status == "active" {
		status = "1"
	}
	if status == "inactive" {
		status = "0"
	}

	customers := []Customer{}

	queryStr := generateQueryStringByStatus(status)
	err := reciever.db.Select(&customers, queryStr)
	if err != nil {
		logger.CustomError("Error while querying customer table" + err.Error())
		return nil, errs.NewInternalServerError("internal service error")
	}

	return customers, nil
}

func (reciever CustomerRepositoryDb) GetCustomerById(id string) (*Customer, *errs.AppErr) {
	var customer Customer
	err := reciever.db.Get(&customer, "select * from customers where customer_id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.CustomError("no rows found" + err.Error())
			return nil, errs.NewNotFoundError("not found")
		}
		logger.CustomError("Error while scanning customer with id " + err.Error())
		return nil, errs.NewInternalServerError("internal server error")
	}

	return &customer, nil
}

type QueryTemplate struct {
	Status string
}

func generateQueryStringByStatus(status string) string {
	tmpl, err := template.New("").Parse(`
		select * from customers 
		{{if .Status}} where status='{{.Status}}'{{end}}`)

	if err != nil {
		logger.CustomError("Error getting status " + err.Error())
	}

	data := QueryTemplate{status}

	var qsBuf bytes.Buffer
	err = tmpl.Execute(&qsBuf, data)
	if err != nil {
		logger.CustomError("Error executing status template " + err.Error())
	}

	return qsBuf.String()
}
