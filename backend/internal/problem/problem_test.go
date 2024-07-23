package problem_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/nishojib/ffxivdailies/internal/problem"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	p := problem.New(problem.WithTitle("title string"), problem.WithCustom("x", "value"))

	assert.JSONEq(t, p.JSONString(), `{"title":"title string", "x": "value"}`)

	b, err := json.Marshal(p)
	require.NoError(t, err)
	assert.JSONEq(t, string(b), `{"title":"title string", "x": "value"}`)

	p = problem.New(
		problem.WithTitle("title string"),
		problem.WithStatus(404),
		problem.WithCustom("x", "value"),
	)
	str := p.JSONString()
	assert.JSONEq(t, str, `{"title":"title string", "x": "value", "status":404}`)

	p.Append(
		problem.WithDetail("some more details"),
		problem.WithInstance("https://example.com/details"),
	)
	str = p.JSONString()
	assert.JSONEq(
		t,
		str,
		`{"title":"title string", "x": "value", "status":404, "detail":"some more details", "instance":"https://example.com/details"}`,
	)

	p = problem.Of(http.StatusAccepted)
	str = p.JSONString()
	assert.JSONEq(
		t,
		str,
		`{"status":202, "title":"Accepted", "type":"https://tools.ietf.org/html/rfc9110#section-15.3.3"}`,
	)

	assert.Equal(t, *p, p.Problem())
}

func TestProblemHTTP(t *testing.T) {
	p := problem.New(
		problem.WithTitle("title string"),
		problem.WithStatus(404),
		problem.WithCustom("x", "value"),
	)
	p.Append(
		problem.WithDetail("some more details"),
		problem.WithInstance("https://example.com/details"),
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.Append(problem.WithType("https://example.com/404"))
		if r.Method == "HEAD" {
			p.WriteHeaderTo(w)
		} else {
			p.WriteTo(w)
		}
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	require.NoError(t, err)

	bodyBytes, err := io.ReadAll(res.Body)
	res.Body.Close()
	require.NoError(t, err)

	assert.Equal(t, res.StatusCode, http.StatusNotFound)
	assert.Equal(t, res.Header.Get("Content-Type"), problem.ContentTypeJson)
	assert.Equal(
		t,
		string(bodyBytes),
		`{"detail":"some more details","instance":"https://example.com/details","status":404,"title":"title string","type":"https://example.com/404","x":"value"}`,
	)

	// Try HEAD request
	res, err = http.Head(ts.URL)
	require.NoError(t, err)

	bodyBytes, err = io.ReadAll(res.Body)
	res.Body.Close()
	require.NoError(t, err)

	assert.Zero(t, len(bodyBytes))
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
	assert.Equal(t, res.Header.Get("Content-Type"), problem.ContentTypeJson)
}

func TestMarshalUnmarshal(t *testing.T) {
	p := problem.New(problem.WithStatus(500), problem.WithTitle("Strange"))

	newProblem := problem.New()
	err := json.Unmarshal(p.JSON(), &newProblem)
	require.NoError(t, err)
	assert.Equal(t, p.Error(), newProblem.Error())
}

func TestErrors(t *testing.T) {
	knownProblem := problem.New(problem.WithStatus(404), problem.WithTitle("not found"))

	responseFromExternalService := http.Response{
		Header: map[string][]string{
			"Content-Type": {problem.ContentTypeJson},
		},
		Body: io.NopCloser(strings.NewReader(`{"status":404, "title":"not found"}`)),
	}
	defer responseFromExternalService.Body.Close()

	if responseFromExternalService.Header.Get("Content-Type") == problem.ContentTypeJson {
		problemDecoder := json.NewDecoder(responseFromExternalService.Body)

		problemFromExternalService := problem.New()
		problemDecoder.Decode(&problemFromExternalService)

		assert.ErrorIs(t, problemFromExternalService, knownProblem)
	}
}

func TestNestedErrors(t *testing.T) {
	rootProblem := problem.New(problem.WithStatus(404), problem.WithTitle("not found"))
	p := problem.New(problem.Wrap(rootProblem), problem.WithTitle("high level error msg"))

	unwrappedProblem := errors.Unwrap(p)
	assert.ErrorIs(t, unwrappedProblem, rootProblem)
	assert.Empty(t, errors.Unwrap(unwrappedProblem))

	// See wrapped error in 'reason'
	assert.JSONEq(
		t,
		p.JSONString(),
		`{"reason":"{\"status\":404,\"title\":\"not found\"}", "title":"high level error msg"}`,
	)

	p = problem.New(problem.WrapSilent(rootProblem), problem.WithTitle("high level error msg"))
	assert.JSONEq(t, p.JSONString(), `{"title":"high level error msg"}`)
}

func TestOSErrorInProblem(t *testing.T) {
	_, err := os.ReadFile("non-existing")
	if err != nil {
		p := problem.New(
			problem.Wrap(err),
			problem.WithTitle("Internal Error"),
			problem.WithStatus(404),
		)
		assert.ErrorIs(t, p, os.ErrNotExist)
		assert.NotErrorIs(t, p, os.ErrPermission)

		var o *os.PathError
		assert.ErrorAs(t, p, &o)

		newErr := errors.New("new error")
		p = problem.New(problem.Wrap(newErr), problem.WithTitle("new problem"))

		assert.ErrorIs(t, p, newErr)
	}
}

func TestWithTitlef(t *testing.T) {
	expected := `{"title":"this is a test"}`
	toTest := problem.New(problem.WithTitlef("this is a %s", "test")).JSONString()
	assert.Contains(t, expected, toTest)
}

func TestWithDetailf(t *testing.T) {
	expected := `{"detail":"this is a test"}`
	toTest := problem.New(problem.WithDetailf("this is a %s", "test")).JSONString()
	assert.Contains(t, expected, toTest)
}

func TestWithInstancef(t *testing.T) {
	expected := `{"instance":"this is a test"}`
	toTest := problem.New(problem.WithInstancef("this is a %s", "test")).JSONString()
	assert.Contains(t, expected, toTest)
}

func TestStatusCode(t *testing.T) {
	p := problem.New(problem.WithStatus(404))
	assert.Equal(t, p.StatusCode(), 404)

	p = problem.New()
	assert.Equal(t, p.StatusCode(), 500)
}
