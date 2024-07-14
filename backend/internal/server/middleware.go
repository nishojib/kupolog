package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/auth"
)

func (s *Server) withAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			api.AuthenticationRequiredResponse(w, r)
			return
		}

		slog.Info("validating token", "token", authToken)

		userID, err := auth.Validate(authToken, s.cfg.AuthSecret)
		if err != nil {
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		if userID == "" {
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		user, err := s.db.GetUserByUserID(r.Context(), userID)
		if err != nil {
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), auth.UserIDKey, user.UserID)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}
