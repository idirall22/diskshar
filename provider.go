package user

import (
	"context"
	"database/sql"
)

// Default sql db used
var driverName = "postgres"

// Provider interface
type Provider interface {
	New(context.Context, string, string, string, string, string, string) (*User, error)
	Get(context.Context, int64, string, string) (*User, error)
	Update(context.Context, *User) error
	Delete(ctx context.Context, id int64) error
}

// RegisterProvider register a provider
func getProvider(db *sql.DB, tableName string) (Provider, error) {

	switch driverName {

	case "postgres":
		return &PostgresProvider{db: db, tableName: tableName}, nil
	}
	return nil, nil
}

// SetDriverName set driver name
func SetDriverName(name string) {
	driverName = name
}
