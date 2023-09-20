package handlers

import (
	"docker-alarms/api/server/domain/response"
	"docker-alarms/configs"
	"encoding/json"
	"os"
)

func UpdateContainersConf(newConfig configs.ContainersConf) response.Status {

	jsonInfo, err := json.Marshal(newConfig)
	if err != nil {
		return response.InternalServerError
	}

	err = os.WriteFile(os.Getenv("CONFIG_FILES_DIR")+"containers.json", jsonInfo, 0644)
	if err != nil {
		return response.InternalServerError
	}

	return response.SuccessfulUpdate
}
