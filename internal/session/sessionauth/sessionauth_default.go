package sessionauth

import (
	"fmt"
	"time"

	"github.com/LNMMusic/msauth/internal/session"
	"github.com/LNMMusic/msauth/internal/session/storage"
	"github.com/LNMMusic/optional"
)

// constructor
func NewSessionAuthManagerDefault(st storage.Storage, config *Config) *SessionAuthManagerDefault {
	// default values
	if !config.MaxSessionsPerUser.IsSome() {
		config.MaxSessionsPerUser = optional.Some(5)
	}

	return &SessionAuthManagerDefault{
		st: st,
		config:  config,
	}
}

// SessionAuthManagerDefault returns a new instance of the default implementation of SessionAuthManager
type Config struct {
	// MaxSessionsPerUser is the maximum number of sessions per user
	MaxSessionsPerUser 	optional.Option[int]
}

type SessionAuthManagerDefault struct {
	// st is the storage for sessions (get and set operations)
	st storage.Storage
	config  *Config
}

// GenerateSession generates a new session for a user
func (sa *SessionAuthManagerDefault) GenerateSession(userId string, s *session.Session) (err error) {
	// get all sessions for a user
	sessions, err := sa.st.Get(userId)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrSessionAuthManagerInternal, err.Error())
		return
	}

	// sync sessions
	var syncedSessions []*session.Session
	for _, session := range sessions {
		// not expired sessions
		if session.ExpireDate.After(time.Now()) {
			syncedSessions = append(syncedSessions, session)
		}
	}

	// check max sessions per user
	MaxSessionsPerUser, _ := sa.config.MaxSessionsPerUser.Unwrap()
	if len(syncedSessions) >= MaxSessionsPerUser {
		err = fmt.Errorf("%w. %d", ErrSessionReachedMaxSessionsPerUser, len(syncedSessions))
		return
	}

	// add new session
	syncedSessions = append(syncedSessions, s)

	// set sessions
	err = sa.st.Set(userId, syncedSessions)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrSessionAuthManagerInternal, err.Error())
		return
	}

	return
}

// ValidateSession validates a session for a user
func (sa *SessionAuthManagerDefault) ValidateSession(userId string, tokenId string) (err error) {
	// get all sessions for a user
	sessions, err := sa.st.Get(userId)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrSessionAuthManagerInternal, err.Error())
		return
	}

	// sync sessions
	var syncedSessions []*session.Session
	for _, session := range sessions {
		// not expired sessions
		if session.ExpireDate.After(time.Now()) {
			syncedSessions = append(syncedSessions, session)
		}
	}

	// check if session is valid
	var validSession bool
	for _, session := range syncedSessions {
		if session.TokenID == tokenId {
			validSession = true
			break
		}
	}

	if !validSession {
		err = fmt.Errorf("%w. %s", ErrSessionAuthManagerUnauthorized, tokenId)
		return
	}

	return
}