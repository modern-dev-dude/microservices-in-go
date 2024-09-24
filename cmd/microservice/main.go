package main

import (
	"github.com/modern-dev-dude/microservices-in-go/pkg/app"
	"github.com/modern-dev-dude/microservices-in-go/pkg/logger"
)

func main() {
	logger.Info("Starting application")
	app.Start()
}
