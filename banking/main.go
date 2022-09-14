package main

import (
	"banking-restapi/app"
	"banking-restapi/logger"
)

func main() {
	logger.Info("Starting the application....")
	app.Start()
}
