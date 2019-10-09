package user

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// TokenExpiration time to exp a token
	TokenExpiration = time.Second * 3600 * 24 * 14
)

// AuthJWT implement Authentication interface
type AuthJWT struct {
	alg jwt.SigningMethod
	exp time.Duration
}
