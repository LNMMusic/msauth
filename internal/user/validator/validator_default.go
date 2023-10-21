package validator

import (
	"fmt"
	"regexp"

	"github.com/LNMMusic/msauth/internal/user"
	"github.com/LNMMusic/msauth/pkg/crypter"

	"github.com/LNMMusic/optional"
)

// constructor
func NewValidatorDefault(emailRegex string, cr crypter.Crypter) (v *ValidatorDefault) {
	// default cfg
	defaultEmailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	if emailRegex != "" {
		defaultEmailRegex = emailRegex
	}

	return &ValidatorDefault{
		emailRegex: regexp.MustCompile(defaultEmailRegex),
		cr:         cr,
	}
}

// ValidatorDefault is the implementation of the Validator interface
type ValidatorDefault struct {
	// emailRegex is the regex used to validate emails
	emailRegex *regexp.Regexp

	// cr is the Crypter interface to encrypt fields
	cr crypter.Crypter
}

// Default sets default values for a user in case of fields being none
func (v *ValidatorDefault) Default(u *user.User) (err error) {
	// set default values
	// -> id is not set, as its not a default value (its handle by the System that manages the database)
	// ...

	// -> is_active is set to false by default (in case its a none value)
	if !u.IsActive.IsSome() {
		u.IsActive = optional.Some(false)
	}

	return
}

// Validate validates required fields and its quality
func (v *ValidatorDefault) Validate(u *user.User) (err error) {
	// check required fields
	if !u.Username.IsSome() {
		err = fmt.Errorf("%w - username is none", ErrValidatorFieldRequired)
		return
	}
	if !u.Password.IsSome() {
		err = fmt.Errorf("%w - password is none", ErrValidatorFieldRequired)
		return
	}
	if !u.Email.IsSome() {
		err = fmt.Errorf("%w - email is none", ErrValidatorFieldRequired)
		return
	}

	// check quality of fields
	username, _ := u.Username.Unwrap()
	if len(username) < 3 || len(username) > 25 {
		err = fmt.Errorf("%w - username, chars should be between 3 and 25", ErrValidatorFieldQuality)
		return
	}
	password, _ := u.Password.Unwrap()
	if len(password) < 8 || len(password) > 25 {
		err = fmt.Errorf("%w - password, chars should be between 8 and 25", ErrValidatorFieldQuality)
		return
	}
	email, _ := u.Email.Unwrap()
	if !v.emailRegex.MatchString(email) {
		err = fmt.Errorf("%w - email, invalid format", ErrValidatorFieldQuality)
		return
	}

	return
}

// Prepare prepares some fields for storage
func (v *ValidatorDefault) Prepare(u *user.User) (err error) {
	// encrypt password
	password, _ := u.Password.Unwrap()
	encryptedPassword, err := v.cr.Encrypt(password)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrValidatorEncryption, err)
		return
	}
	(*u).Password = optional.Some(encryptedPassword)

	return
}
	