package jwtauth

import "github.com/stretchr/testify/mock"

// constructor
func NewJWTAuthMock() *JWTAuthMock {
	return &JWTAuthMock{}
}


// JWTAuthMock is an implementation of JWTAuth interface for mock
type JWTAuthMock struct {
	mock.Mock
}

func (j *JWTAuthMock) GenerateSign(token *Token) (sign string, err error) {
	args := j.Called(token)
	sign = args.String(0)
	err = args.Error(1)
	return
}

func (j *JWTAuthMock) ValidateSign(sign string) (token *Token, err error) {
	args := j.Called(sign)
	token = args.Get(0).(*Token)
	err = args.Error(1)
	return
}