package storage

import (
	"github.com/LNMMusic/msauth/internal/user"
	"github.com/stretchr/testify/mock"
)

// NewStorageWriteMock returns a new mock storage
func NewStorageWriteMock() *StorageWriteMock {
	return &StorageWriteMock{}
}

// StorageWriteMock is a mock implementation of the Storage interface
type StorageWriteMock struct {
	// mock.Mock is a struct that implements the Mock struct for testing purposes
	mock.Mock
}

// Create a new user
func (m *StorageWriteMock) Create(u *user.User) (err error) {
	args := m.Called(u)
	err = args.Error(0)
	return
}

// Update an existing user
func (m *StorageWriteMock) Update(u *user.User) (err error) {
	args := m.Called(u)
	err = args.Error(0)
	return
}

// Delete an existing user
func (m *StorageWriteMock) Delete(id int) (err error) {
	args := m.Called(id)
	err = args.Error(0)
	return
}

// Activate an existing user
func (m *StorageWriteMock) Activate(id int) (err error) {
	args := m.Called(id)
	err = args.Error(0)
	return
}