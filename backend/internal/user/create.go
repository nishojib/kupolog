package user

import (
	"context"
	"errors"
	"time"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nrednav/cuid2"
)

// GetOrCreate returns a user if exists or creates a new one.
func (um *UserModel) GetOrCreate(
	ctx context.Context,
	email Email,
	provider Provider,
	accountID ID,
) (User, error) {
	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err := um.creator.GetUserByProviderID(dbCtx, string(accountID))
	if err != nil {
		if errors.Is(err, repoErrors.ErrRecordNotFound) {
			user = User{
				Name:   "Warrior of Light",
				Email:  email,
				Image:  "https://example.com/image.png",
				UserID: ID(cuid2.Generate()),
			}

			err = um.creator.InsertAndLinkAccount(dbCtx, &user, &Account{
				Provider:          provider,
				ProviderAccountID: accountID,
				Email:             email,
			})
			if err != nil {
				return User{}, err
			}
		} else {
			return User{}, err
		}
	}

	return user, nil
}

//go:generate mockery --with-expecter --name UserCreator
type UserCreator interface {
	GetUserByProviderID(ctx context.Context, providerAccountID string) (User, error)
	InsertAndLinkAccount(ctx context.Context, user *User, account *Account) error
}
