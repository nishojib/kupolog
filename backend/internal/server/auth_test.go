package server_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"
	"time"

	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/auth"
	"github.com/nishojib/ffxivdailies/internal/provider"
	"github.com/nishojib/ffxivdailies/internal/server"
	"github.com/nishojib/ffxivdailies/internal/server/mocks"
	"github.com/nishojib/ffxivdailies/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUser struct {
	mock.Mock
}

func (m *MockUser) GetOrCreate(
	ctx context.Context,
	db user.UserCreator,
	email string,
	provider string,
	accountID string,
) (user.User, error) {
	args := m.Called(ctx, db, email, provider, accountID)
	return args.Get(0).(user.User), args.Error(1)
}

func TestLoginHandler(t *testing.T) {
	expectedUser := user.User{
		ID:        1,
		Name:      "Warrior of Light",
		Email:     "user@example.com",
		Image:     "https://example.com/image.png",
		CreatedAt: time.Now(),
		UserID:    "123456",
	}

	tests := map[string]struct {
		body           string
		buildStubs     func(mockUserModel *mocks.UserModel, mockAuthModel *mocks.AuthModel, mockProvider *mocks.Provider)
		expectedStatus int
		expectedUser   user.User
		expectedFunc   func(rec *httptest.ResponseRecorder)
	}{
		"successful login": {
			body: `{"provider": "test", "access_token": "validtoken", "expires_at": 1893456000, "provider_account_id": "12345"}`,
			buildStubs: func(mockUserModel *mocks.UserModel, mockAuthModel *mocks.AuthModel, mockProvider *mocks.Provider) {
				mockProvider.EXPECT().
					Validate("test", "validtoken").
					Return(string(expectedUser.Email), true, nil)

				mockAuthModel.EXPECT().
					CreateTokens(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return("accesstoken", "refreshtoken", nil)

				mockUserModel.EXPECT().
					GetOrCreate(mock.Anything, mock.AnythingOfType("user.Email"), mock.AnythingOfType("user.Provider"), mock.AnythingOfType("user.ID")).
					Return(expectedUser, nil)
			},
			expectedStatus: http.StatusOK,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "access_token")
				assert.Contains(t, rec.Body.String(), "refresh_token")
				assert.Contains(t, rec.Body.String(), "user")
			},
		},
		"bad account info": {
			body:           `{"provider": "test", "access_token": "validtoken", "expires_at": 1893456000, "provider_account_id": "`,
			buildStubs:     func(_ *mocks.UserModel, _ *mocks.AuthModel, _ *mocks.Provider) {},
			expectedStatus: http.StatusBadRequest,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Bad Request")
			},
		},
		"expired account info": {
			body:           `{"provider": "test", "access_token": "validtoken", "expires_at": 0, "provider_account_id": "12345"}`,
			buildStubs:     func(_ *mocks.UserModel, _ *mocks.AuthModel, _ *mocks.Provider) {},
			expectedStatus: http.StatusUnauthorized,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Unauthorized")
			},
		},
		"invalid provider": {
			body: `{"provider": "invalid", "access_token": "validtoken", "expires_at": 1893456000, "provider_account_id": "12345"}`,
			buildStubs: func(_ *mocks.UserModel, _ *mocks.AuthModel, mockProvider *mocks.Provider) {
				mockProvider.EXPECT().
					Validate("invalid", "validtoken").
					Return("", false, provider.ErrRequestFailed)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Unauthorized")
			},
		},
		"unauthorized provider": {
			body: `{"provider": "google", "access_token": "validtoken", "expires_at": 1893456000, "provider_account_id": "12345"}`,
			buildStubs: func(_ *mocks.UserModel, _ *mocks.AuthModel, mockProvider *mocks.Provider) {
				mockProvider.EXPECT().
					Validate("google", "validtoken").
					Return("", false, nil)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Unauthorized")
			},
		},
		"error getting user": {
			body: `{"provider": "test", "access_token": "validtoken", "expires_at": 1893456000, "provider_account_id": "12345"}`,
			buildStubs: func(mockUserModel *mocks.UserModel, _ *mocks.AuthModel, mockProvider *mocks.Provider) {
				mockProvider.EXPECT().
					Validate("test", "validtoken").
					Return(string(expectedUser.Email), true, nil)

				mockUserModel.EXPECT().
					GetOrCreate(mock.Anything, mock.AnythingOfType("user.Email"), mock.AnythingOfType("user.Provider"), mock.AnythingOfType("user.ID")).
					Return(user.User{}, errors.New("error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Internal Server Error")
			},
		},
		"error creating tokens": {
			body: `{"provider": "test", "access_token": "validtoken", "expires_at": 1893456000, "provider_account_id": "12345"}`,
			buildStubs: func(mockUserModel *mocks.UserModel, mockAuthModel *mocks.AuthModel, mockProvider *mocks.Provider) {
				mockProvider.EXPECT().
					Validate("test", "validtoken").
					Return(string(expectedUser.Email), true, nil)

				mockUserModel.EXPECT().
					GetOrCreate(mock.Anything, mock.AnythingOfType("user.Email"), mock.AnythingOfType("user.Provider"), mock.AnythingOfType("user.ID")).
					Return(expectedUser, nil)

				mockAuthModel.EXPECT().
					CreateTokens(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return("", "", errors.New("error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Internal Server Error")
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mockProvider := mocks.NewProvider(t)
			mockRepo := mocks.NewRepository(t)
			mockUserModel := mocks.NewUserModel(t)
			mockAuthModel := mocks.NewAuthModel(t)

			tc.buildStubs(mockUserModel, mockAuthModel, mockProvider)

			s := server.New(
				mockRepo,
				mockUserModel,
				mockAuthModel,
				mockProvider,
				nil,
				server.Config{
					Limiter:    api.NewLimiter(25, true),
					Env:        "development",
					Version:    "1.0",
					AuthSecret: "testsecret",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			s.LoginHandler(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			tc.expectedFunc(rec)
		})
	}
}

type ResponseWriter struct {
	httptest.ResponseRecorder
	Fail bool
}

func (w *ResponseWriter) WriteHeader(code int) {
	if w.Fail {
		w.ResponseRecorder.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.ResponseRecorder.WriteHeader(code)
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	if w.Fail {
		return 0, errors.New("error")
	}

	return w.ResponseRecorder.Write(b)
}

func TestLoginHandler_ErrorResponse(t *testing.T) {
	mockProvider := mocks.NewProvider(t)
	mockRepo := mocks.NewRepository(t)
	mockUserModel := mocks.NewUserModel(t)
	mockAuthModel := mocks.NewAuthModel(t)

	mockProvider.EXPECT().
		Validate("test", "validtoken").
		Return(string("test@example.com"), true, nil)

	mockUserModel.EXPECT().
		GetOrCreate(mock.Anything, mock.AnythingOfType("user.Email"), mock.AnythingOfType("user.Provider"), mock.AnythingOfType("user.ID")).
		Return(user.User{}, nil)

	mockAuthModel.EXPECT().
		CreateTokens(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return("valid_access_token", "valid_refresh_token", nil)

	s := server.New(
		mockRepo,
		mockUserModel,
		mockAuthModel,
		mockProvider,
		nil,
		server.Config{
			Limiter:    api.NewLimiter(25, true),
			Env:        "development",
			Version:    "1.0",
			AuthSecret: "testsecret",
		},
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/auth/login",
		strings.NewReader(
			`{"provider": "test", "access_token": "validtoken", "expires_at": 1893456000, "provider_account_id": "12345"}`,
		),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := &ResponseWriter{ResponseRecorder: *httptest.NewRecorder(), Fail: true}

	s.LoginHandler(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestRefreshTokenHandler(t *testing.T) {
	tests := map[string]struct {
		header         string
		buildStubs     func(mockAuthModel *mocks.AuthModel)
		expectedStatus int
		expectedFunc   func(rec *httptest.ResponseRecorder)
	}{
		"successful refresh token": {
			header: "Bearer valid_access_token",
			buildStubs: func(mockAuthModel *mocks.AuthModel) {
				mockAuthModel.EXPECT().
					GetBearerToken(mock.Anything).
					Return("valid_access_token", nil)

				mockAuthModel.EXPECT().IsTokenRevoked(mock.Anything, mock.AnythingOfType("string")).
					Return(false, nil)

				mockAuthModel.EXPECT().
					RefreshToken(mock.Anything, mock.AnythingOfType("string")).
					Return("valid_access_token", nil)
			},
			expectedStatus: http.StatusOK,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "access_token")
			},
		},
		"error getting bearer token": {
			header: "Bearer invalid_token",
			buildStubs: func(mockAuthModel *mocks.AuthModel) {
				mockAuthModel.EXPECT().
					GetBearerToken(mock.Anything).
					Return("", auth.ErrNoAuthHeaderIncluded)
			},
			expectedStatus: http.StatusBadRequest,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Bad Request")
			},
		},
		"error getting status of token": {
			header: "Bearer valid_access_token",
			buildStubs: func(mockAuthModel *mocks.AuthModel) {
				mockAuthModel.EXPECT().
					GetBearerToken(mock.Anything).
					Return("valid_access_token", nil)

				mockAuthModel.EXPECT().IsTokenRevoked(mock.Anything, mock.AnythingOfType("string")).
					Return(true, errors.New("error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Internal Server Error")
			},
		},
		"token revoked": {
			header: "Bearer valid_access_token",
			buildStubs: func(mockAuthModel *mocks.AuthModel) {
				mockAuthModel.EXPECT().
					GetBearerToken(mock.Anything).
					Return("valid_access_token", nil)

				mockAuthModel.EXPECT().IsTokenRevoked(mock.Anything, mock.AnythingOfType("string")).
					Return(true, nil)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Unauthorized")
			},
		},
		"error refreshing token": {
			header: "Bearer valid_access_token",
			buildStubs: func(mockAuthModel *mocks.AuthModel) {
				mockAuthModel.EXPECT().
					GetBearerToken(mock.Anything).
					Return("valid_access_token", nil)

				mockAuthModel.EXPECT().
					IsTokenRevoked(mock.Anything, mock.AnythingOfType("string")).
					Return(false, nil)

				mockAuthModel.EXPECT().
					RefreshToken(mock.Anything, mock.AnythingOfType("string")).
					Return("", errors.New("error"))
			},
			expectedStatus: http.StatusUnauthorized,
			expectedFunc: func(rec *httptest.ResponseRecorder) {
				assert.Contains(t, rec.Body.String(), "Unauthorized")
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mockProvider := mocks.NewProvider(t)
			mockRepo := mocks.NewRepository(t)
			mockUserModel := mocks.NewUserModel(t)
			mockAuthModel := mocks.NewAuthModel(t)

			tc.buildStubs(mockAuthModel)

			s := server.New(
				mockRepo,
				mockUserModel,
				mockAuthModel,
				mockProvider,
				nil,
				server.Config{
					Limiter:    api.NewLimiter(25, true),
					Env:        "development",
					Version:    "1.0",
					AuthSecret: "testsecret",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/auth/refresh", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", tc.header)

			rec := httptest.NewRecorder()

			s.RefreshTokenHandler(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			tc.expectedFunc(rec)
		})
	}
}

func TestRefreshTokenHandler_ErrorResponse(t *testing.T) {
	mockProvider := mocks.NewProvider(t)
	mockRepo := mocks.NewRepository(t)
	mockUserModel := mocks.NewUserModel(t)
	mockAuthModel := mocks.NewAuthModel(t)

	mockAuthModel.EXPECT().
		GetBearerToken(mock.Anything).
		Return("valid_access_token", nil)

	mockAuthModel.EXPECT().IsTokenRevoked(mock.Anything, mock.AnythingOfType("string")).
		Return(false, nil)

	mockAuthModel.EXPECT().
		RefreshToken(mock.Anything, mock.AnythingOfType("string")).
		Return("valid_access_token", nil)

	s := server.New(
		mockRepo,
		mockUserModel,
		mockAuthModel,
		mockProvider,
		nil,
		server.Config{
			Limiter:    api.NewLimiter(25, true),
			Env:        "development",
			Version:    "1.0",
			AuthSecret: "testsecret",
		},
	)

	req := httptest.NewRequest(http.MethodPost, "/auth/refresh", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid_access_token")
	rec := &ResponseWriter{ResponseRecorder: *httptest.NewRecorder(), Fail: true}

	s.RefreshTokenHandler(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
