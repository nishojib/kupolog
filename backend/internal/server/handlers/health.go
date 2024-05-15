package handlers

import (
	"net/http"

	"github.com/nishojib/ffxivdailies/internal/api"
)

func Health(env api.Environment, version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := api.Envelope[any]{
			"status": "available",
			"system_info": map[string]string{
				"environment": env.String(),
				"version":     version,
			},
		}

		err := api.WriteJSON(w, http.StatusOK, data, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
		}
	}
}
