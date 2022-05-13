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

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/vehicles", server.createVehicle)
	router.POST("/deliverypoints", server.createDeliveryPoint)
	router.POST("/bags", server.createBags)
	router.POST("/packages", server.createPackages)

	router.POST("/packagebag/", server.addPackageBag)
	router.GET("/packagevehicle", server.addPackageVehicle)
	router.POST("/bagvehicle", server.addBagVehicle)

	server.router = router
}
