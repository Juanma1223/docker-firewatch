package configs

type ContainersConf struct {
	WatchedContainers  []string `json:"watchedContainers"`
	WatchAllContainers bool     `json:"watchAllContainers"`
	RestartContainers  bool     `json:"restartContainers"`
	SendAlert          bool     `json:"sendAlert"`
}
