package api

import (
	"errors"
)

// Environment represents the environment the application is running in.
type Environment string

const (
	// Development environment.
	EnvDevelopment = "development"
	// Production environment.
	EnvProduction = "production"
)

var (
	ErrEnvironment = errors.New("environment must be either development or production")
)

// NewEnvironment creates a new Environment from a string.
func NewEnvironment(env string) (Environment, error) {
	if env != EnvDevelopment && env != EnvProduction {
		return "", ErrEnvironment
	}

	return Environment(env), nil
}

// String returns the string representation of the Environment.
func (env Environment) String() string {
	return string(env)
}

// Limiter contains the rate limit configuration.
type Limiter struct {
	RPS     int
	Enabled bool
}

// NewLimiter creates a new Limiter.
func NewLimiter(rps int, enabled bool) Limiter {
	return Limiter{
		RPS:     rps,
		Enabled: enabled,
	}
}
