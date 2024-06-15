package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/auth"
	"github.com/nishojib/ffxivdailies/internal/data/models"
	"github.com/nishojib/ffxivdailies/internal/data/repository"
	"github.com/nrednav/cuid2"

	"github.com/uptrace/bun"
)

// AccountRequest represents the request body for the login endpoint.
type AccountRequest struct {
	AccessToken       string `json:"access_token"`
	ExpiresAt         int    `json:"expires_at"`
	Provider          string `json:"provider"`
	ProviderAccountID string `json:"provider_account_id"`
}

// Login godoc
//
//	@Summary		login
//	@Description	takes a google or discord account request verifies the account and returns a token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AccountRequest	true	"account request"
//	@Success		200		{object}	object{user=models.User,token=object{access_token=string,refresh_token=string}}
//	@Failure		400		{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		401		{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500		{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/auth/login [post]
func Login(db *bun.DB, authSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the account info
		var a AccountRequest
		if err := api.ReadJSON(w, r, &a); err != nil {
			api.BadRequestResponse(w, r, err)
			return
		}

		if int64(a.ExpiresAt) < time.Now().Unix() {
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		slog.Info(
			"login",
			"access_token",
			a.AccessToken,
			"expires_at",
			a.ExpiresAt,
			"provider",
			a.Provider,
			"provider_account_id",
			a.ProviderAccountID,
		)

		// verify the access token
		var email string
		var isAuthorized bool
		var err error
		if a.Provider == "google" {
			email, isAuthorized, err = auth.ValidateGoogle(a.AccessToken)
			if err != nil {
				api.InvalidAccessTokenResponse(w, r)
				return
			}
			slog.Info("GOOGLE VALIDATED")
		} else if a.Provider == "discord" {
			email, isAuthorized, err = auth.ValidateDiscord(a.AccessToken)
			if err != nil {
				api.InvalidAccessTokenResponse(w, r)
				return
			}
			slog.Info("DISCORD VALIDATED")
		}

		slog.Info("login", "email", email, "isAuthorized", isAuthorized)

		if !isAuthorized {
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		user := models.User{
			Name:   "Warrior of Light",
			Email:  email,
			Image:  "https://example.com/image.png",
			UserID: cuid2.Generate(),
		}

		// user, err := repository.NewUserRepository(db).GetByProviderID(a.ProviderAccountID)
		// if err != nil {
		// 	slog.Info("failed to get user", "error", err)

		// 	if errors.Is(err, repository.ErrRecordNotFound) {
		// 		user = models.User{
		// 			Name:   "Warrior of Light",
		// 			Email:  email,
		// 			Image:  "https://example.com/image.png",
		// 			UserID: cuid2.Generate(),
		// 		}

		// 		account := models.Account{
		// 			Provider:          a.Provider,
		// 			ProviderAccountID: a.ProviderAccountID,
		// 			Email:             email,
		// 		}

		// 		err = repository.NewUserRepository(db).InsertAndLinkAccount(&user, &account)
		// 		if err != nil {
		// 			slog.Info("failed to insert user", "error", err)

		// 			api.ServerErrorResponse(w, r, err)
		// 			return
		// 		}
		// 	} else {
		// 		slog.Info("failed to get user", "error", err)

		// 		api.ServerErrorResponse(w, r, err)
		// 		return
		// 	}
		// }

		// slog.Info("login", "GOT USER", user)

		tokenExpiresIn := time.Hour * 1
		accessToken, err := auth.New(auth.TokenTypeAccess, user.UserID, authSecret, tokenExpiresIn)
		if err != nil {
			api.ServerErrorResponse(
				w,
				r,
				fmt.Errorf("error creating access token %s", err.Error()),
			)
			return
		}

		refreshTokenExpiresIn := time.Hour * 24 * 60
		refreshToken, err := auth.New(
			auth.TokenTypeRefresh,
			user.UserID,
			authSecret,
			refreshTokenExpiresIn,
		)
		if err != nil {
			api.ServerErrorResponse(
				w,
				r,
				fmt.Errorf("error creating refresh token %s", err.Error()),
			)
			return
		}

		err = api.WriteJSON(
			w,
			http.StatusOK,
			api.Envelope[any]{"user": user, "token": struct {
				AccessToken  string `json:"access_token"`
				RefreshToken string `json:"refresh_token"`
			}{AccessToken: accessToken, RefreshToken: refreshToken}},
			nil,
		)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
			return
		}
	}
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
func RefreshToken(db *bun.DB, authSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refreshToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			api.BadRequestResponse(w, r, errors.New("couldn't find JWT"))
			return
		}

		isRevoked, err := repository.NewRevocationRepository(db).IsTokenRevoked(refreshToken)
		if err != nil {
			api.ServerErrorResponse(w, r, errors.New("couldn't check session"))
			return
		}
		if isRevoked {
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		accessToken, err := auth.RefreshToken(refreshToken, authSecret)
		if err != nil {
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		err = api.WriteJSON(
			w,
			http.StatusOK,
			api.Envelope[string]{"access_token": accessToken},
			nil,
		)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
			return
		}
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
func RevokeToken(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refreshToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			api.BadRequestResponse(w, r, fmt.Errorf("couldn't find JWT, %s", err.Error()))
			return
		}

		err = repository.NewRevocationRepository(db).RevokeToken(refreshToken)
		if err != nil {
			api.ServerErrorResponse(w, r, fmt.Errorf("couldn't revoke session, %s", err.Error()))
			return
		}

		err = api.WriteJSON(w, http.StatusOK, api.Envelope[string]{}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
			return
		}
	}
}
