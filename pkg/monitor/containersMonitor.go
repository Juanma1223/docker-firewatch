package monitor

import (
	"bytes"
	"context"
	"docker-alarms/api/server/helpers/configParser"
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

var ContainersDown map[string]bool = make(map[string]bool)

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
		// Check that alert hasn't been thrown yet
		if container.State == "exited" {
			// Check if all containers should be watched or not
			if !Config.WatchAllContainers {
				// Parse container name
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
		} else if ContainersDown[container.Names[0]] {
			// Application was down and now is up and running, send alert
			ContainerRunningProcedure(container)
		}
	}
}

func ContainerDownProcedure(containerInfo types.Container, cli *client.Client) {

	// If alert has already been sent, don't alert again
	if !ContainersDown[containerInfo.Names[0]] {
		// Register container as down
		ContainersDown[containerInfo.Names[0]] = true

		// Get alerts config file info
		alertsConfig := configs.AlertsConfig{}

		configParser.ParseConfigFile(os.Getenv("CONFIG_FILES_DIR")+"alerts.json", &alertsConfig)

		if Config.RestartContainers {
			err := cli.ContainerStart(context.TODO(), containerInfo.ID, types.ContainerStartOptions{})
			if err != nil {
				fmt.Println(err)
			}
		}
		if Config.SendAlert {
			reader, err := cli.ContainerLogs(context.TODO(), containerInfo.ID, types.ContainerLogsOptions{
				ShowStdout: true,
				ShowStderr: true,
				Tail:       alertsConfig.LogsTail,
			})
			if err != nil {
				fmt.Println(err)
			}

			buf := new(bytes.Buffer)
			buf.ReadFrom(reader)
			respBytes := buf.String()

			logs := string(respBytes)

			alerts.SendSlack(":red_circle: App down! "+containerInfo.Names[0]+", restarting it... :red_circle:", logs)
		}
	}
}

func ContainerRunningProcedure(containerInfo types.Container) {
	ContainersDown[containerInfo.Names[0]] = false
	alerts.SendSlack(":large_green_circle:"+containerInfo.Names[0]+" App running again! :large_green_circle:", ":muscle::skin-tone-3:")
}
