package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthModel struct {
	tokenRevoker TokenRevoker
}

func NewModel(revoker TokenRevoker) *AuthModel {
	return &AuthModel{tokenRevoker: revoker}
}

func (m *AuthModel) CreateTokens(
	ctx context.Context,
	userID string,
	authSecret string,
) (string, string, error) {
	tokenExpiresIn := time.Hour * 1
	accessToken, err := NewToken(
		TokenTypeAccess,
		userID,
		authSecret,
		tokenExpiresIn,
	)
	if err != nil {
		return "", "", fmt.Errorf("error creating access token %s", err.Error())
	}

	refreshTokenExpiresIn := time.Hour * 24 * 60
	refreshToken, err := NewToken(
		TokenTypeRefresh,
		userID,
		authSecret,
		refreshTokenExpiresIn,
	)
	if err != nil {
		return "", "", fmt.Errorf("error creating refresh token %s", err.Error())
	}

	return accessToken, refreshToken, nil
}

func (am *AuthModel) RefreshToken(token, secret string) (string, error) {
	claims := jwt.RegisteredClaims{}

	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	userId, err := t.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	issuer, err := t.Claims.GetIssuer()
	if err != nil {
		return "", err
	}

	if issuer != string(TokenTypeRefresh) {
		return "", jwt.ErrTokenInvalidIssuer
	}

	newToken, err := NewToken(TokenTypeAccess, userId, secret, time.Hour)
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func (am *AuthModel) GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "Bearer" {
		return "", ErrMalformedAuthHeader
	}

	return splitAuth[1], nil
}
