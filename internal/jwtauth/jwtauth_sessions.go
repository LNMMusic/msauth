package jwtauth

import (
	"errors"
	"fmt"

	"github.com/LNMMusic/msauth/internal/session"
	"github.com/LNMMusic/msauth/internal/session/sessionauth"
)

// constructor
func NewJWTAuthSessions(jw JWTAuth, ss sessionauth.SessionAuthManager) *JWTAuthSessions {
	return &JWTAuthSessions{
		jw: jw,
		ss: ss,
	}
}

// ______________________________________________________________________________________________________________________________
// JWTAuthSessions is an implementation of JWTAuth interface for sessions
// - It use decorator pattern to add session functionality to JWTAuth interface
type JWTAuthSessions struct {
	// JWTAuth interface to generate and validate signs
	jw JWTAuth
	// SessionAuthManager interface to handle sessions
	ss sessionauth.SessionAuthManager
}

// GenerateSign generates a new sign from token info (encryption) and tracks the session by user id
func (j *JWTAuthSessions) GenerateSign(token *Token) (sign string, err error) {
	// generate sign
	sign, err = j.jw.GenerateSign(token)
	if err != nil {
		return
	}

	// save session
	userID, ok := token.Claims["user_id"].(string)
	if !ok {
		sign = ""
		err = fmt.Errorf("%w. %s", ErrJWTAuthInternal, "user id missing")
		return
	}
	session := &session.Session{
		TokenID: token.ID,
		ExpireDate: token.ExpireDate,
	}

	err = j.ss.GenerateSession(userID, session)
	if err != nil {
		sign = ""
		switch {
			case errors.Is(err, sessionauth.ErrSessionReachedMaxSessionsPerUser):
				err = fmt.Errorf("%w. %s", ErrJWTAuthMaxSessions, err.Error())
			default:
				err = fmt.Errorf("%w. %s", ErrJWTAuthInternal, err.Error())
		}
		return
	}
	
	return
}

// ValidateSign validates a sign and returns token info (decryption) and validates the session
func (j *JWTAuthSessions) ValidateSign(sign string) (token *Token, err error) {
	// validate sign
	token, err = j.jw.ValidateSign(sign)
	if err != nil {
		return
	}

	// validate session
	userID, ok := token.Claims["user_id"].(string)
	if !ok {
		token = nil
		err = fmt.Errorf("%w. %s", ErrJWTAuthUnauthorized, "user id missing")
		return
	}
	session := &session.Session{
		TokenID: token.ID,
		ExpireDate: token.ExpireDate,
	}

	err = j.ss.ValidateSession(userID, session.TokenID)
	if err != nil {
		token = nil
		switch {
			case errors.Is(err, sessionauth.ErrSessionAuthManagerUnauthorized):
				err = fmt.Errorf("%w. %s", ErrJWTAuthUnauthorized, err.Error())
			default:
				err = fmt.Errorf("%w. %s", ErrJWTAuthInternal, err.Error())
		}
		return
	}

	return
}