package user

import (
	"time"

	"github.com/nishojib/ffxivdailies/internal/validator"
)

type Provider string

type Account struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete"`
	Version   int64
	UserID    int64

	Provider          Provider
	ProviderAccountID ID
	Email             Email
}

func (a *Account) Validate(v *validator.Validator) {
	v.Check(a.Provider != "", "provider", "must be provided")
	v.Check(
		a.Provider == "google" || a.Provider == "discord",
		"provider",
		"must be either google or discord",
	)

	v.Check(a.Email != "", "email", "must be provided")
	v.Check(
		validator.Matches(string(a.Email), validator.EmailRX),
		"email",
		"must be a valid email address",
	)

	v.Check(a.ProviderAccountID != "", "provider_account_id", "must be provided")
}
