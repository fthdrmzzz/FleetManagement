package api

import (
	"database/sql"

	db "github.com/DevelopmentHiring/FatihDurmaz/db/sqlc"
	"github.com/DevelopmentHiring/FatihDurmaz/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *db.Queries
	router *gin.Engine
	l      logger.Logging
}

func NewServer(l logger.Logging, conn *sql.DB) *Server {
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
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/vehicles", server.createVehicle)
	router.POST("/deliverypoints", server.createDeliveryPoints)
	router.POST("/bags", server.createBags)
	router.POST("/packages", server.createPackages)
	router.POST("/assignpackages/", server.addPackageBag)
	router.POST("/makedelivery", server.makeDelivery)
	server.router = router
}
