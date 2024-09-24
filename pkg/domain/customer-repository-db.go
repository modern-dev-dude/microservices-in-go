package domain

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/modern-dev-dude/microservices-in-go/pkg/errs"

	_ "github.com/mattn/go-sqlite3"
)

var connectionString string = "./microservice-in-go.db"

type CustomerRepositoryDb struct {
	db *sql.DB
}

func NewCustomerRepositoryDb() (CustomerRepositoryDb, error) {
	db, err := sql.Open("sqlite3", connectionString)
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

	queryStr := generateQueryStringByStatus(status)
	rows, err := reciever.db.Query(queryStr)
	if err != nil {
		fmt.Printf("Error while querying customer table err:%v\n", err)
		return nil, errs.NewInternalServerError("internal service error")
	}

	customers := []Customer{}
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)

		if err != nil {
			fmt.Printf("Error while scanning customers err:%v\n", err)
			return nil, errs.NewInternalServerError("internal service error")
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (reciever CustomerRepositoryDb) GetCustomerById(id string) (*Customer, *errs.AppErr) {
	row := reciever.db.QueryRow("select * from customers where customer_id = ?", id)

	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("not found")
		}
		log.Printf("Error while scanning customer with id: %v\n err:%v\n", id, err)
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
		fmt.Printf("Error getting status:%v\n", tmpl)

	}

	data := QueryTemplate{status}

	var qsBuf bytes.Buffer
	err = tmpl.Execute(&qsBuf, data)
	if err != nil {
		fmt.Printf("Error getting status:%v\n", tmpl)

	}

	return qsBuf.String()
}
