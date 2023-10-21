package credential

import (
	"errors"

	"github.com/LNMMusic/msauth/internal/user/storage"
	"github.com/LNMMusic/msauth/pkg/crypter"
)

// NewCredentialDefault returns a new default credential
func NewCredentialDefault(st storage.StorageRead, cr crypter.Crypter) *CredentialDefault {
	return &CredentialDefault{
		st: st,
		cr: cr,
	}
}

// CredentialDefault is the implementation of the Credential interface
type CredentialDefault struct {
	// st is the StorageRead interface
	st storage.StorageRead
	// cr is the Crypter interface
	cr crypter.Crypter
}

// VerifyByUsername verifies a credential by username
func (c *CredentialDefault) VerifyByUsername(username string, password string) (err error) {
	// get user from storage
	u, err := c.st.GetByUsername(username)
	if err != nil {
		if errors.Is(err, storage.ErrStorageNotFound) {
			err = ErrCredentialUsernameNotFound
		}

		return
	}

	// check if password is set
	if !u.Password.IsSome() {
		err = ErrCredentialInternal
		return
	}
	passwordUser, _ := u.Password.Unwrap()

	// check if password matches
	err = c.cr.Compare(passwordUser, password)
	if err != nil {
		err = ErrCredentialPasswordInvalid
		return
	}

	return
}

// VerifyByEmail verifies a credential by email
func (c *CredentialDefault) VerifyByEmail(email string, password string) (err error) {
	// get user from storage
	u, err := c.st.GetByEmail(email)
	if err != nil {
		if errors.Is(err, storage.ErrStorageNotFound) {
			err = ErrCredentialEmailNotFound
		}

		return
	}

	// check if password is set
	if !u.Password.IsSome() {
		err = ErrCredentialInternal
		return
	}
	passwordUser, _ := u.Password.Unwrap()

	// check if password matches
	err = c.cr.Compare(passwordUser, password)
	if err != nil {
		err = ErrCredentialPasswordInvalid
		return
	}

	return
}