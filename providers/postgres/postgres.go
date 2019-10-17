package userAccount

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"text/template"

	"github.com/idirall22/user/models"
)

var (
	// ErrorParamsGetUser when parameters used in get func are: id=0, username=email=""
	ErrorParamsGetUser = errors.New("Error id should be > 0, username and email != empty string")
)

// PostgresProvider postgres implementation of provider interface
type PostgresProvider struct {
	DB        *sql.DB
	TableName string
}

// New created new user model
func (p *PostgresProvider) New(ctx context.Context,
	username, email, firstName, lastName, password, avatar string) (*models.User, error) {

	query := fmt.Sprintf(`INSERT INTO %s
		(username, first_name, last_name, email, password, avatar)
		VALUES ('%s','%s','%s','%s','%s','%s') RETURNING id, created_at`,
		p.TableName,
		username,
		firstName,
		lastName,
		email,
		password,
		avatar,
	)

	stmt, err := p.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	user := &models.User{}

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
func (p *PostgresProvider) Get(ctx context.Context, id int64, username, email string) (*models.User, error) {

	if id <= 0 && username == "" && email == "" {
		return nil, ErrorParamsGetUser
	}

	query := `
		{{if gt .id 0}}
			SELECT * FROM {{.tableName}} WHERE id={{.id}} AND deleted_at IS NULL
		{{else if (ne .username "")}}
			SELECT * FROM {{.tableName}} WHERE username='{{.username}}' AND deleted_at IS NULL
		{{else if (ne .email "")}}
			SELECT * FROM {{.tableName}} WHERE email='{{.email}}' AND deleted_at IS NULL
		{{else}}

		{{end}}
	`

	t := template.Must(template.New("query").Parse(query))

	data := make(map[string]interface{})
	data["tableName"] = p.TableName
	data["id"] = id
	data["username"] = username
	data["email"] = email

	buf := new(bytes.Buffer)
	t.Execute(buf, data)

	stmt, err := p.DB.Prepare(buf.String())

	if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = stmt.QueryRowContext(ctx).Scan(
		&user.ID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.CreatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update update user model
func (p *PostgresProvider) Update(ctx context.Context, id int64, firstName, lastName, avatar string) error {

	tx, err := p.DB.BeginTx(ctx, nil)
	defer tx.Rollback()

	if err != nil {
		return err
	}

	query := fmt.Sprintf(`SELECT EXISTS( SELECT 1 FROM %s WHERE id = %d)`, p.TableName, id)

	exists := false

	err = tx.QueryRowContext(ctx, query).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return sql.ErrNoRows
	}

	query = fmt.Sprintf(`
		UPDATE %s
		SET first_name = '%s', last_name = '%s', avatar = '%s'
		WHERE(id = %d)
	`, p.TableName, firstName, lastName, avatar, id)

	stmt, err := p.DB.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx)

	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

// Delete delete user model
func (p *PostgresProvider) Delete(ctx context.Context, id int64) error {

	tx, err := p.DB.BeginTx(ctx, nil)
	defer tx.Rollback()

	if err != nil {
		return err
	}

	query := fmt.Sprintf(`SELECT EXISTS( SELECT 1 FROM %s WHERE id = %d)`,
		p.TableName, id)

	exists := false

	err = tx.QueryRowContext(ctx, query).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return sql.ErrNoRows
	}

	query = fmt.Sprintf(`DELETE FROM users WHERE id=%d`, id)
	_, err = tx.ExecContext(ctx, query)

	if err != nil {
		return err
	}

	return nil
}

// GetDB return database
func (p *PostgresProvider) GetDB() *sql.DB {
	return p.DB
}

// GetTableName return table name
func (p *PostgresProvider) GetTableName() string {
	return p.TableName
}
