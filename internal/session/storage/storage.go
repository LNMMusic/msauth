package storage

import (
	"errors"

	"github.com/LNMMusic/msauth/internal/session"
)

var (
	ErrStorageInternal	   = errors.New("internal storage error")
	ErrStorageUserNotFound = errors.New("user id not found")
)

// Storage is an interface for auth to handle auth get and set operations in the database for users sessions.
// - key-value storage: key is the user id, value is an slice of sessions
// - it does not handle sync for expired sessions
type Storage interface {
	// Get returns all sessions for a user
	Get(userId string) (s []*session.Session, err error)

	// Set sets sessions for a user
	Set(userId string, session []*session.Session) (err error)
}
