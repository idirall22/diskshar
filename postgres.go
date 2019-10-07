package user

import (
	"context"
	"database/sql"
	"fmt"
)

// PostgresProvider postgres implementation of provider interface
type PostgresProvider struct {
	db        *sql.DB
	tableName string
}

// New created new user model
func (p *PostgresProvider) New(ctx context.Context,
	username, email, firstName, lastName, password, avatar string) (*User, error) {

	query := fmt.Sprintf(`INSERT INTO %s
		(username, first_name, last_name, email, password, avatar)
		VALUES ('%s','%s','%s','%s','%s','%s') RETURNING id, created_at`,
		p.tableName,
		username,
		firstName,
		lastName,
		email,
		password,
		avatar,
	)

	stmt, err := p.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	user := &User{}

	err = stmt.QueryRowContext(ctx).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return nil, err
	}
	user.Username = username
	user.Email = email
	user.LastName = lastName
	user.FirstName = firstName
	user.Password = password
	user.Avatar = avatar

	return user, nil
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

func (p *PostgresProvider) getDB() *sql.DB {
	return p.db
}

func (p *PostgresProvider) getTableName() string {
	return p.tableName
}
