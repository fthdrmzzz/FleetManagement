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
	// make migrations
	if err := setupDatabaseSchema(l); err != nil {
		l.Error.Fatal("Could not setup db:", err)
	}
	l.Info.Println("Database migrations are done")

	// db connection
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		l.Error.Fatal("Could not connect to db:", err)
	}
	l.Info.Println("Database connection is done")

	// start server
	server := api.NewServer(l, conn)
	err = server.Start(serverAddress)
	if err != nil {
		l.Error.Fatal("Could not start server:", err)
	}
}

func setupDatabaseSchema(l logger.Logging) error {
	m, err := migrate.New(
		migrations,
		dbSource)
	if err != nil {
		l.Error.Printf("Error in creating migrate object %s", err.Error())
		return err
	}
	if err := m.Down(); err != nil {
		l.Error.Printf("Error in migrate down %s", err.Error())
		if err.Error() != "no change" {
			return err
		}
	}
	if err := m.Up(); err != nil {
		l.Error.Printf("Error in migrate up %s", err.Error())
		return err
	}
	return nil
}
