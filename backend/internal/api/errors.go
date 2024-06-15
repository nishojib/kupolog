package api

import (
	"fmt"

	"log/slog"
	"net/http"

	"github.com/nishojib/ffxivdailies/internal/problem"
)

// ServerResponse logs an error and sends a 500 Internal Server Error response.
func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	message := "the server encountered a problem and could not process your request"
	problem := problem.Of(http.StatusInternalServerError).Append(problem.WithDetail(message))
	errorResponse(
		w,
		r,
		problem,
	)
}

// NotFoundResponse sends a 404 Not Found response.
func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	problem := problem.Of(http.StatusNotFound).Append(problem.WithDetail(message))
	errorResponse(w, r, problem)
}

// MethodNotAllowedResponse sends a 405 Method Not Allowed response.
func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	problem := problem.Of(http.StatusMethodNotAllowed).Append(problem.WithDetail(message))
	errorResponse(w, r, problem)
}

// BadRequestResponse sends a 400 Bad Request response.
func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	problem := problem.Of(http.StatusBadRequest).Append(problem.WithDetail(err.Error()))
	errorResponse(w, r, problem)
}

// FailedValidationResponse sends a 422 Unprocessable Entity response.
func FailedValidationResponse(
	w http.ResponseWriter,
	r *http.Request,
	errors map[string]string,
) {
	problem := problem.Of(http.StatusUnprocessableEntity).
		Append(problem.WithCustom("errors", errors))
	errorResponse(w, r, problem)
}

// EditConflictResponse sends a 409 Conflict response.
func EditConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	problem := problem.Of(http.StatusConflict).Append(problem.WithDetail(message))
	errorResponse(w, r, problem)
}

// RateLimitExceededResponse sends a 429 Too Many Requests response.
func RateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	problem := problem.Of(http.StatusTooManyRequests).Append(problem.WithDetail(message))
	errorResponse(w, r, problem)
}

// InvalidCredentialsResponse sends a 401 Unauthorized response.
func InvalidAccessTokenResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid access token"
	problem := problem.Of(http.StatusUnauthorized).Append(problem.WithDetail(message))
	errorResponse(w, r, problem)
}

// InvalidAuthenticationTokenResponse sends a 401 Unauthorized response.
func InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	problem := problem.Of(http.StatusUnauthorized).Append(problem.WithDetail(message))
	errorResponse(w, r, problem)
}

// AuthenticationRequiredResponse sends a 401 Unauthorized response.
func AuthenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	problem := problem.Of(http.StatusUnauthorized).Append(problem.WithDetail(message))
	errorResponse(w, r, problem)
}

// InactiveAccountResponse sends a 403 Forbidden response.
func InactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	problem := problem.Of(http.StatusForbidden).Append(problem.WithDetail(message))
	errorResponse(w, r, problem)
}

// NotPermittedResponse sends a 403 Forbidden response.
func NotPermittedResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	problem := problem.Of(http.StatusForbidden).Append(problem.WithDetail(message))
	errorResponse(w, r, problem)
}

// logError logs an error message with some context.
func logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	slog.Error(err.Error(), "method", method, "uri", uri)
}

// errorResponse writes a error response with a message and status code.
func errorResponse(
	w http.ResponseWriter,
	r *http.Request,
	problem *problem.Problem,
) {
	err := WriteJSON(w, problem.StatusCode(), problem, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
