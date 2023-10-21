package session

import "time"

// Session is a struct that represents a user session
type Session struct {
	TokenID    string    `json:"token_id"`
	ExpireDate time.Time `json:"expire_date"`
}