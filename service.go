package user

import (
	"database/sql"
	"errors"
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

	service = &Service{provider: provider}
	return nil
}
