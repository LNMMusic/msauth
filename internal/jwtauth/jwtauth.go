package jwtauth

import (
	"errors"
	"time"
)

var (
	// ErrJWTAuthInternal is an error that represents an internal error in the auth process
	ErrJWTAuthInternal		 = errors.New("internal jwt auth error")
	// ErrJWTAuthUnauthorized is an error that represents an unauthorized token
	ErrJWTAuthUnauthorized	 = errors.New("unauthorized token")
	// ErrJWTExpired is an error that represents an expired token
	ErrJWTExpired			 = errors.New("token expired")
	// ErrJWTAuthMaxSessions is an error that represents a max sessions reached
	ErrJWTAuthMaxSessions	 = errors.New("max sessions reached")
)

// Token is a struct that represents a user token
type Token struct {
	ID         string			`json:"id"`
	ExpireDate time.Time		`json:"expire_date"`
	Claims	   map[string]any	`json:"claims"`
}

// JWTAuth is an interface for auth to handle authentication operations for users sessions (stateless)
type JWTAuth interface {
	// GenerateSign generates a new sign from token info (encryption)
	GenerateSign(token *Token) (sign string, err error)
	
	// ValidateToken validates a sign and returns token info (decryption)
	ValidateSign(sign string) (token *Token, err error)
}
