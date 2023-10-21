package storage

import (
	"github.com/LNMMusic/msauth/internal/user"

	"github.com/stretchr/testify/mock"
)

// NewStorageReadMock returns a new mock storage
func NewStorageReadMock() *StorageReadMock {
	return &StorageReadMock{}
}

// StorageReadMock is the implementation of the Storage interface
type StorageReadMock struct {
	// mock.Mock is a struct that implements the Mock struct for testing purposes
	mock.Mock
}

// Get a user by id
func (m *StorageReadMock) Get(id string) (u user.User, err error) {
	args := m.Called(id)
	u = args.Get(0).(user.User)
	err = args.Error(1)
	return
}

// Get a user by email
func (m *StorageReadMock) GetByEmail(email string) (u user.User, err error) {
	args := m.Called(email)
	u = args.Get(0).(user.User)
	err = args.Error(1)
	return
}

// Get a user by username
func (m *StorageReadMock) GetByUsername(username string) (u user.User, err error) {
	args := m.Called(username)
	u = args.Get(0).(user.User)
	err = args.Error(1)
	return
}