package scripts

import (
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainerMemoryUsage(ctx context.Context, cli *client.Client, containerID string) (uint64, error) {
	stats, err := cli.ContainerStats(ctx, containerID, false)
	if err != nil {
		return 0, err
	}
	defer stats.Body.Close()

	var memoryUsage uint64

	// Decode the stats response into a JSON object
	var statsJSON types.StatsJSON
	err = json.NewDecoder(stats.Body).Decode(&statsJSON)
	if err != nil {
		return 0, err
	}

	memoryStats := statsJSON.MemoryStats

	memoryUsage = memoryStats.Usage

	return memoryUsage, nil
}
