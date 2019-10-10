package user

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// TokenExpiration time to exp a token
	TokenExpiration = time.Second * 3600 * 24 * 14

	// TODO: move to env var
	// TokenSignedString string to sign token
	TokenSignedString = []byte("random-password")
)

// ClaimsJWT model
type ClaimsJWT struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// generate token
func generateToken(id int64, username string) (string, error) {

	claims := ClaimsJWT{ID: id, Username: username}
	claims.ExpiresAt = time.Now().Add(TokenExpiration).Unix()
	claims.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(TokenSignedString)
}
