package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

var connectionString string = "./microservice-in-go.db"

func main() {
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		log.Fatalf("Error with DB connection: %v\n", err)
	}

	rows, err := db.Query(`select customer_id, name FROM customers`)
	if err != nil {
		log.Fatalf("Error with DB request for customers: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var customer_id string
		var name string
		err = rows.Scan(&customer_id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(customer_id, name)
	}
}
