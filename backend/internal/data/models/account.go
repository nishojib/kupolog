package models

import (
	"time"

	"github.com/nishojib/ffxivdailies/internal/validator"
)

type Account struct {
	ID                int64     `json:"_"                 bun:"id,pk,autoincrement"`
	UserID            int64     `json:"userID"`
	Provider          string    `json:"provider"`
	ProviderAccountID string    `json:"providerAccountId"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"createdAt"         bun:",default:current_timestamp"`
	DeletedAt         time.Time `json:"deletedAt"         bun:",soft_delete"`
	Version           int64     `json:"version"`
}

func (a *Account) Validate(v *validator.Validator) {
	v.Check(a.Provider != "", "provider", "must be provided")
	v.Check(
		a.Provider == "google" || a.Provider == "discord",
		"provider",
		"must be either google or discord",
	)

	v.Check(a.Email != "", "email", "must be provided")
	v.Check(validator.Matches(a.Email, validator.EmailRX), "email", "must be a valid email address")

	v.Check(a.ProviderAccountID != "", "provider_account_id", "must be provided")
}
