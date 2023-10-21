package credential

import "errors"

var (
	// ErrCredentialInternal is returned when an internal error occurs
	ErrCredentialInternal = errors.New("credential internal error")
	// ErrCredentialUsernameNotFound is returned when the username is not found
	ErrCredentialUsernameNotFound = errors.New("credential username not found")
	// ErrCredentialEmailNotFound is returned when the email is not found
	ErrCredentialEmailNotFound = errors.New("credential email not found")
	// ErrCredentialPasswordInvalid is returned when the password is invalid
	ErrCredentialPasswordInvalid = errors.New("credential password invalid")
)

// Credential interface for verifying credentials
type Credential interface {
	// VerifyByUsername verifies a credential by username
	VerifyByUsername(username string, password string) (err error)

	// VerifyByEmail verifies a credential by email
	VerifyByEmail(email string, password string) (err error)
}
