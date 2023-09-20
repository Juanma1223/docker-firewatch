package main

import (
	"docker-alarms/api/server"
	"docker-alarms/pkg/monitor"
	"log"
	"os"
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

	// Start monitoring containers
	go RunContainersMonitor(cli)

	// Start configs API
	StartApi()
}

func RunContainersMonitor(cli *client.Client) {
	// Start monitoring containers
	for {
		monitor.MonitorContainers(cli)
		time.Sleep(10 * time.Second)
	}
}

func StartApi() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4200"
	}
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	serv.Start()
}
