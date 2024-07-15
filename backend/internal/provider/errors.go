package provider

import "errors"

var (
	ErrRequestFailed  = errors.New("request failed")
	ErrMalformedInput = errors.New("malformed response")
	ErrReadTokenInfo  = errors.New("failed to read token info")
	ErrGetTokenInfo   = errors.New("failed to get token info")
	ErrMessage        = errors.New("invalid message")
)
