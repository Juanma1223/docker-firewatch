package configs

type ContainersConf struct {
	WatchedContainers  []string
	WatchAllContainers bool
	RestartContainers  bool
	SendAlert          bool
}
