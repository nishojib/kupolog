package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nishojib/ffxivdailies/internal/data/models"
	"github.com/uptrace/bun"
)

type RevocationRepository struct {
	db *bun.DB
}

func NewRevocationRepository(db *bun.DB) *RevocationRepository {
	return &RevocationRepository{db}
}

func (rr *RevocationRepository) IsTokenRevoked(token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var revocation models.Revocation
	err := rr.db.NewSelect().
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

	if revocation.RevokedAt.IsZero() {
		return false, nil
	}

	return true, nil
}

func (rr *RevocationRepository) RevokeToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	revocation := models.Revocation{
		Token:     token,
		RevokedAt: time.Now(),
	}

	if _, err := rr.db.NewInsert().Model(&revocation).Exec(ctx); err != nil {
		return err
	}

	return nil
}
