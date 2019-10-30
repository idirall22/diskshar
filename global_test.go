package userAccount

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	pr "github.com/idirall22/user/providers/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "diskshar_test"
)

var testService *Service

var testTokenString = ""
var testPassword = "fdpjfd654/*sMLdf"

// Clean db test
func cleanDB(db *sql.DB) error {
	query := fmt.Sprintf(`
		DROP TABLE IF EXISTS users CASCADE;

		CREATE TABLE IF NOT EXISTS users(
		    id SERIAL PRIMARY KEY,
		    username VARCHAR UNIQUE NOT NULL,
		    first_name VARCHAR,
		    last_name VARCHAR,
		    Email VARCHAR UNIQUE NOT NULL,
		    password VARCHAR NOT NULL,
		    avatar VARCHAR,
		    created_at TIMESTAMP with TIME ZONE DEFAULT now(),
		    deleted_at TIMESTAMP DEFAULT NULL
		);
		`)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

func closeDB(db *sql.DB) {
	db.Close()
}

// Connect to db test
func connectDB() error {

	dbInfos := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfos)
	if err != nil {
		return err
	}

	provider := &pr.PostgresProvider{DB: db, TableName: "users"}
	testService = &Service{
		provider: provider,
	}
	err = cleanDB(db)
	if err != nil {
		return err
	}

	return nil
}

func TestGlobal(t *testing.T) {
	if err := connectDB(); err != nil {
		log.Fatal("Error connect database test, ", err)
	}

	defer closeDB(testService.provider.(*pr.PostgresProvider).GetDB())

	// t.Run("create user", testCreateUser)
	// t.Run("authenticate user", testAuthenticate)

	// handlers
	t.Run("register handler", testRegister)
	t.Run("login handler", testLogin)
	t.Run("middelware authenticate user", testAuthnticateUser)
	t.Run("validate token", testValidateToken)

}
