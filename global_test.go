package user

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

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
var testDriverName = "postgres"

// Clean db test
func cleanDB(db *sql.DB) error {
	query := fmt.Sprintf(`
		DROP TABLE IF EXISTS users;

		CREATE TABLE IF NOT EXISTS users(
		    id SERIAL PRIMARY KEY,
		    username VARCHAR NOT NULL,
		    first_name VARCHAR,
		    last_name VARCHAR,
		    Email VARCHAR NOT NULL,
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

	db, err := sql.Open(testDriverName, dbInfos)
	if err != nil {
		return err
	}

	provider := &PostgresProvider{db: db, tableName: "users"}
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

	defer closeDB(testService.provider.(*PostgresProvider).getDB())

	// service.provider
	t.Run("new", testNew)
	t.Run("get", testGet)
	t.Run("update", testUpdate)
	t.Run("delete", testDelete)

}
