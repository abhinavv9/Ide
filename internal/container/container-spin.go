package container

import (
	"context"
	"fmt"
	"log"

	"github.com/abhinavv9/codee/internal"
	"github.com/abhinavv9/codee/types"
	"github.com/docker/docker/client"
)

func SpinContainer(ctx context.Context, cli *client.Client, job types.Job) {

    go func() {
        //Start the container
        containerID, err := runContainer(ctx, cli, job.Code, job.Image)
        if err != nil {
            log.Printf("Error running container for user %s: %v", job.UserID, err)
        }

        //Execute the code passed thru env variable
        output, err := internal.ExecuteCodeInContainer(ctx, cli, containerID)
        if err != nil {
            log.Fatal(err)
        }

        //handle output
        fmt.Println(string(output))

        logs, err := getContainerLogs(ctx, cli, containerID)
        if err != nil {
            log.Printf("Error retrieving logs for user %s: %v", job.UserID, err)
        }

        // memUsage, err := scripts.GetContainerMemoryUsage(ctx, cli, containerID)
        // if err != nil {
        // 	log.Println("Failed to get container memory usage:", err)
        // 	return
        // }

        // Print container memory usage
        // log.Printf("Container memory used: %d bytes", memUsage)

        log.Printf("Container logs for user %s:\n%s\n", job.UserID, logs)

        err = removeContainer(ctx, cli, containerID)
        if err != nil {
            log.Printf("Error removing container for user %s: %v", job.UserID, err)
        }

    }()
}



