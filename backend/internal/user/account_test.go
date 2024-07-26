package user_test

import (
	"testing"

	"github.com/nishojib/ffxivdailies/internal/user"
	"github.com/nishojib/ffxivdailies/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestAccountValidate(t *testing.T) {
	tests := map[string]struct {
		account        user.Account
		expectedErrors map[string]string
		expectedValid  bool
	}{
		"valid account": {
			account: user.Account{
				Provider:          "google",
				ProviderAccountID: "123456",
				Email:             "example@example.com",
			},
			expectedErrors: map[string]string{},
			expectedValid:  true,
		},
		"invalid account": {
			account: user.Account{
				Provider:          "",
				ProviderAccountID: "",
				Email:             "",
			},
			expectedErrors: map[string]string{
				"provider":            "must be provided",
				"email":               "must be provided",
				"provider_account_id": "must be provided",
			},
			expectedValid: false,
		},
		"invalid account with invalid provider and email": {
			account: user.Account{
				Provider:          "myownprovider",
				Email:             "example",
				ProviderAccountID: "123456",
			},
			expectedErrors: map[string]string{
				"provider": "must be either google or discord",
				"email":    "must be a valid email address",
			},
			expectedValid: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v := validator.New()

			tc.account.Validate(v)

			assert.Equal(t, v.Valid(), tc.expectedValid)
			assert.Equal(t, tc.expectedErrors, v.Errors)
		})
	}
}
