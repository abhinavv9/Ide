package container

import (
	"context"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getContainerLogs(ctx context.Context, cli *client.Client, containerID string) (string, error) {
	logOptions := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       "all",
	}

	logs, err := cli.ContainerLogs(ctx, containerID, logOptions)
	if err != nil {
		return "", err
	}
	defer logs.Close()

	logBytes, err := ioutil.ReadAll(logs)
	if err != nil {
		return "", err
	}

	return string(logBytes), nil
}
