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
	db userCreator,
	email Email,
	provider Provider,
	accountID ID,
) (User, error) {
	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err := db.GetUserByProviderID(dbCtx, string(accountID))
	if err != nil {
		if errors.Is(err, repoErrors.ErrRecordNotFound) {
			user = User{
				Name:   "Warrior of Light",
				Email:  email,
				Image:  "https://example.com/image.png",
				UserID: ID(cuid2.Generate()),
			}

			account := Account{
				Provider:          provider,
				ProviderAccountID: accountID,
				Email:             email,
			}

			err = db.InsertAndLinkAccount(dbCtx, &user, &account)
			if err != nil {
				return User{}, err
			}
		} else {
			return User{}, err
		}
	}

	return user, nil
}

//go:generate minimock -i userCreator -s "_mock.go" -o "mocks"
type userCreator interface {
	GetUserByProviderID(ctx context.Context, providerAccountID string) (User, error)
	InsertAndLinkAccount(ctx context.Context, user *User, account *Account) error
}
