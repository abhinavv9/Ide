package cmd

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Use(logger.New())

	// Serve static files from the "public" directory
	app.Static("/", "./public")
}

var LogFile = "./logs/log.txt"

func Start() {
	// Open the log file
	logFile, err := os.OpenFile(LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		// Set the output file for the logs
		Output: logFile,
	}))

	setupRoutes(app)

	// Start the server on port 3000
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
