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
	MessageStruct []MessageBlock
}

type MessageBlock struct {
	Type   string         `json:"type"`
	Text   interface{}    `json:"text"`
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
				Type: "mrkdown",
				Text: "\n" + messageBody,
			},
		},
	}

	message := SlackMessage{
		[]MessageBlock{
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
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
