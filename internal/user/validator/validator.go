package validator

import (
	"errors"

	"github.com/LNMMusic/msauth/internal/user"
)

var (
	// ErrValidatorDefault is returned when a user cannot be defaulted
	ErrValidatorDefault  = errors.New("validator: cannot default user")

	// ErrValidatorFieldRequired is returned when a field is required
	ErrValidatorFieldRequired  = errors.New("validator: field required")

	// ErrValidatorFieldQuality is returned when a field is not of the required quality
	ErrValidatorFieldQuality  = errors.New("validator: field quality")

	// ErrValidatorEncryption is returned when a field cannot be encrypted
	ErrValidatorEncryption  = errors.New("validator: cannot encrypt field")
)

// Validator interface for users to handle user validation before storage
// - handles some default values to certain fields and
// - handles validation of required fields and its quality
// other fields that require validation with the consistency on the database (as id, unique fields, etc) are handled by the db manager
type Validator interface {
	// Default set default values for a user in case of fields being none
	Default(u *user.User) (err error)

	// Validate validates required fields and its quality
	Validate(u *user.User) (err error)

	// Prepare prepares some fields for storage
	Prepare(u *user.User) (err error)
}