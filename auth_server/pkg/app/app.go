package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Start() {
	router := mux.NewRouter()

	authRepo := domain.NewAuthRepository(getDbClient())
	ah := AuthHandler{service.NewLogingService(authRepo, domain.GetRolePermissions())}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	serverAddress := fmt.Sprintf("%s:%s", address, port)

	log.Println("Listening on " + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
