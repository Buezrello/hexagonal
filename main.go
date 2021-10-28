package main

import (
	"example.com/hexagonal/app"
	"example.com/hexagonal/logger"
)

func main() {
	// log.Println("Starting our application.")
	logger.Info("Starting the application...")
	app.Start()
}
