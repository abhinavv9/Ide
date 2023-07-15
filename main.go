package main

import (
	// "log"
	// "time"

	"github.com/abhinavv9/codee/cmd"
	// "github.com/abhinavv9/codee/internal/container"
	// "github.com/docker/docker/client"
)

func main() {
	server := cmd.NewServer()

	if err := server.Start(); err != nil {
		panic(err)
	}
}
