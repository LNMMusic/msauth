package storage_test

import (
	"sync"
	"testing"

	"github.com/LNMMusic/msauth/internal/user"
	"github.com/LNMMusic/msauth/internal/user/storage"
	"github.com/LNMMusic/optional"
	"github.com/stretchr/testify/require"
)

// Tests for StorageWriteMap.Create
func TestStorageWriteMap_Create(t *testing.T) {
	type arrange struct {storage func() *storage.StorageWriteMap}
	type input struct {u *user.User}
	type output struct {id int; err error; errMsg string}
	type testCase struct {
		// name of the test case
		name string
		// structure of the test case
		arrange arrange
		input   input
		output  output
	}

	// test cases
	testCases := []testCase{
		// success cases
		{
			name: "success - create a new user",
			arrange: arrange{
				storage: func() *storage.StorageWriteMap {
					return storage.NewStorageWriteMap(nil, 0)
				},
			},
			input: input{
				u: &user.User{
					Id: 0,
					Username: optional.Some[string]("username"),
					Password: optional.Some[string]("password"),
					Email: optional.Some[string]("email"),
					IsActive: optional.Some[bool](false),
				},
			},
			output: output{
				id: 1,
				err: nil,
				errMsg: "",
			},
		},
	}

	// run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			st := tc.arrange.storage()

			// act
			err := st.Create(tc.input.u)

			// assert
			require.Equal(t, tc.output.id, tc.input.u.Id)
			require.ErrorIs(t, err, tc.output.err)
			if tc.output.err != nil {
				require.EqualError(t, err, tc.output.errMsg)
			}
		})
	}
}

// Tests for StorageWriteMap.Update
func TestStorageWriteMap_Update(t *testing.T) {
	type arrange struct {storage func() *storage.StorageWriteMap}
	type input struct {u *user.User}
	type output struct {err error; errMsg string}
	type testCase struct {
		// name of the test case
		name string
		// structure of the test case
		arrange arrange
		input   input
		output  output
	}

	// test cases
	testCases := []testCase{
		// success cases
		{
			name: "success - update an existing user",
			arrange: arrange{
				storage: func() *storage.StorageWriteMap {
					db := &sync.Map{}
					db.Store(1, user.User{
						Id: 1,
						Username: optional.Some[string]("username"),
						Password: optional.Some[string]("password"),
						Email: optional.Some[string]("email"),
						IsActive: optional.Some[bool](false),
					})

					return storage.NewStorageWriteMap(db, 1)
				},
			},
			input: input{
				u: &user.User{
					Id: 1,
					Username: optional.Some[string]("username_updated"),
					Password: optional.Some[string]("password_updated"),
					Email: optional.Some[string]("email_updated"),
					IsActive: optional.Some[bool](true),
				},
			},
			output: output{
				err: nil,
				errMsg: "",
			},
		},

		// failure cases
		{
			name: "failure - update a non-existing user",
			arrange: arrange{
				storage: func() *storage.StorageWriteMap {
					return storage.NewStorageWriteMap(nil, 0)
				},
			},
			input: input{
				u: &user.User{
					Id: 1,
					Username: optional.Some[string]("username_updated"),
					Password: optional.Some[string]("password_updated"),
					Email: optional.Some[string]("email_updated"),
					IsActive: optional.Some[bool](true),
				},
			},
			output: output{
				err: storage.ErrStorageNotFound,
				errMsg: "storage: user not found - 1",
			},
		},
	}

	// run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			st := tc.arrange.storage()

			// act
			err := st.Update(tc.input.u)

			// assert
			require.ErrorIs(t, err, tc.output.err)
			if tc.output.err != nil {
				require.EqualError(t, err, tc.output.errMsg)
			}
		})
	}
}

// Tests for StorageWriteMap.Delete
func TestStorageWriteMap_Delete(t *testing.T) {
	type arrange struct {storage func() *storage.StorageWriteMap}
	type input struct {id int}
	type output struct {err error; errMsg string}
	type testCase struct {
		// name of the test case
		name string
		// structure of the test case
		arrange arrange
		input   input
		output  output
	}

	// test cases
	testCases := []testCase{
		// success cases
		{
			name: "success - delete an existing user",
			arrange: arrange{
				storage: func() *storage.StorageWriteMap {
					db := &sync.Map{}
					db.Store(1, user.User{
						Id: 1,
						Username: optional.Some[string]("username"),
						Password: optional.Some[string]("password"),
						Email: optional.Some[string]("email"),
						IsActive: optional.Some[bool](false),
					})

					return storage.NewStorageWriteMap(db, 1)
				},
			},
			input: input{
				id: 1,
			},
			output: output{
				err: nil,
				errMsg: "",
			},
		},

		// failure cases
		{
			name: "failure - delete a non-existing user",
			arrange: arrange{
				storage: func() *storage.StorageWriteMap {
					return storage.NewStorageWriteMap(nil, 0)
				},
			},
			input: input{
				id: 1,
			},
			output: output{
				err: storage.ErrStorageNotFound,
				errMsg: "storage: user not found - 1",
			},
		},
	}

	// run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			st := tc.arrange.storage()

			// act
			err := st.Delete(tc.input.id)

			// assert
			require.ErrorIs(t, err, tc.output.err)
			if tc.output.err != nil {
				require.EqualError(t, err, tc.output.errMsg)
			}
		})
	}
}

// Tests for StorageWriteMap.Activate
func TestStorageWriteMap_Activate(t *testing.T) {
	type arrange struct {storage func() *storage.StorageWriteMap}
	type input struct {id int}
	type output struct {err error; errMsg string}
	type testCase struct {
		// name of the test case
		name string
		// structure of the test case
		arrange arrange
		input   input
		output  output
	}

	// test cases
	testCases := []testCase{
		// success cases
		{
			name: "success - activate an existing user",
			arrange: arrange{
				storage: func() *storage.StorageWriteMap {
					db := &sync.Map{}
					db.Store(1, user.User{
						Id: 1,
						Username: optional.Some[string]("username"),
						Password: optional.Some[string]("password"),
						Email: optional.Some[string]("email"),
						IsActive: optional.Some[bool](false),
					})

					return storage.NewStorageWriteMap(db, 1)
				},
			},
			input: input{
				id: 1,
			},
			output: output{
				err: nil,
				errMsg: "",
			},
		},

		// failure cases
		{
			name: "failure - activate a non-existing user",
			arrange: arrange{
				storage: func() *storage.StorageWriteMap {
					return storage.NewStorageWriteMap(nil, 0)
				},
			},
			input: input{
				id: 1,
			},
			output: output{
				err: storage.ErrStorageNotFound,
				errMsg: "storage: user not found - 1",
			},
		},
	}

	// run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			st := tc.arrange.storage()

			// act
			err := st.Activate(tc.input.id)

			// assert
			require.ErrorIs(t, err, tc.output.err)
			if tc.output.err != nil {
				require.EqualError(t, err, tc.output.errMsg)
			}
		})
	}
}