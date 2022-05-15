package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	dbDriver   = "postgres"
	dbSource   = "postgresql://root:postgres@localhost:5432/trendyol?sslmode=disable"
	migrations = "file://../migrations"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	mgrt, err := migrate.New(
		migrations,
		dbSource)
	if err != nil {
		log.Fatal("Could not connect to db:", err)
	}
	if err := mgrt.Down(); err != nil {
		log.Fatal("Could not migrate db:", err)
	}
	if err := mgrt.Up(); err != nil {
		log.Fatal("Could not migrate db:", err)
	}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Could not connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
