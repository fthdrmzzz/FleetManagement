package api

import (
	"fmt"
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

	var req addPackageBagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//get package
	oldPackage, err := s.db.GetPackage(ctx, req.PackageBarcode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//get bag
	bag, err := s.db.GetBag(ctx, req.BagBarcode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// check if package is already loaded.
	if oldPackage.PackageStatus != db.PackagesStatusCreated {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("package is not available, status: %s", oldPackage.PackageStatus)))
		return
	}
	// check if package and bag delivery points match.
	if oldPackage.DeliveryID != bag.DeliveryID {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("pelivery Points do not match: %d-%d", oldPackage.DeliveryID, bag.DeliveryID)))
		return
	}
	//check bag's status
	if bag.BagStatus != db.BagsStatusCreated {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("bag is not available, status: %s", bag.BagStatus)))
		return
	}

	//update package
	argUpdate := db.UpdatePackageParams{
		Barcode:       oldPackage.Barcode,
		PackageStatus: db.PackagesStatusLoaded,
	}
	s.db.UpdatePackage(ctx, argUpdate)

	argAssign := db.AddPackageBagParams{
		PackageBarcode: req.PackageBarcode,
		BagBarcode:     req.BagBarcode,
	}
	pk, err := s.db.AddPackageBag(ctx, argAssign)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, pk)
}

func (s *Server) makeDelivery(ctx *gin.Context) {

}
