package cmd

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	*fiber.App
}

func NewServer() *Server {
	app := fiber.New()

	return &Server{
		App: app,
	}
}

func (s *Server) setupRoutes() {
	s.Use(logger.New())

	// Define your routes here
	s.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
}

var LogFile = "./logs/log.txt"

func (s *Server) Start() error {
	// Open the log file
	logFile, err := os.OpenFile(LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	s.Use(logger.New(logger.Config{
		// Set the output file for the logs
		Output: logFile,
	}))
	s.setupRoutes()

	// Start the server on port 3000
	return s.Listen(":3000")
}
