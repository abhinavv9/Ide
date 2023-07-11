package container

import (
	"context"
	"fmt"

	// "github.com/abinavv9/codee/scripts"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func runContainer(ctx context.Context, cli *client.Client, code string, image string) (string, error) {

	memLimit := int64(50 * 1024 * 1024) // 50MB memory limit
	// cpuLimit := float64(0.5)           // 50% CPU limit

	// Set environment variables for the container
	env := []string{
		fmt.Sprintf("CODE=%s", code),
	}

	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: image, // Replace with your Docker image name
			// Cmd:          []string{"sh", "-c", code},
			AttachStdout: true,
			AttachStderr: true,
			Env:          env,
		},
		&container.HostConfig{
			Resources: container.Resources{
				Memory: memLimit,
				// NanoCPUs: int64(cpuLimit * 1e9),
			},
		},
		nil,
		nil,
		"",
	)
	if err != nil {
		return "", err
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}
