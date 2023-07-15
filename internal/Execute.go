package internal

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/docker/docker/client"
)

func ExecuteCodeInContainer(ctx context.Context, cli *client.Client, containerID string) ([]byte, error) {
	inspect, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return nil, err
	}

	var code string
	for _, env := range inspect.Config.Env {
		if strings.HasPrefix(env, "CODE=") {
			code = strings.TrimPrefix(env, "CODE=")
			break
		}
	}

	if code == "" {
		return nil, fmt.Errorf("code environment variable not found in the container")
	}

	// Execute the code using Go's built-in capabilities
	cmd := exec.Command("python", "-c", code)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// fmt.Println(cmd.Stdout)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return output, nil

}
