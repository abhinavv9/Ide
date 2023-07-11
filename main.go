package main

import (
	// "context"
	"fmt"
	// "log"
	// "time"

	"github.com/abhinavv9/codee/cmd"
	// "github.com/abhinavv9/codee/internal/container"
	// "github.com/docker/docker/client"
)

func main() {
	fmt.Println("Starting server...")
	cmd.Start()
}

// func main() {
// 	startTime := time.Now()

// 	ctx := context.Background()

// 	// Create a Docker client
// 	cli, err := client.NewClientWithOpts(client.FromEnv)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	image := "chat-py" // Replace with the actual image name

// 	container.SpinContainer(ctx, cli, image)

// 	// Calculate the total time taken
// 	totalTime := time.Since(startTime)
// 	log.Printf("Total time taken: %s", totalTime)

// 	// Get container memory usage

// }
