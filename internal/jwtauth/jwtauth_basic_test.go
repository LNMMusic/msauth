package jwtauth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// Tests for ImplJWTAuthDefault
func TestImplJWTAuthDefault_GenerateSign(t *testing.T) {
	type input struct {token *Token}
	type output struct {err error; errMsg string}
	type testCase struct {
		// base
		title  string
		input  input
		output output
		// set-up
		setUpConfig func(cfg *Config)
	}

	cases := []testCase{
		// valid cases
		{
			title: "valid case",
			input: input{
				token: &Token{
					ID: "id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"key": "value",
					},
				},
			},
			output: output{err: nil, errMsg: ""},
			setUpConfig: func(cfg *Config) {
				(*cfg).SigningMethod = jwt.SigningMethodHS256
				(*cfg).Secret = []byte("secret")
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// arrange
			cfg := &Config{}
			c.setUpConfig(cfg)

			j := NewJWTAuthBasic(cfg)

			// act
			_, err := j.GenerateSign(c.input.token)

			// assert
			assert.ErrorIs(t, err, c.output.err)
			if err != nil {
				assert.Equal(t, c.output.errMsg, err.Error())
			}
		})
	}
}

func TestImplJWTAuthDefault_ValidateSign(t *testing.T) {
	// valid cases
	t.Run("valid case", func(t *testing.T) {
		// arrange
		cfg := &Config{}
		(*cfg).SigningMethod = jwt.SigningMethodHS256
		(*cfg).Secret = []byte("secret")

		j := NewJWTAuthBasic(cfg)

		token := &Token{
			ID: "id",
			ExpireDate: time.Now().Add(time.Hour),
			Claims: map[string]interface{}{"key": "value"},
		}
		sign, err := j.GenerateSign(token)
		assert.NoError(t, err)

		// act
		token, err = j.ValidateSign(sign)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, "id", token.ID)
		assert.Equal(t, "value", token.Claims["key"])
	})

	// invalid cases
	t.Run("invalid case - unauthorized sign - malformed", func(t *testing.T) {
		// arrange
		cfg := &Config{}
		(*cfg).SigningMethod = jwt.SigningMethodHS256
		(*cfg).Secret = []byte("secret")

		j := NewJWTAuthBasic(cfg)

		// act
		token, err := j.ValidateSign("invalid")
		
		// assert
		assert.ErrorIs(t, err, ErrJWTAuthUnauthorized)
		assert.Nil(t, token)
	})

	t.Run("invalid case - expired token", func(t *testing.T) {
		// arrange
		cfg := &Config{}
		(*cfg).SigningMethod = jwt.SigningMethodHS256
		(*cfg).Secret = []byte("secret")

		j := NewJWTAuthBasic(cfg)

		token := &Token{
			ID: "id",
			ExpireDate: time.Now().Add(-time.Hour),
			Claims: map[string]interface{}{"key": "value"},
		}
		sign, err := j.GenerateSign(token)
		assert.NoError(t, err)

		// act
		token, err = j.ValidateSign(sign)

		// assert
		assert.ErrorIs(t, err, ErrJWTExpired)
		assert.Nil(t, token)
	})
}