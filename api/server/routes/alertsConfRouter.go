package routes

import (
	"docker-alarms/api/server/handlers"
	"docker-alarms/api/server/helpers/responseHelper"
	"docker-alarms/configs"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type AlertsConfRouter struct {
}

func (ar *AlertsConfRouter) UpdateAlertsConf(w http.ResponseWriter, r *http.Request) {
	editConfig := configs.AlertsConfig{}
	err := json.NewDecoder(r.Body).Decode(&editConfig)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	status := handlers.UpdateAlertsConf(editConfig)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	w.WriteHeader(status.StatusCode())
	responseHelper.ResponseStatusChecker(w, resp)
}

func (ar *AlertsConfRouter) Routes() http.Handler {
	r := chi.NewRouter()

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"https://*", "http://*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowOriginFunc:    func(r *http.Request, origin string) bool { return true },
		AllowCredentials:   true,
		OptionsPassthrough: true,
		Debug:              true,
		MaxAge:             300,
	}))

	r.Put("/", ar.UpdateAlertsConf)

	return r
}
