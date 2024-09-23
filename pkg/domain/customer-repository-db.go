package domain

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var connectionString string = "./microservice-in-go.db"

type CustomerRepositoryDb struct {
	db *sql.DB
}

func (reciever CustomerRepositoryDb) FindAll() ([]Customer, error) {
	rows, err := reciever.db.Query("select * from customers")
	if err != nil {
		fmt.Printf("Error while querying customer table err:%v\n", err)
		return nil, err
	}

	customers := []Customer{}
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)

		if err != nil {
			fmt.Printf("Error while scanning customers err:%v\n", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
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
