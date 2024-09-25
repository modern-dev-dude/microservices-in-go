package main

import (
	"github.com/joho/godotenv"
	"github.com/modern-dev-dude/microservices-in-go/pkg/app"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"
)

func main() {
	logger.Info("Load env variables for development")
	err := godotenv.Load(".env")
	if err != nil {
		logger.CustomError("Error loading .env file")
	}

	logger.Info("Starting application")
	app.Start()
}
