package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	domain2 "github.com/modern-dev-dude/microservices-in-go/api/pkg/domain"
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/logger"
	service2 "github.com/modern-dev-dude/microservices-in-go/api/pkg/service"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	// setup DB
	dbClient, err := connectToDb()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	customerRepoDb := domain2.NewCustomerRepositoryDb(dbClient)
	accountRepoDb := domain2.NewAccountRepositoryDb(dbClient)

	customerHandlers := CustomerHandlers{service2.NewCustomerService(customerRepoDb)}
	accountHandler := AccountHandler{service2.NewAccountService(accountRepoDb)}

	mux.HandleFunc("/customers", customerHandlers.getAllCustomersHandler)
	mux.HandleFunc("/customers/{id}", customerHandlers.getCustomerHandler)
	mux.HandleFunc("/customers/{id}/account", accountHandler.newAccount)
	mux.HandleFunc("/customers/{id}/transaction", accountHandler.createTransaction)

	// get env variables
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	logger.Info(fmt.Sprintf("Starting server on %s:%s", host, port))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), mux))
}

func connectToDb() (*sqlx.DB, error) {
	DbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := sqlx.Open("sqlite3", DbConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
