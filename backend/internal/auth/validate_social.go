package auth

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func ValidateGoogle(token string) (string, bool, error) {
	url := "https://oauth2.googleapis.com/tokeninfo?access_token=" + token
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", false, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("failed to get token info for google", "error", err)
		return "", false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read token info for google", "error", err)
		return "", false, err
	}

	var input struct {
		Email   string `json:"email"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &input); err != nil {
		slog.Error("failed to unmarshal social info", "error", err)
		return "", false, err
	}

	if input.Message != "" {
		slog.Error("invalid message", "message", input.Message)
		return "", false, nil
	}

	return input.Email, true, nil
}

func ValidateDiscord(token string) (string, bool, error) {
	url := "https://discord.com/api/v10/users/@me"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		slog.Error("failed to create request for discord", "error", err)
		return "", false, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("failed to get user info for discord", "error", err)
		return "", false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read user info for discord", "error", err)
		return "", false, err
	}

	var input struct {
		Email   string `json:"email"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &input); err != nil {
		slog.Error("failed to unmarshal social info", "error", err)
		return "", false, err
	}

	if input.Message != "" {
		slog.Error("invalid message", "message", input.Message)
		return "", false, nil
	}

	return input.Email, true, nil
}
