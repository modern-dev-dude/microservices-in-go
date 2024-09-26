package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/modern-dev-dude/microservices-in-go/pkg/domain"
	"github.com/modern-dev-dude/microservices-in-go/pkg/service"
)

func Start() {
	// setup DB
	dbClient, err := connectToDb()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	customerRepoDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepoDb := domain.NewAccountRepositoryDb(dbClient)

	customerHandlers := CustomerHandlers{service.NewCustomerService(customerRepoDb)}
	accountHandler := AccountHandler{service.NewAccountService(accountRepoDb)}

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
