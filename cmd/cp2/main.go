package main

import (
	"log"
	"net/http"

	"github.com/modern-dev-dude/microservices-in-go/pkg/cp2"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", cp2.Greet)
	mux.HandleFunc("/customers", cp2.GetAllCustomers)
	mux.HandleFunc("/customers/{cusomter_id}", cp2.GetCustomer)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
