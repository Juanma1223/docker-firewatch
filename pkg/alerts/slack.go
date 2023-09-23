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

type Blocks struct {
	Blocks []MessageBlock `json:"blocks"`
}

type MessageBlock struct {
	Type   string         `json:"type"`
	Text   interface{}    `json:"text,omitempty"`
	Fields []MessageBlock `json:"fields,omitempty"`
}

func SendSlack(messageHeader, messageBody string) {

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

	messageHeaderBlock := MessageBlock{
		Type: "header",
		Text: MessageBlock{
			Type: "plain_text",
			Text: messageHeader,
		},
	}

	messageBodyBlock := MessageBlock{
		Type: "section",
		Fields: []MessageBlock{
			{
				Type: "plain_text",
				Text: messageBody,
			},
		},
	}

	message := Blocks{
		Blocks: []MessageBlock{
			messageHeaderBlock,
			messageBodyBlock,
		},
	}

	body, _ := json.Marshal(message)
	bodyReader := bytes.NewReader(body)

	// Create a HTTP post request
	r, err := http.NewRequest(http.MethodPost, alertsConfig.SlackWebhook, bodyReader)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error while sending slack notification, status code:", res.StatusCode)
		fmt.Println("Body:", string(body))
	}
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
