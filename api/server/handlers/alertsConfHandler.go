package handlers

import (
	"docker-alarms/api/server/domain/response"
	"docker-alarms/configs"
	"encoding/json"
	"os"
)

func UpdateAlertsConf(newConfig configs.AlertsConfig) response.Status {

	jsonInfo, err := json.Marshal(newConfig)
	if err != nil {
		return response.InternalServerError
	}

	err = os.WriteFile(os.Getenv("CONFIG_FILES_DIR")+"alerts.json", jsonInfo, 0644)
	if err != nil {
		return response.InternalServerError
	}

	return response.SuccessfulUpdate
}
