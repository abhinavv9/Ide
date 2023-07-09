package main

import (
	"context"
	"log"
	"time"

	"github.com/abinavv9/codee/internal"
	"github.com/docker/docker/client"
)

func main() {
	startTime := time.Now()

	ctx := context.Background()

	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	image := "chat-py" // Replace with the actual image name

	internal.SpinContainer(ctx, cli, image)

	// Calculate the total time taken
	totalTime := time.Since(startTime)
	log.Printf("Total time taken: %s", totalTime)

	// Get container memory usage

}
