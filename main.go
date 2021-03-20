package main

import (
	"github.com/PrathameshKesarkar/banking-app/app"
	"github.com/PrathameshKesarkar/banking-app/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
