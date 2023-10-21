package storage

import (
	"github.com/LNMMusic/msauth/internal/session"
	"github.com/stretchr/testify/mock"
)

// constructor
func NewStorageMock() *StorageMock {
	return &StorageMock{}
}

// StorageMock is a mock implementation of Storage
type StorageMock struct {
	mock.Mock
}

func (st *StorageMock) Get(userId string) (s []*session.Session, err error) {
	args := st.Called(userId)
	s = args.Get(0).([]*session.Session)
	err = args.Error(1)
	return
}

func (st *StorageMock) Set(userId string, session []*session.Session) (err error) {
	args := st.Called(userId, session)
	err = args.Error(0)
	return
}