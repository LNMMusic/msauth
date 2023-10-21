package jwtauth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// constructor
func NewJWTAuthBasic(config *Config) *JWTAuthBasic {
	return &JWTAuthBasic{
		config: config,
	}
}

// ______________________________________________________________________________________________________________________________
type CustomClaims struct {
	jwt.RegisteredClaims
	Claims map[string]interface{}
}

// JWTAuthBasic is the default implementation of JWTAuth interface
type Config struct {
	// SigningMethod is the signing method to encrypt and decrypt tokens
	SigningMethod	jwt.SigningMethod
	// Secret is the secret key to encrypt and decrypt tokens
	Secret			[]byte
}

type JWTAuthBasic struct {
	config *Config
}

// GenerateSign generates a new sign from token info (encryption)
// > process
// - claims: token-info (id, expire date, claims) deserialization
// - token: contains claims
// - | encryption
// - sign: token is signed ([]byte or string) with signing method and secret-key
func (j *JWTAuthBasic) GenerateSign(token *Token) (sign string, err error) {
	// -> claims: id, expire date, claims (token-info deserialization)
	claims := &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: token.ID,
			ExpiresAt: jwt.NewNumericDate(token.ExpireDate),
		},
		Claims: token.Claims,
	}
	// -> token: claims
	jwttoken := jwt.NewWithClaims(j.config.SigningMethod, claims)

	// -> sign token (using signing method and secret-key) (encryption)
	sign, err = jwttoken.SignedString(j.config.Secret)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrJWTAuthInternal, err)
	}

	return
}

// ValidateToken validates a sign and returns token info (decryption)
// process
// - sign
// - token/claims: token is parsed with claims from sign (use secret-key and sign method)
// 				   token-info (id, expire date, claims) serialization
func (j *JWTAuthBasic) ValidateSign(sign string) (token *Token, err error) {
	// parse token with claims from sign (using secret-key and sign method) (decryption)
	var jwttoken *jwt.Token
	jwttoken, err = jwt.ParseWithClaims(sign, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {return j.config.Secret, nil})
	if err != nil {
		switch {
		// -> sign malformed
		case errors.Is(err, jwt.ErrTokenMalformed):
			err = fmt.Errorf("%w. %v", ErrJWTAuthUnauthorized, err)
		// -> token expired
		case errors.Is(err, jwt.ErrTokenExpired):
			err = fmt.Errorf("%w. %v", ErrJWTExpired, err)
		default:
			err = fmt.Errorf("%w. %v", ErrJWTAuthInternal, err)
		}
		return
	}

	// -> token-info (id, expire date, claims) serialization
	claims, ok := jwttoken.Claims.(*CustomClaims)
	if !ok {
		err = fmt.Errorf("%w. %v", ErrJWTAuthInternal, err)
		return
	}

	token = &Token{
		ID: claims.ID,
		ExpireDate: claims.ExpiresAt.Time,
		Claims: claims.Claims,
	}
	return
}
