package user_test

import (
	"context"
	"errors"
	"testing"
	"time"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nishojib/ffxivdailies/internal/user"
	"github.com/nishojib/ffxivdailies/internal/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetOrCreate(t *testing.T) {
	newUser := user.User{
		ID:        1,
		Name:      "Warrior of Light",
		Email:     "test@example.com",
		Image:     "https://example.com/image.png",
		UserID:    "123456",
		Version:   1,
		CreatedAt: time.Now(),
		DeletedAt: time.Time{},
	}

	newAccount := user.Account{
		ID:                1,
		Provider:          user.Provider("google"),
		ProviderAccountID: "123456",
		Email:             "test@example.com",
		Version:           1,
		CreatedAt:         time.Now(),
		DeletedAt:         time.Time{},
		UserID:            1,
	}

	dbError := errors.New("error")

	tests := map[string]struct {
		user        user.User
		account     user.Account
		db          func(creator *mocks.UserCreator, u *user.User)
		expectedErr error
	}{
		"create user when nothing exists": {
			user:    newUser,
			account: newAccount,
			db: func(creator *mocks.UserCreator, u *user.User) {
				creator.EXPECT().
					GetUserByProviderID(mock.Anything, mock.AnythingOfType("string")).
					Return(user.User{}, repoErrors.ErrRecordNotFound)

				creator.EXPECT().
					InsertAndLinkAccount(mock.Anything, mock.AnythingOfType("*user.User"), mock.AnythingOfType("*user.Account")).
					RunAndReturn(func(_ context.Context, usr *user.User, _ *user.Account) error {
						usr.ID = u.ID
						usr.UserID = u.UserID
						usr.Name = u.Name
						usr.Email = u.Email
						usr.Image = u.Image
						usr.CreatedAt = u.CreatedAt
						usr.DeletedAt = u.DeletedAt
						usr.Version = u.Version
						return nil
					})
			},
			expectedErr: nil,
		},
		"error when creating user": {
			user:    newUser,
			account: newAccount,
			db: func(creator *mocks.UserCreator, _ *user.User) {
				creator.EXPECT().
					GetUserByProviderID(mock.Anything, mock.AnythingOfType("string")).
					Return(user.User{}, repoErrors.ErrRecordNotFound)

				creator.EXPECT().
					InsertAndLinkAccount(mock.Anything, mock.AnythingOfType("*user.User"), mock.AnythingOfType("*user.Account")).
					Return(dbError)
			},
			expectedErr: dbError,
		},
		"getting existing user": {
			user:    newUser,
			account: newAccount,
			db: func(creator *mocks.UserCreator, u *user.User) {
				creator.EXPECT().
					GetUserByProviderID(mock.Anything, mock.AnythingOfType("string")).
					Return(*u, nil)
			},
			expectedErr: nil,
		},
		"error getting existing user": {
			user:    newUser,
			account: newAccount,
			db: func(creator *mocks.UserCreator, u *user.User) {
				creator.EXPECT().
					GetUserByProviderID(mock.Anything, mock.AnythingOfType("string")).
					Return(user.User{}, dbError)
			},
			expectedErr: dbError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mockUserCreator := mocks.NewUserCreator(t)

			tc.db(mockUserCreator, &tc.user)

			userModel := user.NewModel(mockUserCreator)

			got, err := userModel.GetOrCreate(
				context.Background(),
				tc.user.Email,
				tc.account.Provider,
				tc.account.ProviderAccountID,
			)
			require.ErrorIs(t, err, tc.expectedErr)

			if tc.expectedErr == nil {
				assert.Equal(t, tc.user, got)
			}

		})
	}
}
