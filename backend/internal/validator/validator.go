package validator

import (
	"net/url"
	"regexp"
	"slices"
)

// EmailRX is a regular expression to validate email addresses.
var (
	EmailRX = regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
	)
)

// Validator is a struct that contains an Errors map.
type Validator struct {
	Errors map[string]string
}

// New returns a new Validator.
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid returns true if the item is valid (i.e., the errors map is empty).
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds an error message to the errors map.
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the errors map if the condition is false.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// PermittedValue returns true if the value is in the permitted values.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}

// Matches returns true if the value matches the regular expression.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique returns true if all the values in the slice are unique.
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}

func Url(value string) bool {
	u, err := url.Parse(value)
	return err == nil && u.Scheme != "" && u.Host != ""

}
