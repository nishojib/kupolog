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
		authToken, err := s.authModel.GetBearerToken(r.Header)
		if err != nil {
			slog.Error("failed to get bearer token", "error", err)
			api.AuthenticationRequiredResponse(w, r)
			return
		}

		userID, err := auth.ValidateToken(authToken, s.cfg.AuthSecret)
		if err != nil {
			slog.Error("failed to validate token", "error", err)
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		if userID == "" {
			slog.Error("user id is empty")
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		user, err := s.db.GetUserByUserID(r.Context(), userID)
		if err != nil {
			slog.Error("failed to get user", "error", err)
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), auth.UserIDKey, user.UserID)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}
