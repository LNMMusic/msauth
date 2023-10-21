package validator

import (
	"github.com/LNMMusic/msauth/internal/user"

	"github.com/stretchr/testify/mock"
)

// constructor
func NewValidatorMock() *ValidatorMock {
	return &ValidatorMock{
		MethodDefault: func(u *user.User) {},
		MethodValidate: func(u *user.User) {},
		MethodPrepare: func(u *user.User) {},
	}
}

// ValidatorMock returns a new mock validator
type ValidatorMock struct {
	mock.Mock
	MethodDefault func(u *user.User)
	MethodValidate func(u *user.User)
	MethodPrepare func(u *user.User)
}

func (m *ValidatorMock) Default(u *user.User) (err error) {
	args := m.Called(u)
	
	m.MethodDefault(u)
	
	err = args.Error(0)
	return
}

func (m *ValidatorMock) Validate(u *user.User) (err error) {
	args := m.Called(u)
	
	m.MethodValidate(u)
	
	err = args.Error(0)
	return
}

func (m *ValidatorMock) Prepare(u *user.User) (err error) {
	args := m.Called(u)
	
	m.MethodPrepare(u)

	err = args.Error(0)
	return
}