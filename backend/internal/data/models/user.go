package models

import (
	"time"

	"github.com/nishojib/ffxivdailies/internal/validator"
)

// User represents a user in the database.
type User struct {
	ID        int64     `json:"-"          bun:"id,pk,autoincrement"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at" bun:",default:current_timestamp"`
	DeletedAt time.Time `json:"-"          bun:",soft_delete"`
	Version   int64     `json:"-"`
}

func (u *User) Validate(v *validator.Validator) {
	v.Check(u.UserID != "", "user_id", "must be provided")

	v.Check(u.Name != "", "name", "must be provided")
	v.Check(len(u.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(u.Email != "", "email", "must be provided")
	v.Check(validator.Matches(u.Email, validator.EmailRX), "email", "must be a valid email address")

	v.Check(validator.Url(u.Image), "image", "must be a valid url")
}
