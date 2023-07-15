package cmd

import (
	"context"
	"log"
	"os"

	"github.com/abhinavv9/codee/middleware"
	"github.com/abhinavv9/codee/routes"
	"github.com/abhinavv9/codee/types"
	"github.com/docker/docker/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewServer() *Server {
	app := fiber.New()

	return &Server{
		App: app,
	}
}

func (s *Server) setupRoutes() {
	s.Use(logger.New())

	s.Post("/execute", middleware.JobMiddleware(), routes.CodeExecutionRoute)

}

var LogFile = "./logs/log.txt"

func (s *types.Server) Start() error {

	//Setting up docker client and context
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	// Set the docker client and context
	s.DockerClient = cli
	s.Ctx = ctx

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
	return s.Listen(":5000")
}
