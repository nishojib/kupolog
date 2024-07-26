package auth_test

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nishojib/ffxivdailies/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	token, err := auth.NewToken(testIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	// Generate an access token
	token, err := auth.NewToken(testIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate the token
	userId, err := auth.ValidateToken(token, testSecret)
	assert.NoError(t, err)
	assert.Equal(t, testUserId, userId)
}

func TestValidateToken_InvalidIssuer(t *testing.T) {
	// Generate a refresh token
	token, err := auth.NewToken(testRefreshIssuer, testUserId, testSecret, testTokenExpiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Attempt to validate the refresh token (should fail)
	_, err = auth.ValidateToken(token, testSecret)
	assert.Error(t, err)
	assert.Equal(t, jwt.ErrTokenInvalidIssuer, err)
}
