package cp2

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "hello")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := generateCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
