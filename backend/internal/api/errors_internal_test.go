package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nishojib/ffxivdailies/internal/problem"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/error", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	prob := problem.New(
		problem.WithStatus(http.StatusBadRequest),
		problem.WithDetail("bad request"),
	)

	errorResponse(rr, req, prob)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "bad request")
}

func TestLogError(t *testing.T) {
	t.Skip("TODO: fix this test")
}
