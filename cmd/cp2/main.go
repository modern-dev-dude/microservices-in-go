package main

import (
	"log"
	"net/http"

	"github.com/modern-dev-dude/microservices-in-go/pkg/cp2"
)

func main() {
	http.HandleFunc("/greet", cp2.Greet)
	http.HandleFunc("/customers", cp2.GetAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
