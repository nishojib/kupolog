package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nishojib/ffxivdailies/internal/auth"
)

func (r *Repository) IsTokenRevoked(ctx context.Context, token string) (bool, error) {
	var revocation auth.Revocation
	err := r.db.NewSelect().
		Model(&revocation).
		Where("token = ?", token).
		Scan(ctx)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func (r *Repository) RevokeToken(ctx context.Context, token string) error {
	revocation := auth.Revocation{
		Token:     token,
		RevokedAt: time.Now(),
	}

	if _, err := r.db.NewInsert().Model(&revocation).Exec(ctx); err != nil {
		return err
	}

	return nil
}
