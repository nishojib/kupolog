package auth_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nishojib/ffxivdailies/internal/auth"
	"github.com/stretchr/testify/assert"
)

const (
	testSecret        = "testsecret"
	testUserId        = "testuser"
	testIssuer        = auth.TokenTypeAccess
	testRefreshIssuer = auth.TokenTypeRefresh
	testTokenExpiry   = time.Hour
)

func TestNew(t *testing.T) {
	token, err := auth.New(testIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestRefreshToken(t *testing.T) {
	token, err := auth.New(testRefreshIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	refreshToken, err := auth.RefreshToken(token, testSecret)
	assert.NoError(t, err)
	assert.NotEmpty(t, refreshToken)
}

func TestRefreshToken_InvalidIssuer(t *testing.T) {
	// Generate an access token
	token, err := auth.New(testIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Attempt to refresh the access token (should fail)
	_, err = auth.RefreshToken(token, testSecret)
	assert.Error(t, err)
	assert.Equal(t, jwt.ErrTokenInvalidIssuer, err)
}

func TestValidate(t *testing.T) {
	// Generate an access token
	token, err := auth.New(testIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate the token
	userId, err := auth.Validate(token, testSecret)
	assert.NoError(t, err)
	assert.Equal(t, testUserId, userId)
}

func TestGetBearerToken(t *testing.T) {
	// Prepare a mock request with a Bearer token
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	token, err := auth.New(testIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+token)

	// Extract the Bearer token
	extractedToken, err := auth.GetBearerToken(req.Header)
	assert.NoError(t, err)
	assert.Equal(t, token, extractedToken)

	// Test with no Authorization header
	req.Header.Del("Authorization")
	_, err = auth.GetBearerToken(req.Header)
	assert.Equal(t, auth.ErrNoAuthHeaderIncluded, err)

	// Test with a malformed Authorization header
	req.Header.Set("Authorization", "InvalidToken")
	_, err = auth.GetBearerToken(req.Header)
	assert.Equal(t, auth.ErrMalformedAuthHeader, err)
}

func TestValidate_InvalidIssuer(t *testing.T) {
	// Generate a refresh token
	token, err := auth.New(testRefreshIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Attempt to validate the refresh token (should fail)
	_, err = auth.Validate(token, testSecret)
	assert.Error(t, err)
	assert.Equal(t, jwt.ErrTokenInvalidIssuer, err)
}
