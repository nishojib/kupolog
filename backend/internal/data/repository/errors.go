package repository

import "errors"

var (
	ErrRecordNotFound     = errors.New("record not found")
	ErrEditConflict       = errors.New("edit conflict")
	ErrContextMissingUser = errors.New("context missing user")
	ErrServer             = errors.New("server error")
	ErrDuplicateEmail     = errors.New("duplicate email")
)
