package user

import (
	"github.com/LNMMusic/optional"
)

// Interfaces
type User struct {
	// Id is the unique identifier of the user
	Id			int
	// Username is the username of the user
	Username 	optional.Option[string]
	// Password is the password of the user
	// - is always hashed
	Password 	optional.Option[string]
	// Email is the email of the user
	Email 		optional.Option[string]
	// IsActive is the status of the user
	IsActive 	optional.Option[bool]
}