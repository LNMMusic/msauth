package crypter

import "errors"

var (
	// ErrCrypterEncryption is returned when an encryption error occurs
	ErrCrypterEncryption = errors.New("crypter: encryption error")
	// ErrCrypterComparison is returned when a comparison error occurs
	ErrCrypterComparison = errors.New("crypter: comparison error")
)

// Crypter interface for encrypting and decrypting strings
type Crypter interface {
	// Encrypt a string
	Encrypt(s string) (e string, err error)

	// Compare an encrypted string with a plain string and
	// validate if they are the same
	Compare(e string, s string) (err error)
}