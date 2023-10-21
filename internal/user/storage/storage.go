package storage

import (
	"errors"

	"github.com/LNMMusic/msauth/internal/user"
)

var (
	// ErrStorageNotFound is returned when a user is not found
	ErrStorageNotFound = errors.New("storage: user not found")
	// ErrStorageExists is returned when a user already exists
	ErrStorageExists   = errors.New("storage: user already exists")
	// ErrStorageInvalid is returned when a user is invalid
	ErrStorageInvalid  = errors.New("storage: invalid user")
	// ErrStorageEncrypt is returned when a user password cannot be encrypted
	ErrStorageEncrypt  = errors.New("storage: cannot encrypt password")
)

// StorageRead interface for users to handle user read operations in the database
type StorageRead interface {
	// Get a user by id
	Get(id int) (u user.User, err error)
	// Get a user by email
	GetByEmail(email string) (u user.User, err error)
	// Get a user by username
	GetByUsername(username string) (u user.User, err error)
}

// StorageWrite interface for users to handle user operations in the database
type StorageWrite interface {
	// Create a new user
	Create(u *user.User) (err error)
	// Update an existing user
	Update(u *user.User) (err error)
	// Delete an existing user
	Delete(id int) (err error)
	// Activate an existing user
	Activate(id int) (err error)
}