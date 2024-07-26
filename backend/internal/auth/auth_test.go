package auth_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nishojib/ffxivdailies/internal/auth"
	"github.com/nishojib/ffxivdailies/internal/auth/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	testSecret        = "testsecret"
	testUserId        = "testuser"
	testIssuer        = auth.TokenTypeAccess
	testRefreshIssuer = auth.TokenTypeRefresh
	testTokenExpiry   = time.Hour
)

func TestRefreshToken(t *testing.T) {
	token, err := auth.NewToken(testRefreshIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	authModel := auth.NewModel(mocks.NewTokenRevoker(t))

	refreshToken, err := authModel.RefreshToken(token, testSecret)
	assert.NoError(t, err)
	assert.NotEmpty(t, refreshToken)
}

func TestRefreshToken_InvalidIssuer(t *testing.T) {
	// Generate an access token
	token, err := auth.NewToken(testIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	authModel := auth.NewModel(mocks.NewTokenRevoker(t))

	// Attempt to refresh the access token (should fail)
	_, err = authModel.RefreshToken(token, testSecret)
	assert.Error(t, err)
	assert.Equal(t, jwt.ErrTokenInvalidIssuer, err)
}

func TestGetBearerToken(t *testing.T) {
	// Prepare a mock request with a Bearer token
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	token, err := auth.NewToken(testIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+token)

	authModel := auth.NewModel(mocks.NewTokenRevoker(t))

	// Extract the Bearer token
	extractedToken, err := authModel.GetBearerToken(req.Header)
	assert.NoError(t, err)
	assert.Equal(t, token, extractedToken)

	// Test with no Authorization header
	req.Header.Del("Authorization")
	_, err = authModel.GetBearerToken(req.Header)
	assert.Equal(t, auth.ErrNoAuthHeaderIncluded, err)

	// Test with a malformed Authorization header
	req.Header.Set("Authorization", "InvalidToken")
	_, err = authModel.GetBearerToken(req.Header)
	assert.Equal(t, auth.ErrMalformedAuthHeader, err)
}
