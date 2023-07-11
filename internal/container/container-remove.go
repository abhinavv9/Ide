package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func removeContainer(ctx context.Context, cli *client.Client, containerID string) error {
	return cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})
}
