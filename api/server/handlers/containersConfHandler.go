package handlers

import "docker-alarms/api/server/domain/response"

func UpdateContainersConf() response.Status {
	return response.SuccessfulUpdate
}
