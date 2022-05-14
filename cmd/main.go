package main

import (
	"database/sql"
	"log"

	"github.com/DevelopmentHiring/FatihDurmaz/api"
	"github.com/DevelopmentHiring/FatihDurmaz/logger"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:postgres@localhost:5432/trendyol?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	l := logger.New()
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Could not connect to db:", err)
	}
	server := api.NewServer(l, conn)
	err = server.Start(serverAddress)
}
