package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/nishojib/ffxivdailies/internal/auth"
	"github.com/nishojib/ffxivdailies/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

func TestIsTokenRevoked_NotRevoked_NotFound(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	isRevoked, err := repo.IsTokenRevoked(context.Background(), "test")
	require.NoError(t, err)

	assert.False(t, isRevoked)
}

func TestIsTokenRevoked_Revoked(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	revocation := auth.Revocation{
		RevokedAt: time.Now(),
		Token:     "test",
	}

	_, err := bunDB.NewInsert().Model(&revocation).Exec(context.Background())
	require.NoError(t, err)

	isRevoked, err := repo.IsTokenRevoked(context.Background(), revocation.Token)
	require.NoError(t, err)

	assert.True(t, isRevoked)
}

func TestIsTokenRevoked_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	revocation := auth.Revocation{
		RevokedAt: time.Now(),
		Token:     "test",
	}

	_, err := bunDB.NewInsert().Model(&revocation).Exec(context.Background())
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = repo.IsTokenRevoked(ctx, revocation.Token)
	require.Error(t, err)
}

func TestRevokeToken_Success(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	err := repo.RevokeToken(context.Background(), "test")
	require.NoError(t, err)

	isRevoked, err := repo.IsTokenRevoked(context.Background(), "test")
	require.NoError(t, err)

	assert.True(t, isRevoked)
}

func TestRevokeToken_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := repo.RevokeToken(ctx, "test")
	require.Error(t, err)
}
