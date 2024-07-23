package repository_test

import (
	"context"
	"testing"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nishojib/ffxivdailies/internal/repository"
	"github.com/nishojib/ffxivdailies/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

func TestInsertAndLinkAccount_CreateNewUser(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	u := user.User{
		UserID: "user-123456",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}

	a := user.Account{
		Provider:          "google",
		ProviderAccountID: "provider-123456",
		Email:             "john@doe.com",
	}

	got, err := repo.InsertAndLinkAccount(context.Background(), u, a)
	require.NoError(t, err)

	var expectedUser user.User
	err = bunDB.NewSelect().
		Model(&expectedUser).
		Where("user_id = ?", u.UserID).
		Scan(context.Background())
	require.NoError(t, err)

	assert.Equal(t, got, expectedUser)

	var expectedAccount user.Account
	err = bunDB.NewSelect().
		Model(&expectedAccount).
		Where("provider_account_id = ?", a.ProviderAccountID).
		Scan(context.Background())
	require.NoError(t, err)

	assert.Equal(t, got.ID, expectedAccount.UserID)
}

func TestInsertAndLinkAccount_UpdateExistingUser(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	u := user.User{
		UserID: "user-234567",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}

	_, err := bunDB.NewInsert().Model(&u).Exec(context.Background())
	require.NoError(t, err)

	newUser := user.User{
		UserID: "user-123456",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}

	a := user.Account{
		Provider:          "google",
		ProviderAccountID: "provider-123456",
		Email:             "john@doe.com",
	}

	got, err := repo.InsertAndLinkAccount(context.Background(), newUser, a)
	require.NoError(t, err)

	var expectedUser user.User
	err = bunDB.NewSelect().
		Model(&expectedUser).
		Where("user_id = ?", u.UserID).
		Scan(context.Background())
	require.NoError(t, err)

	assert.Equal(t, got, expectedUser)

	var expectedAccount user.Account
	err = bunDB.NewSelect().
		Model(&expectedAccount).
		Where("provider_account_id = ?", a.ProviderAccountID).
		Scan(context.Background())
	require.NoError(t, err)

	assert.Equal(t, got.ID, expectedAccount.UserID)
}

func TestInsertAndLinkAccount_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := repo.InsertAndLinkAccount(ctx, user.User{}, user.Account{})
	require.Error(t, err)
}

func TestGetUserByProviderID_Success(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	expected, err := repo.InsertAndLinkAccount(context.Background(), user.User{
		UserID: "user-123456",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}, user.Account{
		Provider:          "google",
		ProviderAccountID: "provider-123456",
		Email:             "john@doe.com",
	})
	require.NoError(t, err)

	got, err := repo.GetUserByProviderID(context.Background(), "provider-123456")
	require.NoError(t, err)

	assert.Equal(t, got, expected)
}

func TestGetUserByProviderID_NotFoundError(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	_, err := repo.GetUserByProviderID(context.Background(), "provider-123456")
	require.ErrorIs(t, err, repoErrors.ErrRecordNotFound)
}

func TestGetUserByProviderID_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	_, err := repo.InsertAndLinkAccount(context.Background(), user.User{
		UserID: "user-123456",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}, user.Account{
		Provider:          "google",
		ProviderAccountID: "provider-123456",
		Email:             "john@doe.com",
	})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = repo.GetUserByProviderID(ctx, "provider-123456")
	require.Error(t, err)
}

func TestUpdateUser_Success(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	u, err := repo.InsertAndLinkAccount(context.Background(), user.User{
		UserID: "user-123456",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}, user.Account{})
	require.NoError(t, err)

	updatedUser := user.User{
		ID:     u.ID,
		UserID: u.UserID,
		Name:   "Jane Doe",
		Email:  "jane@doe.com",
		Image:  u.Image,
	}

	err = repo.UpdateUser(context.Background(), &updatedUser)
	require.NoError(t, err)

	var gotUser user.User
	err = bunDB.NewSelect().
		Model(&gotUser).
		Where("user_id = ?", u.UserID).
		Scan(context.Background())
	require.NoError(t, err)

	assert.Equal(t, gotUser.Name, updatedUser.Name)
	assert.Equal(t, gotUser.Email, updatedUser.Email)
	assert.Equal(t, gotUser.Version, u.Version+1)
}

func TestUpdateUser_EditConflictError(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	u := user.User{
		UserID: "user-123456",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}

	err := repo.UpdateUser(context.Background(), &u)
	require.ErrorIs(t, err, repoErrors.ErrEditConflict)
}

func TestUpdateUser_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	_, err := repo.InsertAndLinkAccount(context.Background(), user.User{
		UserID: "user-123456",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}, user.Account{})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err = repo.UpdateUser(ctx, &user.User{})
	require.Error(t, err)
}

func TestGetByUserID_Success(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	u, err := repo.InsertAndLinkAccount(context.Background(), user.User{
		UserID: "user-123456",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}, user.Account{})
	require.NoError(t, err)

	gotUser, err := repo.GetUserByUserID(context.Background(), string(u.UserID))
	require.NoError(t, err)

	assert.Equal(t, u, gotUser)
}

func TestGetUserByUserID_NotFoundError(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	_, err := repo.GetUserByUserID(context.Background(), "user-123456")
	require.ErrorIs(t, err, repoErrors.ErrRecordNotFound)
}

func TestGetUserByUserID_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	_, err := repo.InsertAndLinkAccount(context.Background(), user.User{
		UserID: "user-123456",
		Name:   "John Doe",
		Email:  "john@doe.com",
		Image:  "https://example.com/image.png",
	}, user.Account{})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = repo.GetUserByUserID(ctx, "user-123456")
	require.Error(t, err)
}
