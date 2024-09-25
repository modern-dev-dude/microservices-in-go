package app

import (
	"fmt"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"
	"log"
	"net/http"
	"os"

	"github.com/modern-dev-dude/microservices-in-go/pkg/domain"
	"github.com/modern-dev-dude/microservices-in-go/pkg/service"
)

func Start() {
	mux := http.NewServeMux()

	// customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	domain, err := domain.NewCustomerRepositoryDb()
	if err != nil {
		log.Fatalf("Error setting up connection to Db err:%v\n", err)
	}
	customerHandlers := CustomerHandlers{service.NewCustomerService(domain)}

	mux.HandleFunc("/customers", customerHandlers.getAllCustomersHandler)
	mux.HandleFunc("/customers/{id}", customerHandlers.getCustomerHandler)

	// get env variables
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	logger.Info(fmt.Sprintf("Starting server on %s:%s", host, port))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), mux))
}
