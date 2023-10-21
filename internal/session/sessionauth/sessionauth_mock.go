package sessionauth

import (
	"github.com/LNMMusic/msauth/internal/session"
	"github.com/stretchr/testify/mock"
)

// constructor
func NewSessionAuthMock() *SessionAuthMock {
	return &SessionAuthMock{}
}

// SessionAuthMock is a mock implementation of SessionAuthManager interface
type SessionAuthMock struct {
	mock.Mock
}

func (m *SessionAuthMock) GenerateSession(userID string, session *session.Session) (err error) {
	args := m.Called(userID, session)
	err = args.Error(0)
	return
}

func (m *SessionAuthMock) ValidateSession(userID string, sessionID string) (err error) {
	args := m.Called(userID, sessionID)
	err = args.Error(0)
	return
}