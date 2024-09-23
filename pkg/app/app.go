package app

import (
	"log"
	"net/http"

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

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
