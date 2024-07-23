package user

import (
	"context"
	"errors"
	"time"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nrednav/cuid2"
)

func GetOrCreate(
	ctx context.Context,
	db UserCreator,
	email Email,
	provider Provider,
	accountID ID,
) (User, error) {
	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err := db.GetUserByProviderID(dbCtx, string(accountID))
	if err != nil {
		if errors.Is(err, repoErrors.ErrRecordNotFound) {
			user, err = db.InsertAndLinkAccount(dbCtx, User{
				Name:   "Warrior of Light",
				Email:  email,
				Image:  "https://example.com/image.png",
				UserID: ID(cuid2.Generate()),
			}, Account{
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
	InsertAndLinkAccount(ctx context.Context, user User, account Account) (User, error)
}
