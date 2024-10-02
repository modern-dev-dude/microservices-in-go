package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/modern-dev-dude/microservices-in-go/auth_server/pkg/domain"
	"github.com/modern-dev-dude/microservices-in-go/auth_server/pkg/service"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	checkEnvVars()
	router := mux.NewRouter()

	authRepo := domain.NewAuthRepository(connectToDb())
	ah := AuthHandler{service.NewLoginService(authRepo, domain.GetRolePermissions())}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/refresh", ah.Refresh).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	serverAddress := fmt.Sprintf("%s:%s", address, port)

	log.Println("Listening on " + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func connectToDb() *sqlx.DB {
	DbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := sqlx.Open("sqlite3", DbConnectionString)
	if err != nil {
		log.Fatal("Error connecting to DB")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func checkEnvVars() {
	envProps := []string{
		"SERVER_ADDRESS", "SERVER_PORT", "DB_CONNECTION_STRING",
	}

	for _, prop := range envProps {
		if os.Getenv(prop) == "" {
			log.Fatal("Environment variable " + prop + " not set")
		}
	}
}
