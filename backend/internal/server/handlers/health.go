package handlers

import (
	"net/http"

	"github.com/nishojib/ffxivdailies/internal/api"
)

// ServerStatus Response for the health check
//
//	@Description Response for the health check
type ServerStatus struct {
	// Status is the health status of the service
	Status string `json:"status"`
	// SystemInfo contains information about the system
	SystemInfo ServerInfo `json:"system_info"`
}

type ServerInfo struct {
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

// Health
//
//	@Summary		Health check
//	@Description	Checks the health of the service
//	@Tags			health
//	@Produce		json
//	@Success		200	{object} 	ServerStatus
//	@Failure		500	{object}	problem.Problem
//	@Router			/health [get]
func Health(env api.Environment, version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := ServerStatus{
			Status: "available",
			SystemInfo: ServerInfo{
				Environment: env.String(),
				Version:     version,
			},
		}

		err := api.WriteJSON(w, http.StatusOK, data, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
		}
	}
}
