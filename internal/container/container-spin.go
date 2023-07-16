package container

import (
	"context"
	"fmt"
	"log"

	"github.com/abhinavv9/codee/internal"
	"github.com/abhinavv9/codee/types"
	"github.com/docker/docker/client"
)

type Error struct {
    run error
    execute error
}

func SpinContainer(ctx context.Context, cli *client.Client, job types.Job) (string, Error) {
    var output string;
    errs := Error{}

    go func() {
        //Start the container
        containerID, err := runContainer(ctx, cli, job.Code, job.Image)
        if err != nil {
            errs.run = err
            log.Printf("Error running container for user %s: %v", job.UserID, err)
        }

        //Execute the code passed thru env variable
        outputBytes, err := internal.ExecuteCodeInContainer(ctx, cli, containerID)
        if err != nil {
            errs.execute = err
            log.Fatal(err)
        }

        output = string(outputBytes)

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

    return output, errs

}



