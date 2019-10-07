package user

import (
	"database/sql"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// TokenExpiration time to exp a token
	TokenExpiration = time.Second * 3600 * 24 * 14
)

var (
	// ErrorDataBaseSS when database value is nil
	ErrorDataBaseSS = errors.New("Error database used to start service is NIL")
	// ErrorTableNameSS when table name used is empty
	ErrorTableNameSS = errors.New("Error table name not valid")
)

// service
var service *Service

// Service model
type Service struct {
	auth     Authentication
	provider Provider
}

// StartService start service user
func StartService(db *sql.DB, tableName string) error {

	if db == nil {
		return ErrorDataBaseSS
	}
	if tableName == "" {
		return ErrorTableNameSS
	}

	provider, err := getProvider(db, tableName)
	if err != nil {
		return err
	}

	auth := &AuthJWT{alg: &jwt.SigningMethodHMAC{}, exp: TokenExpiration}
	service = &Service{provider: provider, auth: auth}
	return nil
}
