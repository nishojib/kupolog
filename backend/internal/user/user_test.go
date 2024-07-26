package user_test

import (
	"testing"

	"github.com/nishojib/ffxivdailies/internal/user"
	"github.com/nishojib/ffxivdailies/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	tests := map[string]struct {
		user           user.User
		expectedErrors map[string]string
		expectedValid  bool
	}{
		"valid user": {
			user: user.User{
				UserID: "123456",
				Name:   "Warrior of Light",
				Email:  "test@example.com",
				Image:  "https://example.com/image.png",
			},
			expectedErrors: map[string]string{},
			expectedValid:  true,
		},
		"invalid user": {
			user: user.User{
				UserID: "",
				Name:   "",
				Email:  "",
				Image:  "example",
			},
			expectedErrors: map[string]string{
				"user_id": "must be provided",
				"name":    "must be provided",
				"email":   "must be provided",
				"image":   "must be a valid url",
			},
			expectedValid: false,
		},
		"invalid user with invalid name and email": {
			user: user.User{
				UserID: "123456",
				Name:   "LongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongNameLongName",
				Email:  "example",
				Image:  "https://example.com/image.png",
			},
			expectedErrors: map[string]string{
				"name":  "must not be more than 500 bytes long",
				"email": "must be a valid email address",
			},
			expectedValid: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v := validator.New()

			tc.user.Validate(v)

			assert.Equal(t, v.Valid(), tc.expectedValid)
			assert.Equal(t, tc.expectedErrors, v.Errors)
		})
	}
}
