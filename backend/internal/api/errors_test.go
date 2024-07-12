package api_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorResponses(t *testing.T) {
	tests := map[string]struct {
		statusCode int
		f          func(w http.ResponseWriter, r *http.Request)
		errMessage string
	}{
		"server error": {
			statusCode: http.StatusInternalServerError,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.ServerErrorResponse(w, r, errors.New("test error"))
			},
			errMessage: "the server encountered a problem and could not process your request",
		},
		"not found": {
			statusCode: http.StatusNotFound,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.NotFoundResponse(w, r)
			},
			errMessage: "the requested resource could not be found",
		},
		"method not allowed": {
			statusCode: http.StatusMethodNotAllowed,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.MethodNotAllowedResponse(w, r)
			},
			errMessage: "the GET method is not supported for this resource",
		},
		"bad request": {
			statusCode: http.StatusBadRequest,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.BadRequestResponse(w, r, errors.New("test error"))
			},
			errMessage: "test error",
		},
		"failed validation": {
			statusCode: http.StatusUnprocessableEntity,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.FailedValidationResponse(w, r, map[string]string{"name": "name is required"})
			},
			errMessage: "name is required",
		},
		"edit conflict": {
			statusCode: http.StatusConflict,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.EditConflictResponse(w, r)
			},
			errMessage: "unable to update the record due to an edit conflict, please try again",
		},
		"rate limit exceeded": {
			statusCode: http.StatusTooManyRequests,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.RateLimitExceededResponse(w, r)
			},
			errMessage: "rate limit exceeded",
		},
		"invalid access token": {
			statusCode: http.StatusUnauthorized,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.InvalidAccessTokenResponse(w, r)
			},
			errMessage: "invalid access token",
		},
		"invalid authentication token": {
			statusCode: http.StatusUnauthorized,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.InvalidAuthenticationTokenResponse(w, r)
			},
			errMessage: "invalid or missing authentication token",
		},
		"authentication required": {
			statusCode: http.StatusUnauthorized,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.AuthenticationRequiredResponse(w, r)
			},
			errMessage: "you must be authenticated to access this resource",
		},
		"inactive account": {
			statusCode: http.StatusForbidden,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.InactiveAccountResponse(w, r)
			},
			errMessage: "your user account must be activated to access this resource",
		},
		"not permitted": {
			statusCode: http.StatusForbidden,
			f: func(w http.ResponseWriter, r *http.Request) {
				api.NotPermittedResponse(w, r)
			},
			errMessage: "your user account doesn't have the necessary permissions to access this resource",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/", nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tc.f)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tc.statusCode, rr.Code)
			assert.Contains(
				t,
				rr.Body.String(),
				tc.errMessage,
			)
		})
	}
}
