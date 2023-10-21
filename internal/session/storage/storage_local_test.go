package storage

import (
	"testing"
	"time"

	"github.com/LNMMusic/msauth/internal/session"
	"github.com/stretchr/testify/assert"
)

// Tests for StorageLocal
func TestStorageLocal_Get(t *testing.T) {
	type input struct {userId string}
	type output struct {sessions []*session.Session; err error; errMsg string}
	type testCase struct {
		// base
		title  string
		input  input
		output output
		// set-up
		setUpDb func (db *map[string][]*session.Session)
	}

	cases := []testCase{
		// valid cases
		{
			title: "user has 0 sessions",
			input: input{userId: "user1"},
			output: output{sessions: []*session.Session{}, err: nil, errMsg: ""},
			setUpDb: func (db *map[string][]*session.Session) {
				(*db)["user1"] = []*session.Session{}
			},
		},
		{
			title: "user has 1 session",
			input: input{userId: "user1"},
			output: output{
				sessions: []*session.Session{
					{TokenID: "token1", ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				},
				err: nil,
				errMsg: "",
			},
			setUpDb: func (db *map[string][]*session.Session) {
				(*db)["user1"] = []*session.Session{
					{TokenID: "token1", ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				}
			},
		},
		{
			title: "user has 5 sessions",
			input: input{userId: "user1"},
			output: output{
				sessions: []*session.Session{
					{TokenID: "token1", ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					{TokenID: "token2", ExpireDate: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
					{TokenID: "token3", ExpireDate: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
					{TokenID: "token4", ExpireDate: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
					{TokenID: "token5", ExpireDate: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				},
				err: nil,
				errMsg: "",
			},
			setUpDb: func (db *map[string][]*session.Session) {
				(*db)["user1"] = []*session.Session{
					{TokenID: "token1", ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					{TokenID: "token2", ExpireDate: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
					{TokenID: "token3", ExpireDate: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
					{TokenID: "token4", ExpireDate: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
					{TokenID: "token5", ExpireDate: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				}
			},
		},
		
		// invalid cases
		{
			title: "user not found",
			input: input{userId: "user1"},
			output: output{sessions: nil, err: ErrStorageUserNotFound, errMsg: "user id not found: user1"},
			setUpDb: func (db *map[string][]*session.Session) {
				(*db)["user2"] = []*session.Session{}
			},
		},
	}
	
	// run tests
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// arrange
			db := make(map[string][]*session.Session)
			c.setUpDb(&db)
			
			st := NewStorageLocal(db)

			// act
			sessions, err := st.Get(c.input.userId)

			// assert
			assert.Equal(t, c.output.sessions, sessions)
			assert.ErrorIs(t, err, c.output.err)
			if c.output.err != nil {
				assert.EqualError(t, err, c.output.errMsg)
			}
		})
	}
}

func TestStorageLocal_Set(t *testing.T) {
	type input struct {userId string; sessions []*session.Session}
	type output struct {db map[string][]*session.Session; err error; errMsg string}
	type testCase struct {
		// base
		title  string
		input  input
		output output
		// set-up
		setUpDb func (db *map[string][]*session.Session)
	}

	cases := []testCase{
		// valid cases
		{
			title: "user not exists",
			input: input{
				userId: "user1",
				sessions: []*session.Session{
					{TokenID: "token1", ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				},
			},
			output: output{
				db: map[string][]*session.Session{
					"user1": {
						{TokenID: "token1", ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
					},
					"user2": {},
				},
				err: nil,
				errMsg: "",
			},
			setUpDb: func (db *map[string][]*session.Session) {
				(*db)["user2"] = []*session.Session{}
			},
		},
		{
			title: "user exists",
			input: input{
				userId: "user1",
				sessions: []*session.Session{
					{TokenID: "token2", ExpireDate: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				},
			},
			output: output{
				db: map[string][]*session.Session{
					"user1": {
						{TokenID: "token2", ExpireDate: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
					},
					"user2": {},
				},
				err: nil,
				errMsg: "",
			},
			setUpDb: func (db *map[string][]*session.Session) {
				(*db)["user1"] = []*session.Session{
					{TokenID: "token1", ExpireDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				}
				(*db)["user2"] = []*session.Session{}
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// arrange
			db := make(map[string][]*session.Session)
			c.setUpDb(&db)

			st := NewStorageLocal(db)

			// act
			err := st.Set(c.input.userId, c.input.sessions)

			// assert
			assert.ErrorIs(t, err, c.output.err)
			if c.output.err != nil {
				assert.EqualError(t, err, c.output.errMsg)
			}
			assert.Equal(t, c.output.db, db)
		})
	}
}