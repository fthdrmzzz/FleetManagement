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
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("delivery points do not match: package %d - bag %d", oldPackage.DeliveryID, bag.DeliveryID)))
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
		PackageStatus: db.PackagesStatusLoadedToBag,
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
	var req makeDeliveryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	vehicle := req.Plate
	// load the shipments.
	for _, route := range req.Routes {
		for _, delivery := range route.Deliveries {
			curBag, err := s.db.GetBag(ctx, delivery.Barcode)
			if err == nil { //if bag
				// handle bag load to vehicle
				if err := s.handleBagLoad(ctx, db.AddBagVehicleParams{
					BagBarcode:   curBag.Barcode,
					VehiclePlate: vehicle,
				}); err != nil {
					fmt.Printf("Error while loading bag into vehicle %s", err.Error())
				}

			} else {
				curPackage, _ := s.db.GetPackage(ctx, delivery.Barcode)
				s.db.UpdatePackage(ctx, db.UpdatePackageParams{
					Barcode:       curPackage.Barcode,
					PackageStatus: db.PackagesStatusLoaded,
				})
			}
		}
	}
	// unload the shipments
	for _, route := range req.Routes {
		curDeliveryPoint := route.DeliveryPoint
		for _, delivery := range route.Deliveries {
			curBag, err := s.db.GetBag(ctx, delivery.Barcode)
			if err == nil { //if bag
				if curDeliveryPoint == int(curBag.DeliveryID) {
					// handle bag unload from vehicle
					if err := s.handleBagUnload(ctx, db.AddBagVehicleParams{
						BagBarcode:   curBag.Barcode,
						VehiclePlate: vehicle,
					}); err != nil {
						fmt.Printf("Error while unloading bag %s", err.Error())
					}
				}
			} else {
				curPackage, _ := s.db.GetPackage(ctx, delivery.Barcode)
				s.db.UpdatePackage(ctx, db.UpdatePackageParams{
					Barcode:       curPackage.Barcode,
					PackageStatus: db.PackagesStatusLoaded,
				})
			}
		}
	}
}
func (s *Server) handleBagUnload(ctx *gin.Context, arg db.AddBagVehicleParams) error {
	// change bag status
	s.db.UpdateBag(ctx, db.UpdateBagParams{
		Barcode:   arg.BagBarcode,
		BagStatus: db.BagsStatusUnloaded})
	// change packages status
	packageBagList, err := s.db.ListPackagesInBag(ctx, arg.BagBarcode)
	if err != nil {
		return err
	}
	for _, pkBag := range packageBagList {
		s.db.UpdatePackage(ctx, db.UpdatePackageParams{
			Barcode:       pkBag.PackageBarcode,
			PackageStatus: db.PackagesStatusUnloaded,
		})
	}
	// delete bag vehicle
	if err = s.db.DeleteBagVehicle(ctx, arg.BagBarcode); err != nil {
		return err
	}
	return nil
}

func (s *Server) handleBagLoad(ctx *gin.Context, arg db.AddBagVehicleParams) error {

	s.db.AddBagVehicle(ctx, arg)

	s.db.UpdateBag(ctx, db.UpdateBagParams{
		Barcode:   arg.BagBarcode,
		BagStatus: db.BagsStatusLoaded})

	packageBagList, err := s.db.ListPackagesInBag(ctx, arg.BagBarcode)
	if err != nil {
		return err
	}
	for _, pkBag := range packageBagList {
		s.db.UpdatePackage(ctx, db.UpdatePackageParams{
			Barcode:       pkBag.PackageBarcode,
			PackageStatus: db.PackagesStatusLoaded,
		})
	}

	return nil
}
