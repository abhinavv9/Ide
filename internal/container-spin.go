package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	// "github.com/abinavv9/codee/scripts"
	"github.com/docker/docker/client"
)

type Job struct {
	UserID string
	Code   string
}

func SpinContainer(ctx context.Context, cli *client.Client, image string) {

	// Create a Docker client
	// cli, err := client.NewClientWithOpts(client.FromEnv)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Set up channels and wait group
	jobCh := make(chan Job)
	doneCh := make(chan struct{})
	var wg sync.WaitGroup

	// Define the number of workers (number of containers to run concurrently)
	numWorkers := 1

	// Start the workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for job := range jobCh {
				//Start the container
				containerID, err := runContainer(ctx, cli, job.Code, image)
				if err != nil {
					log.Printf("Error running container for user %s: %v", job.UserID, err)
					continue
				}

				//Execute the code passed thru env variable
				output, err := ExecuteCodeInContainer(ctx, cli, containerID)
				if err != nil {
					log.Fatal(err)
				}

				//handle output
				fmt.Println(string(output))

				// logs, err := getContainerLogs(ctx, cli, containerID)
				// if err != nil {
				// 	log.Printf("Error retrieving logs for user %s: %v", job.UserID, err)
				// 	continue
				// }

				// memUsage, err := scripts.GetContainerMemoryUsage(ctx, cli, containerID)
				// if err != nil {
				// 	log.Println("Failed to get container memory usage:", err)
				// 	return
				// }

				// Print container memory usage
				// log.Printf("Container memory used: %d bytes", memUsage)

				// log.Printf("Container logs for user %s:\n%s\n", job.UserID, logs)

				err = removeContainer(ctx, cli, containerID)
				if err != nil {
					log.Printf("Error removing container for user %s: %v", job.UserID, err)
					continue
				}
			}
		}()
	}

	// Start a goroutine to wait for completion
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	// Simulate multiple users submitting code jobs
	users := []string{"user1"} // Add more users as needed
	// filePath := "job1.py"

	// Read the content of the file
	code, err := os.ReadFile("E:/webdev/Golang/Ide/codee/jobs/py/job1.py")
	if err != nil {
		log.Fatal(err)
	}

	codeStr := string(code)
	// codeStr := "print('hello, world!!')"

	for _, user := range users {
		job := Job{
			UserID: user,
			Code:   codeStr, // Replace with the user's code
		}

		jobCh <- job
	}

	close(jobCh) // Close the job channel to indicate no more jobs

	// Wait for all workers to complete
	<-doneCh

}
