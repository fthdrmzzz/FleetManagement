package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createVehicle(ctx *gin.Context) {
	var req createVehicleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.Plate

	vehicle, err := s.db.CreateVehicle(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}
func (s *Server) createDeliveryPoint(ctx *gin.Context) {

}
func (s *Server) createBags(ctx *gin.Context) {

}
func (s *Server) createPackages(ctx *gin.Context) {

}

func (s *Server) addPackageBag(ctx *gin.Context) {

}
func (s *Server) addPackageVehicle(ctx *gin.Context) {

}
func (s *Server) addBagVehicle(ctx *gin.Context) {

}
