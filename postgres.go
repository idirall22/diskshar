package user

import (
	"context"
	"database/sql"
)

// PostgresProvider postgres implementation of provider interface
type PostgresProvider struct {
	db        *sql.DB
	tableName string
}

// New created new user model
func (p *PostgresProvider) New(ctx context.Context,
	username, email, firstName, lastName, password, avatar string) (*User, error) {
	return nil, nil
}

// Get get a user model
func (p *PostgresProvider) Get(ctx context.Context,
	id int64, username, email string) (*User, error) {
	return nil, nil
}

// Update update user model
func (p *PostgresProvider) Update(ctx context.Context, user *User) error {
	return nil
}

// Delete delete user model
func (p *PostgresProvider) Delete(ctx context.Context, id int64) error {
	return nil
}
