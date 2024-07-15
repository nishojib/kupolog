package validator_test

import (
	"testing"

	"github.com/nishojib/ffxivdailies/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	v := validator.New()
	assert.NotNil(t, v)
}

func TestValid(t *testing.T) {
	v := validator.New()

	assert.True(t, v.Valid())

	v.AddError("key", "error")
	assert.False(t, v.Valid())
}

func TestCheck(t *testing.T) {
	v := validator.New()

	v.Check(true, "key", "error")
	assert.True(t, v.Valid())

	v.Check(false, "key", "error")
	assert.False(t, v.Valid())
}

func TestPermittedValue(t *testing.T) {
	assert.True(t, validator.PermittedValue(1, 1, 2, 3))
	assert.False(t, validator.PermittedValue(1, 2, 3))
}

func TestMatches(t *testing.T) {
	assert.True(t, validator.Matches("test@example.com", validator.EmailRX))
	assert.False(t, validator.Matches("test", validator.EmailRX))
}

func TestUnique(t *testing.T) {
	assert.True(t, validator.Unique([]int{1, 2, 3}))
	assert.False(t, validator.Unique([]string{"test", "test", "test"}))
}

func TestUrl(t *testing.T) {
	assert.True(t, validator.Url("https://example.com"))
	assert.False(t, validator.Url("example.com"))
}
