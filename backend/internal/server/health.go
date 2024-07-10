package server

import (
	"net/http"

	"github.com/nishojib/ffxivdailies/internal/api"
)

// ServerStatus Response for the health check
//
//	@Description	Response for the health check
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

// Health godoc
//
//	@Summary		Health check
//	@Description	Checks the health of the service
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	ServerStatus
//	@Failure		500	{object}	problem.Problem
//	@Router			/health [get]
func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	data := ServerStatus{
		Status: "available",
		SystemInfo: ServerInfo{
			Environment: s.cfg.Env.String(),
			Version:     s.cfg.Version,
		},
	}

	err := api.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		api.ServerErrorResponse(w, r, err)
	}
}
