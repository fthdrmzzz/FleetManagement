package api

import (
	"database/sql"

	db "github.com/DevelopmentHiring/FatihDurmaz/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *db.Queries
	router *gin.Engine
}

func NewServer(conn *sql.DB) *Server {
	server := &Server{}
	query := db.New(conn)

	server.db = query
	server.setupRouter()
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/vehicles", server.createVehicle)
	router.POST("/deliverypoints", server.createDeliveryPoint)
	router.POST("/bags", server.createBags)
	router.POST("/packages", server.createPackages)

	router.POST("/assignpackage/", server.addPackageBag)
	router.POST("/makedelivery", server.makeDelivery)

	server.router = router
}
