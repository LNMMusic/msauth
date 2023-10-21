package jwtauth

import (
	"fmt"
	"testing"
	"time"

	"github.com/LNMMusic/msauth/internal/session"
	"github.com/LNMMusic/msauth/internal/session/sessionauth"
	"github.com/stretchr/testify/assert"
)

// Tests for ImplJWTAuthSessions
func TestImplJWTAuthSessions_GenerateSign(t *testing.T) {
	type input struct { token *Token }
	type output struct { sign string; err error; errMsg string }
	type testCase struct {
		// base
		title  string
		input  input
		output output
		// set-up
		setUpJWTAuth     func(mk *JWTAuthMock)
		setUpSessionAuth func(mk *sessionauth.SessionAuthMock)
	}

	cases := []testCase{
		// valid cases
		{
			title: "valid case",
			input: input{
				token: &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				},
			},
			output: output{
				sign: "sign",
				err: nil,
				errMsg: "",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("GenerateSign", &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				}).Return("sign", nil)
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {
				mk.On("GenerateSession", "#01", &session.Session{
					TokenID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(nil)
			},
		},

		// invalid cases
		// -> jw error
		{
			title: "jw error",
			input: input{
				token: &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				},
			},
			output: output{
				sign: "",
				err: ErrJWTAuthInternal,
				errMsg: "internal jwt auth error. extra message",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("GenerateSign", &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				}).Return("", fmt.Errorf("%w. %s", ErrJWTAuthInternal, "extra message"))
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {},
		},

		// -> ss error
		{
			title: "ss error - invalid user id",
			input: input{
				token: &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{},
				},
			},
			output: output{
				sign: "",
				err: ErrJWTAuthInternal,
				errMsg: "internal jwt auth error. user id missing",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("GenerateSign", &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{},
				}).Return("sign", nil)
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {},
		},
		{
			title: "ss error - max session reached",
			input: input{
				token: &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				},
			},
			output: output{
				sign: "",
				err: ErrJWTAuthMaxSessions,
				errMsg: "max session reached. max session per user reached",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("GenerateSign", &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				}).Return("sign", nil)
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {
				mk.On("GenerateSession", "#01", &session.Session{
					TokenID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(sessionauth.ErrSessionReachedMaxSessionsPerUser)
			},
		},
		{
			title: "ss error - internal error",
			input: input{
				token: &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				},
			},
			output: output{
				sign: "",
				err: ErrJWTAuthInternal,
				errMsg: "internal jwt auth error. internal session auth manager error",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("GenerateSign", &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				}).Return("sign", nil)
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {
				mk.On("GenerateSession", "#01", &session.Session{
					TokenID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				}).Return(sessionauth.ErrSessionAuthManagerInternal)
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// arrange
			jw := NewJWTAuthMock()
			c.setUpJWTAuth(jw)

			ss := sessionauth.NewSessionAuthMock()
			c.setUpSessionAuth(ss)

			j := NewJWTAuthSessions(jw, ss)

			// act
			sign, err := j.GenerateSign(c.input.token)

			// assert
			assert.ErrorIs(t, err, c.output.err)
			if c.output.err != nil {
				assert.Equal(t, c.output.errMsg, err.Error())
			}
			assert.Equal(t, c.output.sign, sign)
			jw.AssertExpectations(t)
			ss.AssertExpectations(t)
		})
	}
}

func TestImplJWTAuthSessions_ValidateSign(t *testing.T) {
	type input struct { sign string }
	type output struct { token *Token; err error; errMsg string }
	type testCase struct {
		// base
		title  string
		input  input
		output output
		// set-up
		setUpJWTAuth     func(mk *JWTAuthMock)
		setUpSessionAuth func(mk *sessionauth.SessionAuthMock)
	}

	cases := []testCase{
		// valid cases
		{
			title: "valid case",
			input: input{sign: "sign"},
			output: output{
				token: &Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				},
				err: nil,
				errMsg: "",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("ValidateSign", "sign").Return(&Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				}, nil)
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {
				mk.On("ValidateSession", "#01", "token_id").Return(nil)
			},
		},

		// invalid cases
		// -> jw error
		{
			title: "jw error",
			input: input{sign: "sign"},
			output: output{
				token: &Token{},
				err: ErrJWTAuthInternal,
				errMsg: "internal jwt auth error. extra message",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("ValidateSign", "sign").Return(&Token{}, fmt.Errorf("%w. %s", ErrJWTAuthInternal, "extra message"))
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {},
		},
		// -> ss error
		{
			title: "ss error - invalid user id",
			input: input{sign: "sign"},
			output: output{
				token: nil,
				err: ErrJWTAuthUnauthorized,
				errMsg: "unauthorized token. user id missing",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("ValidateSign", "sign").Return(&Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{},
				}, nil)
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {},
		},
		{
			title: "ss error - invalid token id",
			input: input{sign: "sign"},
			output: output{
				token: nil,
				err: ErrJWTAuthUnauthorized,
				errMsg: "unauthorized token. unauthorized session",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("ValidateSign", "sign").Return(&Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				}, nil)
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {
				mk.On("ValidateSession", "#01", "token_id").Return(sessionauth.ErrSessionAuthManagerUnauthorized)
			},
		},
		{
			title: "ss error - internal error",
			input: input{sign: "sign"},
			output: output{
				token: nil,
				err: ErrJWTAuthInternal,
				errMsg: "internal jwt auth error. internal session auth manager error",
			},
			setUpJWTAuth: func(mk *JWTAuthMock) {
				mk.On("ValidateSign", "sign").Return(&Token{
					ID: "token_id",
					ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					Claims: map[string]interface{}{
						"user_id": "#01",
					},
				}, nil)
			},
			setUpSessionAuth: func(mk *sessionauth.SessionAuthMock) {
				mk.On("ValidateSession", "#01", "token_id").Return(sessionauth.ErrSessionAuthManagerInternal)
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// arrange
			jw := NewJWTAuthMock()
			c.setUpJWTAuth(jw)

			ss := sessionauth.NewSessionAuthMock()
			c.setUpSessionAuth(ss)

			j := NewJWTAuthSessions(jw, ss)

			// act
			token, err := j.ValidateSign(c.input.sign)

			// assert
			assert.ErrorIs(t, err, c.output.err)
			if c.output.err != nil {
				assert.Equal(t, c.output.errMsg, err.Error())
			}
			assert.Equal(t, c.output.token, token)
			jw.AssertExpectations(t)
			ss.AssertExpectations(t)
		})
	}
}
