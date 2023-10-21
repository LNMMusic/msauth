package crypter

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// NewCrypterDefault returns a new CrypterDefault
func NewCrypterDefault(cost int) *CrypterDefault {
	// default config
	defaultCost := bcrypt.DefaultCost
	if cost >= bcrypt.MinCost && cost <= bcrypt.MaxCost {
		defaultCost = cost
	}

	return &CrypterDefault{
		cost: defaultCost,
	}
}

// CrypterDefault is the implementation of the Crypter interface
type CrypterDefault struct {
	// cost is the cost of the bcrypt algorithm
	cost int
}

// Encrypt encrypts an string
func (c *CrypterDefault) Encrypt(s string) (e string, err error) {
	var b []byte
	b, err = bcrypt.GenerateFromPassword([]byte(s), c.cost)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrCrypterEncryption, err.Error())
		return
	}

	e = string(b)
	return
}

// Compare compares an encrypted string with a plain string
func (c *CrypterDefault) Compare(e string, s string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(e), []byte(s))
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrCrypterComparison, err.Error())
		return
	}

	return
}
