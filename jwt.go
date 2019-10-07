package user

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// AuthJWT implement Authentication interface
type AuthJWT struct {
	alg jwt.SigningMethod
	exp time.Duration
}

// login
func (a *AuthJWT) login(ctx context.Context,
	username, email, password string) (string, error) {
	return "", nil
}

// register
func (a *AuthJWT) register(ctx context.Context,
	username, email, password string) error {
	return nil
}
