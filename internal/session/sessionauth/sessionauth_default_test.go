package sessionauth

import (
	"testing"
	"time"

	"github.com/LNMMusic/msauth/internal/session"
	"github.com/LNMMusic/msauth/internal/session/storage"
	"github.com/stretchr/testify/assert"
)

// Tests for SessionAuthManagerDefault
func TestSessionAuthManagerDefault_GenerateSession(t *testing.T) {
	type input struct {userId string; s *session.Session}
	type output struct {err error; errMsg string}
	type testCase struct {
		// base
		title  string
		input  input
		output output
		// set-up
		setUpStorage func(mk *storage.StorageMock)
		setUpConfig  func(cfg *Config)
	}

	cases := []testCase{
		// valid cases
		{
			title: "success",
			input: input{userId: "user-id", s: &session.Session{
				TokenID: "token-id",
				ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			}},
			output: output{err: nil, errMsg: ""},
			setUpStorage: func(mk *storage.StorageMock) {
				mk.On("Get", "user-id").Return([]*session.Session{}, nil)
				mk.On("Set", "user-id", []*session.Session{
					{
						TokenID: "token-id",
						ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				}).Return(nil)
			},
			setUpConfig: func(cfg *Config) {},
		},
		{
			title: "success - 5 sessions, all expired (should be filtered by sync)",
			input: input{userId: "user-id", s: &session.Session{
				TokenID: "token-id",
				ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			}},
			output: output{err: nil, errMsg: ""},
			setUpStorage: func(mk *storage.StorageMock) {
				mk.On("Get", "user-id").Return([]*session.Session{
					{ExpireDate: time.Now().Add(-1 * time.Hour)},
					{ExpireDate: time.Now().Add(-1 * time.Hour)},
					{ExpireDate: time.Now().Add(-1 * time.Hour)},
					{ExpireDate: time.Now().Add(-1 * time.Hour)},
					{ExpireDate: time.Now().Add(-1 * time.Hour)},
				}, nil)
				mk.On("Set", "user-id", []*session.Session{
					{
						TokenID: "token-id",
						ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				}).Return(nil)
			},
			setUpConfig: func(cfg *Config) {},
		},

		// invalid cases
		// -> storage
		{
			title: "storage error - get",
			input: input{userId: "user-id", s: &session.Session{}},
			output: output{err: ErrSessionAuthManagerInternal, errMsg: "internal session auth manager error. internal storage error"},
			setUpStorage: func(mk *storage.StorageMock) {
				mk.On("Get", "user-id").Return([]*session.Session{}, storage.ErrStorageInternal)
			},
			setUpConfig: func(cfg *Config) {},
		},
		{
			title: "storage error - set",
			input: input{userId: "user-id", s: &session.Session{
				TokenID: "token-id",
				ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			}},
			output: output{err: ErrSessionAuthManagerInternal, errMsg: "internal session auth manager error. internal storage error"},
			setUpStorage: func(mk *storage.StorageMock) {
				mk.On("Get", "user-id").Return([]*session.Session{}, nil)
				mk.On("Set", "user-id", []*session.Session{
					{
						TokenID: "token-id",
						ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				}).Return(storage.ErrStorageInternal)
			},
			setUpConfig: func(cfg *Config) {},
		},
		// -> config
		{
			title: "config error - max sessions per user - 5 sessions non expired",
			input: input{userId: "user-id", s: &session.Session{}},
			output: output{err: ErrSessionReachedMaxSessionsPerUser, errMsg: "max sessions per user reached. 5"},
			setUpStorage: func(mk *storage.StorageMock) {
				mk.On("Get", "user-id").Return([]*session.Session{
					{ExpireDate: time.Now().Add(1 * time.Hour)},
					{ExpireDate: time.Now().Add(1 * time.Hour)},
					{ExpireDate: time.Now().Add(1 * time.Hour)},
					{ExpireDate: time.Now().Add(1 * time.Hour)},
					{ExpireDate: time.Now().Add(1 * time.Hour)},
				}, nil)
			},
			setUpConfig: func(cfg *Config) {},
		},

	}

	// run tests
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// arrange
			st := storage.NewStorageMock()
			c.setUpStorage(st)

			cfg := &Config{}
			c.setUpConfig(cfg)
			ss := NewSessionAuthManagerDefault(st, cfg)

			// act
			err := ss.GenerateSession(c.input.userId, c.input.s)

			// assert
			assert.ErrorIs(t, err, c.output.err)
			if c.output.err != nil {
				assert.Equal(t, c.output.errMsg, err.Error())
			}
			st.AssertExpectations(t)
		})
	}
}

func TestSessionAuthManagerDefault_ValidateSession(t *testing.T) {
	type input struct {userId string; tokenId string}
	type output struct {err error; errMsg string}
	type testCase struct {
		// base
		title  string
		input  input
		output output
		// set-up
		setUpStorage func(mk *storage.StorageMock)
		setUpConfig  func(cfg *Config)
	}

	cases := []testCase{
		// valid cases
		{
			title: "success",
			input: input{userId: "user-id", tokenId: "token-id"},
			output: output{err: nil, errMsg: ""},
			setUpStorage: func(mk *storage.StorageMock) {
				mk.On("Get", "user-id").Return([]*session.Session{
					{
						TokenID: "token-id",
						ExpireDate: time.Now().Add(1 * time.Hour),
					},
				}, nil)
			},
			setUpConfig: func(cfg *Config) {},
		},

		// invalid cases
		// -> storage
		{
			title: "storage error - get",
			input: input{userId: "user-id", tokenId: "token-id"},
			output: output{err: ErrSessionAuthManagerInternal, errMsg: "internal session auth manager error. internal storage error"},
			setUpStorage: func(mk *storage.StorageMock) {
				mk.On("Get", "user-id").Return([]*session.Session{}, storage.ErrStorageInternal)
			},
			setUpConfig: func(cfg *Config) {},
		},
		// -> validation
		{
			title: "validation error - session expired",
			input: input{userId: "user-id", tokenId: "token-id"},
			output: output{err: ErrSessionAuthManagerUnauthorized, errMsg: "unauthorized session. token-id"},
			setUpStorage: func(mk *storage.StorageMock) {
				mk.On("Get", "user-id").Return([]*session.Session{
					{
						TokenID: "token-id",
						ExpireDate: time.Now().Add(-1 * time.Hour),
					},
				}, nil)
			},
			setUpConfig: func(cfg *Config) {},
		},
		{
			title: "validation error - session not found",
			input: input{userId: "user-id", tokenId: "token-id"},
			output: output{err: ErrSessionAuthManagerUnauthorized, errMsg: "unauthorized session. token-id"},
			setUpStorage: func(mk *storage.StorageMock) {
				mk.On("Get", "user-id").Return([]*session.Session{}, nil)
			},
			setUpConfig: func(cfg *Config) {},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// arrange
			st := storage.NewStorageMock()
			c.setUpStorage(st)

			cfg := &Config{}
			c.setUpConfig(cfg)
			ss := NewSessionAuthManagerDefault(st, cfg)

			// act
			err := ss.ValidateSession(c.input.userId, c.input.tokenId)

			// assert
			assert.ErrorIs(t, err, c.output.err)
			if c.output.err != nil {
				assert.Equal(t, c.output.errMsg, err.Error())
			}
			st.AssertExpectations(t)
		})
	}
}