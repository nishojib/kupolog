package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type Provider struct {
}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) Validate(provider string, token string) (string, bool, error) {
	var req *http.Request
	var err error

	if provider == "google" {
		url := "https://oauth2.googleapis.com/tokeninfo?access_token=" + token
		req, err = http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			slog.Error("failed to create request for google", "error", err)
			return "", false, ErrRequestFailed
		}
	} else if provider == "discord" {
		url := "https://discord.com/api/v10/users/@me"
		req, err = http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			slog.Error("failed to create request for discord", "error", err)
			return "", false, ErrRequestFailed
		}
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get token info for %s", provider), "error", err)
		return "", false, ErrGetTokenInfo
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to read token info for %s", provider), "error", err)
		return "", false, ErrReadTokenInfo
	}

	var input struct {
		Email   string `json:"email"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &input); err != nil {
		slog.Error("failed to unmarshal social info", "error", err)
		return "", false, ErrMalformedInput
	}

	if input.Message != "" {
		slog.Error("invalid message", "message", input.Message)
		return "", false, ErrMessage
	}

	return input.Email, true, nil
}
