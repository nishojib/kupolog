package repository

import "github.com/uptrace/bun"

type Repository struct {
	db *bun.DB
}

func New(db *bun.DB) *Repository {
	return &Repository{db}
}
