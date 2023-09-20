package monitor

import (
	"context"
	"docker-alarms/configs"
	"docker-alarms/pkg/alerts"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var Config configs.ContainersConf

func MonitorContainers(cli *client.Client) {

	configFile, err := os.Open(os.Getenv("CONFIG_FILES_DIR") + "containers.json")
	if err != nil {
		fmt.Println(err)
	}

	configFileBytes, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(configFileBytes, &Config)

	watchedContainers := make(map[string]bool)

	// Create map from containers name to make a more efficient access
	if !Config.WatchAllContainers {
		for _, containerName := range Config.WatchedContainers {
			watchedContainers[containerName] = true
		}
	}

	// List all containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.State == "exited" {
			// Check if all containers should be watched or not
			if !Config.WatchAllContainers {
				containerName := ""
				splitName := strings.Split(container.Names[0], "/")
				if len(splitName) > 1 {
					containerName = string(splitName[1])
				} else {
					containerName = string(splitName[0])
				}
				if watchedContainers[containerName] {
					ContainerDownProcedure(container, cli)
				}
			} else {
				ContainerDownProcedure(container, cli)
			}
		}
	}
}

func ContainerDownProcedure(containerInfo types.Container, cli *client.Client) {
	if Config.RestartContainers {
		err := cli.ContainerStart(context.TODO(), containerInfo.ID, types.ContainerStartOptions{})
		if err != nil {
			fmt.Println(err)
		}
	}
	if Config.SendAlert {
		alerts.SendSlack("Container down! " + containerInfo.Names[0] + ", restarting it...")
	}
}
