package auth

import (
	"context"
	"time"
)

func (am *AuthModel) IsTokenRevoked(
	ctx context.Context,
	token string,
) (bool, error) {
	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return am.tokenRevoker.IsTokenRevoked(dbCtx, token)
}

//go:generate mockery --with-expecter --name TokenRevoker
type TokenRevoker interface {
	IsTokenRevoked(ctx context.Context, token string) (bool, error)
}
