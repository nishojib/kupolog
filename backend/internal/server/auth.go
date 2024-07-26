package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/user"
)

// AccountRequest represents the request body for the login endpoint.
type AccountRequest struct {
	AccessToken       string `json:"access_token"`
	ExpiresAt         int    `json:"expires_at"`
	Provider          string `json:"provider"`
	ProviderAccountID string `json:"provider_account_id"`
}

// LoginResponse represents the response for the login endpoint.
type LoginResponse struct {
	User  LoginUserResponse  `json:"user"`
	Token LoginTokenResponse `json:"token"`
}

// LoginUserResponse represents the response for the user of the login endpoint.
type LoginUserResponse struct {
	UserID    string    `json:"userID"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
}

// LoginTokenResponse represents the response for the token of the login endpoint.
type LoginTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login godoc
//
//	@Summary		login
//	@Description	takes a google or discord account request verifies the account and returns a token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AccountRequest	true	"account request"
//	@Success		200		{object}	LoginResponse
//	@Failure		400		{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		401		{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500		{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/auth/login [post]
func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// get the account info
	var a AccountRequest
	if err := api.ReadJSON(w, r, &a); err != nil {
		slog.Error("failed to read json", "error", err)
		api.BadRequestResponse(w, r, err)
		return
	}

	if int64(a.ExpiresAt) < time.Now().Unix() {
		slog.Error("access token expired")
		api.InvalidAccessTokenResponse(w, r)
		return
	}

	// verify the access token
	var email string
	var isAuthorized bool
	var err error

	email, isAuthorized, err = s.provider.Validate(a.Provider, a.AccessToken)
	if err != nil {
		slog.Error("failed to validate access token", "error", err)
		api.InvalidAccessTokenResponse(w, r)
		return
	}

	if !isAuthorized {
		slog.Error("access token is not authorized")
		api.InvalidAccessTokenResponse(w, r)
		return
	}

	u, err := s.userModel.GetOrCreate(
		r.Context(),
		user.Email(email),
		user.Provider(a.Provider),
		user.ID(a.ProviderAccountID),
	)
	if err != nil {
		api.ServerErrorResponse(w, r, err)
		return
	}

	accessToken, refreshToken, err := s.authModel.CreateTokens(
		r.Context(),
		string(u.UserID),
		s.cfg.AuthSecret,
	)
	if err != nil {
		api.ServerErrorResponse(
			w,
			r,
			fmt.Errorf("error creating access token %s", err.Error()),
		)
		return
	}

	err = api.WriteJSON(
		w,
		http.StatusOK,
		LoginResponse{
			User: LoginUserResponse{
				UserID:    string(u.UserID),
				Name:      string(u.Name),
				Email:     string(u.Email),
				Image:     string(u.Image),
				CreatedAt: u.CreatedAt,
			}, Token: LoginTokenResponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		},
		nil,
	)
	if err != nil {
		api.ServerErrorResponse(w, r, err)
		return
	}
}

// RefreshTokenResponse represents the response for the refresh token endpoint.
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// RefreshToken godoc
//
//	@Summary		refresh token
//	@Description	refreshes the access token for the user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Param			Authorization	header		string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Success		200				{object}	object{access_token=string}
//	@Failure		400				{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		401				{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500				{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/auth/refresh [post]
func (s *Server) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := s.authModel.GetBearerToken(r.Header)
	if err != nil {
		api.BadRequestResponse(w, r, errors.New("couldn't find JWT"))
		return
	}

	isRevoked, err := s.authModel.IsTokenRevoked(r.Context(), refreshToken)
	if err != nil {
		api.ServerErrorResponse(w, r, errors.New("couldn't check session"))
		return
	}
	if isRevoked {
		api.InvalidAccessTokenResponse(w, r)
		return
	}

	accessToken, err := s.authModel.RefreshToken(refreshToken, s.cfg.AuthSecret)
	if err != nil {
		api.InvalidAccessTokenResponse(w, r)
		return
	}

	err = api.WriteJSON(
		w,
		http.StatusOK,
		AccountRequest{AccessToken: accessToken},
		nil,
	)
	if err != nil {
		api.ServerErrorResponse(w, r, err)
		return
	}
}

// RevokeToken godoc
//
//	@Summary		revoke token
//	@Description	revokes the refresh token for the user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Param			Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Success		200
//	@Failure		400	{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500	{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/auth/revoke [post]
func (s *Server) RevokeTokenHandler(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := s.authModel.GetBearerToken(r.Header)
	if err != nil {
		api.BadRequestResponse(w, r, fmt.Errorf("couldn't find JWT, %s", err.Error()))
		return
	}

	dbCtx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	err = s.db.RevokeToken(dbCtx, refreshToken)
	if err != nil {
		api.ServerErrorResponse(w, r, fmt.Errorf("couldn't revoke session, %s", err.Error()))
		return
	}

	err = api.WriteJSON(w, http.StatusOK, struct{}{}, nil)
	if err != nil {
		api.ServerErrorResponse(w, r, err)
		return
	}
}
