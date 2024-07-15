package auth

import "time"

type Revocation struct {
	ID        string    `bun:"id,pk,autoincrement"`
	RevokedAt time.Time `bun:",default:current_timestamp"`
	Token     string
}
