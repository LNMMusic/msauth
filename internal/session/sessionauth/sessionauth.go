package sessionauth

import (
	"errors"

	"github.com/LNMMusic/msauth/internal/session"
)

var (
	ErrSessionAuthManagerInternal       = errors.New("internal session auth manager error")
	ErrSessionReachedMaxSessionsPerUser = errors.New("max sessions per user reached")
	ErrSessionAuthManagerUnauthorized   = errors.New("unauthorized session")
)

// SessionAuthManager: an interface for auth to handle authentication operations for users sessions
// - it handles sync for expired sessions
type SessionAuthManager interface {
	// GenerateSession generates a new session for a user
	GenerateSession(userId string, s *session.Session) (err error)

	// ValidateSession validates a session for a user
	ValidateSession(userId string, tokenId string) (err error)
}