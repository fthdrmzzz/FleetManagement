package api

import (
	"net/http"

	db "github.com/DevelopmentHiring/FatihDurmaz/db/sqlc"
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
	var req createDeliveryPointRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateDeliveryPointParams{
		ID:   req.ID,
		Name: req.Name,
	}

	deliveryPoint, err := s.db.CreateDeliveryPoint(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, deliveryPoint)
}

func (s *Server) createBags(ctx *gin.Context) {
	var req createBagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBagParams{
		Barcode:    req.Barcode,
		BagStatus:  db.BagsStatusCreated,
		DeliveryID: req.DeliveryID,
	}

	bag, err := s.db.CreateBag(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, bag)
}

func (s *Server) createPackages(ctx *gin.Context) {
	var req createPackageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePackageParams{
		Barcode:       req.Barcode,
		PackageStatus: db.PackagesStatusCreated,
		DeliveryID:    req.DeliveryID,
	}

	pk, err := s.db.CreatePackage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, pk)
}

func (s *Server) addPackageBag(ctx *gin.Context) {
	/*
		var req addPackageBagRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		arg := db.AddPackageBagParams{
			PackageBarcode: req.PackageBarcode,
			BagBarcode:     req.BagBarcode,
		}

		pk, err := s.db.CreatePackage(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, pk)*/
}

func (s *Server) addPackageVehicle(ctx *gin.Context) {

}
func (s *Server) addBagVehicle(ctx *gin.Context) {

}
