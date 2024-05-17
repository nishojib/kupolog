package api_test

import (
	"testing"

	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEnvironment(t *testing.T) {
	testCases := map[string]struct {
		input       string
		expected    api.Environment
		expectedErr error
	}{
		"development environment": {
			input:       "development",
			expected:    api.Environment("development"),
			expectedErr: nil,
		},
		"production environment": {
			input:       "development",
			expected:    api.Environment("development"),
			expectedErr: nil,
		},
		"unknown environment": {
			input:       "unknown",
			expected:    api.Environment(""),
			expectedErr: api.ErrEnvironment,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := api.NewEnvironment(tc.input)
			if err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestEnvironmentString(t *testing.T) {
	input := "development"
	env, err := api.NewEnvironment(input)
	require.NoError(t, err)

	got := env.String()
	assert.Equal(t, "development", got)
}

func TestNewLimiter(t *testing.T) {
	limiter := api.NewLimiter(25, true)

	assert.ObjectsAreEqual(limiter, api.Limiter{25, true})

}
