package storage

import (
	"fmt"

	"github.com/LNMMusic/msauth/internal/session"
)

// constructor
func NewStorageLocal(db map[string][]*session.Session) *StorageLocal {
	return &StorageLocal{db: db}
}

// StorageLocal is a local implementation of Storage interface
type StorageLocal struct {
	db map[string][]*session.Session
}

// Get returns all sessions for a user
func (s *StorageLocal) Get(userId string) (sessions []*session.Session, err error) {
	var ok bool
	sessions, ok = s.db[userId]
	if !ok {
		err = fmt.Errorf("%w: %s", ErrStorageUserNotFound, userId)
		return
	}

	return
}

// Set sets sessions for a user
func (s *StorageLocal) Set(userId string, sessions []*session.Session) (err error) {
	s.db[userId] = sessions
	return
}
