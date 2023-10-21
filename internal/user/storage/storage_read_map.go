package storage

import (
	"fmt"
	"sync"

	"github.com/LNMMusic/msauth/internal/user"
)

// NewStorageReadMap returns a new map storage
func NewStorageReadMap(db *sync.Map) *StorageReadMap {
	// default config
	defaultDb := &sync.Map{}
	if db != nil {
		defaultDb = db
	}

	return &StorageReadMap{defaultDb}
}

// StorageReadMap is the implementation of the Storage interface
type StorageReadMap struct {
	// db is a map of users
	// - key: user id
	// - value: user
	db *sync.Map
}

// Get a user by id
func (m *StorageReadMap) Get(id int) (u user.User, err error) {
	// get user from db
	value, ok := m.db.Load(id)
	if !ok {
		err = fmt.Errorf("%w - %d", ErrStorageNotFound, id)
		return
	}

	// type assertion
	u = value.(user.User)

	return
}

// Get a user by email
func (m *StorageReadMap) GetByEmail(email string) (u user.User, err error) {
	// loop over users
	var exists bool; var key int
	m.db.Range(func(k, v any) bool {
		// check if email matches
		userMap := v.(user.User)

		if userMap.Email.IsSome() {
			userMapEmail, _ := userMap.Email.Unwrap()
			if userMapEmail == email {
				exists = true
				key = k.(int)
				return false
			}
		}

		return true
	})

	// check if user exists
	if !exists {
		err = fmt.Errorf("%w - %s", ErrStorageNotFound, email)
		return
	}

	// get user from db
	value, _ := m.db.Load(key)

	// type assertion
	u = value.(user.User)

	return
}

// Get a user by username
func (m *StorageReadMap) GetByUsername(username string) (u user.User, err error) {
	// loop over users
	var exists bool; var key int
	m.db.Range(func(k, v any) bool {
		// check if username matches
		userMap := v.(user.User)

		if userMap.Username.IsSome() {
			userMapUsername, _ := userMap.Username.Unwrap()
			if userMapUsername == username {
				exists = true
				key = k.(int)
				return false
			}
		}

		return true
	})

	// check if user exists
	if !exists {
		err = fmt.Errorf("%w - %s", ErrStorageNotFound, username)
		return
	}

	// get user from db
	value, _ := m.db.Load(key)

	// type assertion
	u = value.(user.User)

	return
}