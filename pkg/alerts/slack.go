package alerts

import (
	"bytes"
	"docker-alarms/configs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SlackMessage struct {
	Message string `json:"text"`
}

func SendSlack(message string) {

	configFile, err := os.Open(os.Getenv("CONFIG_FILES_DIR") + "alerts.json")
	if err != nil {
		fmt.Println(err)
	}

	configFileBytes, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Println(err)
	}

	alertsConfig := configs.AlertsConfig{}

	json.Unmarshal(configFileBytes, &alertsConfig)

	slackMessage := SlackMessage{
		Message: message,
	}

	body, _ := json.Marshal(slackMessage)
	bodyReader := bytes.NewReader(body)

	// Create a HTTP post request
	r, err := http.NewRequest(http.MethodPost, alertsConfig.SlackWebhook, bodyReader)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
