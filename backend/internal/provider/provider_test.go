package provider_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/nishojib/ffxivdailies/internal/provider"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	tests := map[string]struct {
		statusCode    int
		response      interface{}
		provider      string
		token         string
		url           string
		expectedErr   error
		expectedEmail string
		expectedValid bool
	}{
		"success google": {
			statusCode:    http.StatusOK,
			response:      map[string]string{"email": "test@example.com"},
			provider:      "google",
			url:           "https://oauth2.googleapis.com/tokeninfo?access_token=valid-google-token",
			token:         "valid-google-token",
			expectedEmail: "test@example.com",
			expectedValid: true,
			expectedErr:   nil,
		},
		"success discord": {
			statusCode:    http.StatusOK,
			response:      map[string]string{"email": "test@example.com"},
			provider:      "discord",
			url:           "https://discord.com/api/v10/users/@me",
			token:         "valid-discord-token",
			expectedEmail: "test@example.com",
			expectedValid: true,
			expectedErr:   nil,
		},
		"invalid token": {
			statusCode:    http.StatusOK,
			response:      map[string]string{},
			provider:      "google",
			url:           "https://oauth2.googleapis.com/tokeninfo?access_token=invalid-google-token",
			token:         "invalid-message-token",
			expectedEmail: "",
			expectedValid: false,
			expectedErr:   provider.ErrGetTokenInfo,
		},
		"not empty message": {
			statusCode:    http.StatusOK,
			response:      map[string]string{"message": "invalid"},
			provider:      "google",
			url:           "https://oauth2.googleapis.com/tokeninfo?access_token=invalid-google-token",
			token:         "invalid-google-token",
			expectedEmail: "",
			expectedValid: false,
			expectedErr:   provider.ErrMessage,
		},
		"malformed response from provider": {
			statusCode:    http.StatusOK,
			response:      map[string]int{"email": 3},
			provider:      "google",
			url:           "https://oauth2.googleapis.com/tokeninfo?access_token=invalid-google-token",
			token:         "invalid-google-token",
			expectedEmail: "",
			expectedValid: false,
			expectedErr:   provider.ErrMalformedInput,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			response, err := json.Marshal(tc.response)
			require.NoError(t, err)
			httpmock.RegisterResponder(
				"GET",
				tc.url,
				httpmock.NewStringResponder(tc.statusCode, string(response)),
			)

			// Create a Provider instance with the test server's client
			p := provider.New()

			gotEmail, gotValid, err := p.Validate(tc.provider, tc.token)
			require.ErrorIs(t, err, tc.expectedErr)

			if tc.expectedErr == nil {
				require.Equal(t, tc.expectedEmail, gotEmail)
				require.Equal(t, tc.expectedValid, gotValid)
			}
		})
	}
}
