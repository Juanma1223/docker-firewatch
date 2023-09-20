package main

import (
	"docker-alarms/pkg/monitor"
	"time"

	"github.com/docker/docker/client"
)

var Cli *client.Client

func main() {

	// Initialize connection with docker socket interface
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// Main application loop
	for {
		monitor.MonitorContainers(cli)
		time.Sleep(10 * time.Second)
	}
}
