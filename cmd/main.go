package main

import (
	"database/sql"
	"log"

	"github.com/DevelopmentHiring/FatihDurmaz/api"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:postgres@localhost:5432/trendyol?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Could not connect to db:", err)
	}

	server := api.NewServer(conn)
	err = server.Start(serverAddress)
}
