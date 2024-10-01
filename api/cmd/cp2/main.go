package main

import (
	"github.com/modern-dev-dude/microservices-in-go/api/pkg/cp2"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", cp2.Greet)
	mux.HandleFunc("/customers", cp2.GetAllCustomers)
	mux.HandleFunc("/customers/{cusomter_id}", cp2.GetCustomer)
	mux.HandleFunc("/customers/new", cp2.AddCustomer)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
