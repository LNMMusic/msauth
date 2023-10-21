package storage_test

import (
	"sync"
	"testing"

	"github.com/LNMMusic/msauth/internal/user"
	"github.com/LNMMusic/msauth/internal/user/storage"
	"github.com/LNMMusic/optional"
	"github.com/stretchr/testify/require"
)

// Tests for StorageReadMap.Get
func TestStorageReadMap_Get(t *testing.T) {
	// test case
	type arrange struct{
		setUpStorage func() *storage.StorageReadMap
	}
	type input struct{id int}
	type output struct{u user.User; err error; errMsg string}
	type test struct {
		// test case name
		name    string
		// test structure
		arrange arrange
		input   input
		output  output
	}

	cases := []test{
		// success
		{
			name: "success to get user by id",
			arrange: arrange{
				setUpStorage: func() *storage.StorageReadMap {
					db := &sync.Map{}
					db.Store(1, user.User{
						Id: 1,
						Username: optional.Some[string]("johndoe"),
						Password: optional.Some[string]("hashedPassword"),
						Email: optional.Some[string]("johndoe@gmail.com"),
						IsActive: optional.Some[bool](false),
					})

					return storage.NewStorageReadMap(db)
				},
			},
			input: input{id: 1},
			output: output{
				u: user.User{
					Id: 1,
					Username: optional.Some[string]("johndoe"),
					Password: optional.Some[string]("hashedPassword"),
					Email: optional.Some[string]("johndoe@gmail.com"),
					IsActive: optional.Some[bool](false),
				},
				err: nil,
				errMsg: "",
			},
		},

		// failure
		{
			name: "failure to get user by id - user not found",
			arrange: arrange{
				setUpStorage: func() *storage.StorageReadMap {
					return storage.NewStorageReadMap(nil)
				},
			},
			input: input{id: 1},
			output: output{
				u: user.User{},
				err: storage.ErrStorageNotFound,
				errMsg: "storage: user not found - 1",
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			st := c.arrange.setUpStorage()

			// act
			u, err := st.Get(c.input.id)

			// assert
			require.Equal(t, c.output.u, u)
			require.ErrorIs(t, err, c.output.err)
			if c.output.err != nil {
				require.EqualError(t, err, c.output.errMsg)
			}
		})
	}
}

// Tests for StorageReadMap.GetByEmail
func TestStorageReadMap_GetByEmail(t *testing.T) {
	// test case
	type arrange struct{
		setUpStorage func() *storage.StorageReadMap
	}
	type input struct{email string}
	type output struct{u user.User; err error; errMsg string}
	type test struct {
		// test case name
		name    string
		// test structure
		arrange arrange
		input   input
		output  output
	}

	cases := []test{
		// success
		{
			name: "success to get user by email",
			arrange: arrange{
				setUpStorage: func() *storage.StorageReadMap {
					db := &sync.Map{}
					db.Store(1, user.User{
						Id: 1,
						Username: optional.Some[string]("johndoe"),
						Password: optional.Some[string]("hashedPassword"),
						Email: optional.Some[string]("johndoe@gmail.com"),
						IsActive: optional.Some[bool](false),
					})
					
					return storage.NewStorageReadMap(db)
				},
			},
			input: input{email: "johndoe@gmail.com"},
			output: output{
				u: user.User{
					Id: 1,
					Username: optional.Some[string]("johndoe"),
					Password: optional.Some[string]("hashedPassword"),
					Email: optional.Some[string]("johndoe@gmail.com"),
					IsActive: optional.Some[bool](false),
				},
				err: nil,
				errMsg: "",
			},
		},

		// failure
		{
			name: "failure to get user by email - user not found",
			arrange: arrange{
				setUpStorage: func() *storage.StorageReadMap {
					return storage.NewStorageReadMap(nil)
				},
			},
			input: input{email: "johndoe@gmail.com"},
			output: output{
				u: user.User{},
				err: storage.ErrStorageNotFound,
				errMsg: "storage: user not found - johndoe@gmail.com",
			},
		},
		{
			name: "failure to get user by email - user not found (db not empty)",
			arrange: arrange{
				setUpStorage: func() *storage.StorageReadMap {
					db := &sync.Map{}
					db.Store(1, user.User{
						Id: 1,
						Username: optional.Some[string]("johndoe"),
						Password: optional.Some[string]("hashedPassword"),
						Email: optional.Some[string](""), // empty email
						IsActive: optional.Some[bool](false),
					})
					
					return storage.NewStorageReadMap(db)
				},
			},
			input: input{email: "johndoe@gmail.com"},
			output: output{
				u: user.User{},
				err: storage.ErrStorageNotFound,
				errMsg: "storage: user not found - johndoe@gmail.com",
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			st := c.arrange.setUpStorage()

			// act
			u, err := st.GetByEmail(c.input.email)

			// assert
			require.Equal(t, c.output.u, u)
			require.ErrorIs(t, err, c.output.err)
			if c.output.err != nil {
				require.EqualError(t, err, c.output.errMsg)
			}
		})
	}
}

// Tests for StorageReadMap.GetByUsername
func TestStorageReadMap_GetByUsername(t *testing.T) {
	// test case
	type arrange struct{
		setUpStorage func() *storage.StorageReadMap
	}
	type input struct{username string}
	type output struct{u user.User; err error; errMsg string}
	type test struct {
		// test case name
		name    string
		// test structure
		arrange arrange
		input   input
		output  output
	}

	cases := []test{
		// success
		{
			name: "success to get user by username",
			arrange: arrange{
				setUpStorage: func() *storage.StorageReadMap {
					db := &sync.Map{}
					db.Store(1, user.User{
						Id: 1,
						Username: optional.Some[string]("johndoe"),
						Password: optional.Some[string]("hashedPassword"),
						Email: optional.Some[string]("johndoe@gmail.com"),
						IsActive: optional.Some[bool](false),
					})
					
					return storage.NewStorageReadMap(db)
				},
			},
			input: input{username: "johndoe"},
			output: output{
				u: user.User{
					Id: 1,
					Username: optional.Some[string]("johndoe"),
					Password: optional.Some[string]("hashedPassword"),
					Email: optional.Some[string]("johndoe@gmail.com"),
					IsActive: optional.Some[bool](false),
				},
				err: nil,
				errMsg: "",
			},
		},

		// failure
		{
			name: "failure to get user by username - user not found",
			arrange: arrange{
				setUpStorage: func() *storage.StorageReadMap {
					return storage.NewStorageReadMap(nil)
				},
			},
			input: input{username: "johndoe"},
			output: output{
				u: user.User{},
				err: storage.ErrStorageNotFound,
				errMsg: "storage: user not found - johndoe",
			},
		},
		{
			name: "failure to get user by username - user not found (db not empty)",
			arrange: arrange{
				setUpStorage: func() *storage.StorageReadMap {
					db := &sync.Map{}
					db.Store(1, user.User{
						Id: 1,
						Username: optional.Some[string](""),
						Password: optional.Some[string]("hashedPassword"),
						Email: optional.Some[string]("johndoe@gmail.com"),
						IsActive: optional.Some[bool](false),
					})

					return storage.NewStorageReadMap(db)
				},
			},
			input: input{username: "johndoe"},
			output: output{
				u: user.User{},
				err: storage.ErrStorageNotFound,
				errMsg: "storage: user not found - johndoe",
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			st := c.arrange.setUpStorage()

			// act
			u, err := st.GetByUsername(c.input.username)

			// assert
			require.Equal(t, c.output.u, u)
			require.ErrorIs(t, err, c.output.err)
			if c.output.err != nil {
				require.EqualError(t, err, c.output.errMsg)
			}
		})
	}
}