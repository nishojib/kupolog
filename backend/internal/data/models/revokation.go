package models

import "time"

type Revocation struct {
	ID        string    `json:"id"        bun:"id,pk,autoincrement"`
	Token     string    `json:"token"`
	RevokedAt time.Time `json:"revokedAt" bun:",default:current_timestamp"`
}
