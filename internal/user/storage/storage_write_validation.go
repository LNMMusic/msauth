package storage

import (
	"fmt"

	"github.com/LNMMusic/msauth/internal/user"
	"github.com/LNMMusic/msauth/internal/user/validator"
)

// NewStorageWriteValidation creates a new StorageWriteValidation
func NewStorageWriteValidation(st StorageWrite, vl validator.Validator) *StorageWriteValidation {
	return &StorageWriteValidation{st, vl}
}

// StorageWriteValidation wraps a StorageWrite interface and adds validation and encryption
type StorageWriteValidation struct {
	// st is the StorageWrite interface to be wrapped
	st StorageWrite
	// vl is the Validator interface to add validation to the StorageWrite interface
	vl validator.Validator
}


// Create a new user
func (s *StorageWriteValidation) Create(u *user.User) (err error) {
	// validator
	// - default values
	err = s.vl.Default(u)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInvalid, err.Error())
		return
	}
	// - validate user
	err = s.vl.Validate(u)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInvalid, err.Error())
		return
	}
	// - prepare user
	err = s.vl.Prepare(u)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInvalid, err.Error())
		return
	}

	// create user
	err = s.st.Create(u)
	return
}

// Update an existing user
func (s *StorageWriteValidation) Update(u *user.User) (err error) {
	// validator
	// - default values
	err = s.vl.Default(u)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInvalid, err.Error())
		return
	}
	// - validate user
	err = s.vl.Validate(u)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInvalid, err.Error())
		return
	}
	// - prepare user
	err = s.vl.Prepare(u)
	if err != nil {
		err = fmt.Errorf("%w. %s", ErrStorageInvalid, err.Error())
		return
	}

	// update user
	err = s.st.Update(u)
	return
}

// Delete an existing user
func (s *StorageWriteValidation) Delete(id int) (err error) {
	// delete user
	err = s.st.Delete(id)
	return
}

// Activate an existing user
func (s *StorageWriteValidation) Activate(id int) (err error) {
	// activate user
	err = s.st.Activate(id)
	return
}