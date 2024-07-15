package user

import (
	"time"

	"github.com/nishojib/ffxivdailies/internal/validator"
)

// ID is the identifier for a user.
type ID string

// Name is a short string that represents the name of a user.
type Name string

// Email is the email address of a user.
type Email string

// Image is the URL of the user's profile image.
type Image string

// User represents a user in the database.
type User struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete"`
	Version   int64

	UserID ID
	Name   Name
	Email  Email
	Image  Image
}

func (u *User) Validate(v *validator.Validator) {
	v.Check(u.UserID != "", "user_id", "must be provided")

	v.Check(u.Name != "", "name", "must be provided")
	v.Check(len(u.Name) <= 500, "name", "must not be more than 500 bytes long")

	v.Check(u.Email != "", "email", "must be provided")
	v.Check(
		validator.Matches(string(u.Email), validator.EmailRX),
		"email",
		"must be a valid email address",
	)

	v.Check(validator.Url(string(u.Image)), "image", "must be a valid url")
}
