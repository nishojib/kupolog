package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

type UserKey string

const UserIDKey = UserKey("user_id")

const (
	TokenTypeAccess  TokenType = "access-token"
	TokenTypeRefresh TokenType = "refresh-token"
)

var (
	ErrNoAuthHeaderIncluded = errors.New("no auth header included in request")
	ErrMalformedAuthHeader  = errors.New("malformed authorization header")
)

func New(issuer TokenType, userId string, secret string, expiresIn time.Duration) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    string(issuer),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Subject:   userId,
	})

	return claims.SignedString([]byte(secret))
}

func RefreshToken(token, secret string) (string, error) {
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

	newToken, err := New(TokenTypeAccess, userId, secret, time.Hour)
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func Validate(token, secret string) (string, error) {
	claims := jwt.RegisteredClaims{}

	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	issuer, err := t.Claims.GetIssuer()
	if err != nil {
		return "", err
	}

	if issuer != string(TokenTypeAccess) {
		return "", jwt.ErrTokenInvalidIssuer
	}

	userIdString, err := t.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return userIdString, nil
}

func GetBearerToken(headers http.Header) (string, error) {
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
