package crypter

import "github.com/stretchr/testify/mock"

// NewCrypterMock creates a new mock for the crypter interface
func NewCrypterMock() *CrypterMock {
	return &CrypterMock{}
}

// CrypterMock is the mock for the crypter interface
type CrypterMock struct {
	mock.Mock
}

// Encrypt is the mock for the Encrypt method
func (c *CrypterMock) Encrypt(s string) (e string, err error) {
	args := c.Called(s)
	e = args.String(0)
	err = args.Error(1)
	return
}

// Compare is the mock for the Compare method
func (c *CrypterMock) Compare(s, e string) (err error) {
	args := c.Called(s, e)
	err = args.Error(0)
	return
}