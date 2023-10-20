package configs

type AlertsConfig struct {
	SlackWebhook string `json:"slackWebhook"`
	LogsTail     string `json:"logsTail"`
}
