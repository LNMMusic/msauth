package storage_test

import (
	"fmt"
	"testing"

	"github.com/LNMMusic/msauth/internal/user/storage"
	"github.com/LNMMusic/msauth/internal/user/validator"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Tests for StorageWriteValidation.Create
func TestStorageWriteValidation_Create(t *testing.T) {
	t.Run("success to create", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(nil)
		vlMk.On("Validate", mock.Anything).Return(nil)
		vlMk.On("Prepare", mock.Anything).Return(nil)

		// - storage: mock
		stMk := storage.NewStorageWriteMock()
		stMk.On("Create", mock.Anything).Return(nil)

		// - storage: validation
		st := storage.NewStorageWriteValidation(stMk, vlMk)

		// act
		err := st.Create(nil)

		// assert
		require.NoError(t, err)
		vlMk.AssertExpectations(t)
		stMk.AssertExpectations(t)
	})

	t.Run("fail to create - default error", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(validator.ErrValidatorDefault)

		// - storage: mock
		// ...

		// - storage: validation
		st := storage.NewStorageWriteValidation(nil, vlMk)

		// act
		err := st.Create(nil)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageInvalid)
		require.EqualError(t, err, fmt.Sprintf("%v. %v", storage.ErrStorageInvalid, validator.ErrValidatorDefault))
		vlMk.AssertExpectations(t)
	})

	t.Run("fail to create - validate error", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(nil)
		vlMk.On("Validate", mock.Anything).Return(validator.ErrValidatorFieldRequired)

		// - storage: mock
		// ...

		// - storage: validation
		st := storage.NewStorageWriteValidation(nil, vlMk)

		// act
		err := st.Create(nil)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageInvalid)
		require.EqualError(t, err, fmt.Sprintf("%v. %v", storage.ErrStorageInvalid, validator.ErrValidatorFieldRequired))
		vlMk.AssertExpectations(t)
	})

	t.Run("fail to create - prepare error", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(nil)
		vlMk.On("Validate", mock.Anything).Return(nil)
		vlMk.On("Prepare", mock.Anything).Return(validator.ErrValidatorEncryption)

		// - storage: mock
		// ...

		// - storage: validation
		st := storage.NewStorageWriteValidation(nil, vlMk)

		// act
		err := st.Create(nil)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageInvalid)
		require.EqualError(t, err, fmt.Sprintf("%v. %v", storage.ErrStorageInvalid, validator.ErrValidatorEncryption))
		vlMk.AssertExpectations(t)
	})

	t.Run("fail to create - storage error", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(nil)
		vlMk.On("Validate", mock.Anything).Return(nil)
		vlMk.On("Prepare", mock.Anything).Return(nil)

		// - storage: mock
		stMk := storage.NewStorageWriteMock()
		stMk.On("Create", mock.Anything).Return(storage.ErrStorageExists)

		// - storage: validation
		st := storage.NewStorageWriteValidation(stMk, vlMk)

		// act
		err := st.Create(nil)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageExists)
		vlMk.AssertExpectations(t)
		stMk.AssertExpectations(t)
	})
}

// Tests for StorageWriteValidation.Update
func TestStorageWriteValidation_Update(t *testing.T) {
	t.Run("success to update", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(nil)
		vlMk.On("Validate", mock.Anything).Return(nil)
		vlMk.On("Prepare", mock.Anything).Return(nil)

		// - storage: mock
		stMk := storage.NewStorageWriteMock()
		stMk.On("Update", mock.Anything).Return(nil)

		// - storage: validation
		st := storage.NewStorageWriteValidation(stMk, vlMk)

		// act
		err := st.Update(nil)

		// assert
		require.NoError(t, err)
		vlMk.AssertExpectations(t)
		stMk.AssertExpectations(t)
	})

	t.Run("fail to update - default error", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(validator.ErrValidatorDefault)

		// - storage: mock
		// ...

		// - storage: validation
		st := storage.NewStorageWriteValidation(nil, vlMk)

		// act
		err := st.Update(nil)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageInvalid)
		require.EqualError(t, err, fmt.Sprintf("%v. %v", storage.ErrStorageInvalid, validator.ErrValidatorDefault))
		vlMk.AssertExpectations(t)
	})

	t.Run("fail to update - validate error", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(nil)
		vlMk.On("Validate", mock.Anything).Return(validator.ErrValidatorFieldRequired)

		// - storage: mock
		// ...

		// - storage: validation
		st := storage.NewStorageWriteValidation(nil, vlMk)

		// act
		err := st.Update(nil)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageInvalid)
		require.EqualError(t, err, fmt.Sprintf("%v. %v", storage.ErrStorageInvalid, validator.ErrValidatorFieldRequired))
		vlMk.AssertExpectations(t)
	})

	t.Run("fail to update - prepare error", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(nil)
		vlMk.On("Validate", mock.Anything).Return(nil)
		vlMk.On("Prepare", mock.Anything).Return(validator.ErrValidatorEncryption)

		// - storage: mock
		// ...

		// - storage: validation
		st := storage.NewStorageWriteValidation(nil, vlMk)

		// act
		err := st.Update(nil)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageInvalid)
		require.EqualError(t, err, fmt.Sprintf("%v. %v", storage.ErrStorageInvalid, validator.ErrValidatorEncryption))
		vlMk.AssertExpectations(t)
	})

	t.Run("fail to update - storage error", func(t *testing.T) {
		// arrange
		// - validator: mock
		vlMk := validator.NewValidatorMock()
		vlMk.On("Default", mock.Anything).Return(nil)
		vlMk.On("Validate", mock.Anything).Return(nil)
		vlMk.On("Prepare", mock.Anything).Return(nil)

		// - storage: mock
		stMk := storage.NewStorageWriteMock()
		stMk.On("Update", mock.Anything).Return(storage.ErrStorageNotFound)

		// - storage: validation
		st := storage.NewStorageWriteValidation(stMk, vlMk)

		// act
		err := st.Update(nil)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageNotFound)
		vlMk.AssertExpectations(t)
		stMk.AssertExpectations(t)
	})
}

// Tests for StorageWriteValidation.Delete
func TestStorageWriteValidation_Delete(t *testing.T) {
	t.Run("success to delete", func(t *testing.T) {
		// arrange
		// - storage: mock
		stMk := storage.NewStorageWriteMock()
		stMk.On("Delete", mock.Anything).Return(nil)

		// - storage: validation
		st := storage.NewStorageWriteValidation(stMk, nil)

		// act
		err := st.Delete(0)

		// assert
		require.NoError(t, err)
		stMk.AssertExpectations(t)
	})

	t.Run("fail to delete - storage error", func(t *testing.T) {
		// arrange
		// - storage: mock
		stMk := storage.NewStorageWriteMock()
		stMk.On("Delete", mock.Anything).Return(storage.ErrStorageNotFound)

		// - storage: validation
		st := storage.NewStorageWriteValidation(stMk, nil)

		// act
		err := st.Delete(0)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageNotFound)
		stMk.AssertExpectations(t)
	})
}

// Tests for StorageWriteValidation.Activate
func TestStorageWriteValidation_Activate(t *testing.T) {
	t.Run("success to activate", func(t *testing.T) {
		// arrange
		// - storage: mock
		stMk := storage.NewStorageWriteMock()
		stMk.On("Activate", mock.Anything).Return(nil)

		// - storage: validation
		st := storage.NewStorageWriteValidation(stMk, nil)

		// act
		err := st.Activate(0)

		// assert
		require.NoError(t, err)
		stMk.AssertExpectations(t)
	})

	t.Run("fail to activate - storage error", func(t *testing.T) {
		// arrange
		// - storage: mock
		stMk := storage.NewStorageWriteMock()
		stMk.On("Activate", mock.Anything).Return(storage.ErrStorageNotFound)

		// - storage: validation
		st := storage.NewStorageWriteValidation(stMk, nil)

		// act
		err := st.Activate(0)

		// assert
		require.ErrorIs(t, err, storage.ErrStorageNotFound)
		stMk.AssertExpectations(t)
	})
}