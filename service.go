package userAccount

import (
	"database/sql"
	"errors"

	pr "github.com/idirall22/user/providers/postgres"
)

var (
	// ErrorDataBaseSS when database value is nil
	ErrorDataBaseSS = errors.New("Error database used to start service is NIL")
	// ErrorTableNameSS when table name used is empty
	ErrorTableNameSS = errors.New("Error table name not valid")
)

// Service model
type Service struct {
	provider Provider
}

// StartService start service user
func StartService(db *sql.DB, tableName string) *Service {

	provider := &pr.PostgresProvider{DB: db, TableName: tableName}
	return &Service{provider: provider}
}
