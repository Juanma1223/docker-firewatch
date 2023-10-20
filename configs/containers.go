package configs

type ContainersConf struct {
	HostMachine        string   `json:"hostMachine"`
	WatchedContainers  []string `json:"watchedContainers"`
	WatchAllContainers bool     `json:"watchAllContainers"`
	RestartContainers  bool     `json:"restartContainers"`
	SendAlert          bool     `json:"sendAlert"`
}
