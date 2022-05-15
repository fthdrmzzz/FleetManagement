package main

import (
	"database/sql"

	"github.com/DevelopmentHiring/FatihDurmaz/api"
	"github.com/DevelopmentHiring/FatihDurmaz/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:postgres@host.docker.internal:5432/trendyol?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
	migrations    = "file://db/migrations"
)

func main() {
	l := logger.New()
	if err := setupDatabaseSchema(); err != nil {
		l.Error.Fatal("Could not setup db:", err)
	}
	l.Info.Println("Database migrations are done")

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		l.Error.Fatal("Could not connect to db:", err)
	}
	l.Info.Println("Database connection is done")

	server := api.NewServer(l, conn)
	err = server.Start(serverAddress)
	if err != nil {
		l.Error.Fatal("Could not start server:", err)
	}
}

func setupDatabaseSchema() error {
	m, err := migrate.New(
		migrations,
		dbSource)
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}
