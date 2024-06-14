package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/auth"
	"github.com/nishojib/ffxivdailies/internal/data/models"
	"github.com/nishojib/ffxivdailies/internal/data/repository"
	"github.com/nishojib/ffxivdailies/internal/validator"
	"github.com/nrednav/cuid2"
	"github.com/uptrace/bun"
)

type AccountInput struct {
	AccessToken       string `json:"access_token"`
	ExpiresAt         int    `json:"expires_at"`
	Provider          string `json:"provider"`
	ProviderAccountID string `json:"provider_account_id"`
}

func Login(db *bun.DB, authSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the account info
		var a AccountInput
		if err := api.ReadJSON(w, r, &a); err != nil {
			api.BadRequestResponse(w, r, err)
			return
		}

		if int64(a.ExpiresAt) < time.Now().Unix() {
			api.InvalidAccessTokenResponse(w, r)
			return
		}

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
		} else if a.Provider == "discord" {
			email, isAuthorized, err = auth.ValidateDiscord(a.AccessToken)
			if err != nil {
				api.InvalidAccessTokenResponse(w, r)
				return
			}
		}

		if !isAuthorized {
			api.InvalidAccessTokenResponse(w, r)
			return
		}

		user, err := repository.NewUserRepository(db).GetByProviderID(a.ProviderAccountID)
		if err != nil {
			if errors.Is(err, repository.ErrRecordNotFound) {
				user = models.User{
					Name:   "Warrior of Light",
					Email:  email,
					Image:  "https://example.com/image.png",
					UserID: cuid2.Generate(),
				}

				v := validator.New()
				if user.Validate(v); !v.Valid() {
					api.FailedValidationResponse(w, r, v.Errors)
					return
				}

				account := models.Account{
					Provider:          a.Provider,
					ProviderAccountID: a.ProviderAccountID,
					Email:             email,
				}

				if account.Validate(v); !v.Valid() {
					api.FailedValidationResponse(w, r, v.Errors)
					return
				}

				err = repository.NewUserRepository(db).InsertAndLinkAccount(&user, &account)
				if err != nil {
					api.ServerErrorResponse(w, r, err)
					return
				}

			}
			api.ServerErrorResponse(w, r, err)
			return
		}

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

		err = api.WriteJSON(w, http.StatusOK, api.Envelope[string]{"token": accessToken}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
			return
		}
	}
}

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
