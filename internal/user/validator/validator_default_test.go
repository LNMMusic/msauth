package validator_test

import (
	"testing"

	"github.com/LNMMusic/msauth/internal/user"
	"github.com/LNMMusic/msauth/internal/user/validator"
	"github.com/LNMMusic/msauth/pkg/crypter"

	"github.com/LNMMusic/optional"
	"github.com/stretchr/testify/require"
)

// Test for ValidatorDefault.Default
func TestValidator_Default(t *testing.T) {
	t.Run("default case - all fields none", func(t *testing.T) {
		// arrange
		vl := validator.NewValidatorDefault("", nil)
		u := &user.User{}

		// act
		err := vl.Default(u)

		// assert
		require.NoError(t, err)
		require.True(t, u.IsActive.IsSome())
		isActive, err := u.IsActive.Unwrap()
		require.NoError(t, err)
		require.False(t, isActive)
	})

	t.Run("default case - all fields some", func(t *testing.T) {
		// arrange
		vl := validator.NewValidatorDefault("", nil)
		u := &user.User{
			Id: 1,
			Username: optional.Some("username"),
			Password: optional.Some("password"),
			Email: optional.Some("johndoe@gmail.com"),
			IsActive: optional.Some(true),
		}

		// act
		err := vl.Default(u)

		// assert
		require.NoError(t, err)
		require.True(t, u.IsActive.IsSome())
		isActive, err := u.IsActive.Unwrap()
		require.NoError(t, err)
		require.True(t, isActive)
	})
}

// Test for ValidatorDefault.Validate
func TestValidator_Validate(t *testing.T) {
	type input struct {user *user.User}
	type output struct {err error; errMsg string}
	type testCase struct {
		title  string
		input  input
		output output
	}

	cases := []testCase{
		// valid cases
		{
			title: "valid case - some fields none",
			input: input{user: &user.User{
				Id: 1,
				Username: optional.Some("username"),
				Password: optional.Some("password"),
				Email: optional.Some("john_doe@gmail.com"),
			}},
			output: output{err: nil},
		},
		{
			title: "valid case - all fields some",
			input: input{user: &user.User{
				Id: 1,
				Username: optional.Some("username"),
				Password: optional.Some("password"),
				Email: optional.Some("john_doe@gmail.com"),
				IsActive: optional.Some(true),
			}},
			output: output{err: nil},
		},

		// invalid cases
		// - required fields
		{
			title: "invalid case - username none",
			input: input{user: &user.User{
				Password: optional.Some("password"),
			}},
			output: output{
				err: validator.ErrValidatorFieldRequired,
				errMsg: "validator: field required - username is none",
			},
		},
		{
			title: "invalid case - password none",
			input: input{user: &user.User{
				Username: optional.Some("username"),
			}},
			output: output{
				err: validator.ErrValidatorFieldRequired,
				errMsg: "validator: field required - password is none",
			},
		},
		{
			title: "invalid case - email none",
			input: input{user: &user.User{
				Username: optional.Some("username"),
				Password: optional.Some("password"),
			}},
			output: output{
				err: validator.ErrValidatorFieldRequired,
				errMsg: "validator: field required - email is none",
			},
		},
		// - quality of fields
		{
			title: "invalid case - username too short",
			input: input{user: &user.User{
				Username: optional.Some("us"),
				Password: optional.Some("password"),
				Email: optional.Some("johndoe@gmail.com"),
			}},
			output: output{
				err: validator.ErrValidatorFieldQuality,
				errMsg: "validator: field quality - username, chars should be between 3 and 25",
			},
		},
		{
			title: "invalid case - username too long",
			input: input{user: &user.User{
				Username: optional.Some("usernameusernameusernameusername"),
				Password: optional.Some("password"),
				Email: optional.Some("johndoe@gmail.com"),
			}},
			output: output{
				err: validator.ErrValidatorFieldQuality,
				errMsg: "validator: field quality - username, chars should be between 3 and 25",
			},
		},
		{
			title: "invalid case - password too short",
			input: input{user: &user.User{
				Username: optional.Some("username"),
				Password: optional.Some("pass"),
				Email: optional.Some("johndoe@gmail.com"),
			}},
			output: output{
				err: validator.ErrValidatorFieldQuality,
				errMsg: "validator: field quality - password, chars should be between 8 and 25",
			},
		},
		{
			title: "invalid case - password too long",
			input: input{user: &user.User{
				Username: optional.Some("username"),
				Password: optional.Some("passwordpasswordpasswordpassword"),
				Email: optional.Some("johndoe@gmail.com"),
			}},
			output: output{
				err: validator.ErrValidatorFieldQuality,
				errMsg: "validator: field quality - password, chars should be between 8 and 25",
			},
		},
		{
			title: "invalid case - email invalid format",
			input: input{user: &user.User{
				Username: optional.Some("username"),
				Password: optional.Some("password"),
				Email: optional.Some("johndoegmail.com"),
			}},
			output: output{
				err: validator.ErrValidatorFieldQuality,
				errMsg: "validator: field quality - email, invalid format",
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// arrange
			vl := validator.NewValidatorDefault("", nil)

			// act
			err := vl.Validate(c.input.user)

			// assert
			require.ErrorIs(t, err, c.output.err)
			if err != nil {
				require.Equal(t, c.output.errMsg, err.Error())
			}
		})
	}
}

// Test for ValidatorDefault.Prepare
func TestValidator_Prepare(t *testing.T) {
	t.Run("success to encrypt password", func(t *testing.T) {
		// arrange
		// - crypter: mock
		cr := crypter.NewCrypterMock()
		cr.On("Encrypt", "password").Return("encrypted_password", nil)

		// - validator: default
		vl := validator.NewValidatorDefault("", cr)

		// act
		u := &user.User{
			Password: optional.Some("password"),
		}
		err := vl.Prepare(u)

		// assert
		require.NoError(t, err)
		require.Equal(t, optional.Some[string]("encrypted_password"), u.Password)
		cr.AssertExpectations(t)
	})

	t.Run("fail to encrypt password", func(t *testing.T) {
		// arrange
		// - crypter: mock
		cr := crypter.NewCrypterMock()
		cr.On("Encrypt", "password").Return("", crypter.ErrCrypterEncryption)

		// - validator: default
		vl := validator.NewValidatorDefault("", cr)

		// act
		u := &user.User{
			Password: optional.Some("password"),
		}
		err := vl.Prepare(u)

		// assert
		require.ErrorIs(t, err, validator.ErrValidatorEncryption)
		require.EqualError(t, err, "validator: cannot encrypt field. crypter: encryption error")
		require.Equal(t, optional.Some("password"), u.Password)
		cr.AssertExpectations(t)
	})
}