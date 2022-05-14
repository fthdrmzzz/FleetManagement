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
		s.l.Error.Printf("createVehicle: %s\n", errorResponse(err))
		return
	}

	arg := req.Plate
	vehicle, err := s.db.CreateVehicle(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("createVehicle: %s\n", errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
	s.l.Info.Printf("createVehicle: vehicle %s\n created", vehicle)
}

func (s *Server) createDeliveryPoints(ctx *gin.Context) {
	var req createDeliveryPointsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("createDeliveryPoints: %s\n", errorResponse(err))
		return
	}

	for _, deliveryPoint := range req.DeliveryPoints {
		arg := db.CreateDeliveryPointParams(deliveryPoint)
		deliveryPoint, err := s.db.CreateDeliveryPoint(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			s.l.Error.Printf("createDeliveryPoints: %s\n", errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, deliveryPoint)
		s.l.Info.Printf("createDeliveryPoints: delivery point %d, %s created\n", deliveryPoint.ID, deliveryPoint.Name)
	}
}

func (s *Server) createBags(ctx *gin.Context) {
	var req createBagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("createBags: %s\n", errorResponse(err))
		return
	}
	for _, bag := range req.Bags {
		arg := db.CreateBagParams{
			Barcode:    bag.Barcode,
			BagStatus:  db.BagsStatusCreated,
			DeliveryID: bag.DeliveryID,
		}

		bag, err := s.db.CreateBag(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			s.l.Error.Printf("createBags: %s\n", errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, bag)
		s.l.Info.Printf("createBags: bag %s created\n", bag.Barcode)
	}
}

func (s *Server) createPackages(ctx *gin.Context) {
	var req createPackageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("createPackages: %s\n", errorResponse(err))
		return
	}
	for _, pk := range req.Packages {
		arg := db.CreatePackageParams{
			Barcode:       pk.Barcode,
			PackageStatus: db.PackagesStatusCreated,
			PackageWeight: pk.PackageWeight,
			DeliveryID:    pk.DeliveryID,
		}

		pk, err := s.db.CreatePackage(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			s.l.Error.Printf("createPackages: %s\n", errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, pk)
		s.l.Info.Printf("createPackages: package %s created\n", pk.Barcode)
	}
}

func (s *Server) addPackageBag(ctx *gin.Context) {

	var req addPackagesToBagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("addPackageBag: %s\n", errorResponse(err))
		return
	}
	for _, assignment := range req.Assignments {
		//get package
		oldPackage, err := s.db.GetPackage(ctx, assignment.PackageBarcode)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			s.l.Error.Printf("addPackageBag: %s\n", errorResponse(err))
			return
		}
		//get bag
		bag, err := s.db.GetBag(ctx, assignment.BagBarcode)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			s.l.Error.Printf("addPackageBag: %s\n", errorResponse(err))
			return
		}
		// check if package is already loaded.
		if oldPackage.PackageStatus != db.PackagesStatusCreated {
			errorStr := errorResponse(fmt.Errorf("package %s is not available, status: %s", oldPackage.Barcode, oldPackage.PackageStatus))
			ctx.JSON(http.StatusBadRequest, errorStr)
			s.l.Error.Printf("addPackageBag: %s\n", errorStr)
			return
		}
		// check if package and bag delivery points match.
		if oldPackage.DeliveryID != bag.DeliveryID {
			errorStr := errorResponse(fmt.Errorf("delivery mismatch: package (%s->%d) - bag (%s->%d)", oldPackage.Barcode, oldPackage.DeliveryID, bag.Barcode, bag.DeliveryID))
			ctx.JSON(http.StatusBadRequest, errorStr)
			s.l.Error.Printf("addPackageBag: %s\n", errorStr)
			return
		}
		//check bag's status
		if bag.BagStatus != db.BagsStatusCreated {
			errorStr := errorResponse(fmt.Errorf("bag %s is not available, status: %s", bag.Barcode, bag.BagStatus))
			ctx.JSON(http.StatusBadRequest, errorStr)
			s.l.Error.Printf("addPackageBag: %s\n", errorStr)
			return
		}

		//update package
		argUpdate := db.UpdatePackageParams{
			Barcode:       oldPackage.Barcode,
			PackageStatus: db.PackagesStatusLoadedToBag,
		}
		s.db.UpdatePackage(ctx, argUpdate)

		argAssign := db.AddPackageBagParams{
			PackageBarcode: assignment.PackageBarcode,
			BagBarcode:     assignment.BagBarcode,
		}
		pk, err := s.db.AddPackageBag(ctx, argAssign)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, pk)
		s.l.Info.Printf("addPackageBag: package %s added to bag %s \n", pk.PackageBarcode, pk.BagBarcode)

	}

}

func (s *Server) makeDelivery(ctx *gin.Context) {
	var req makeDeliveryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("makeDelivery: %s \n", errorResponse(err))
		return
	}

	vehicle := req.Plate
	// load the shipments.
	for _, route := range req.Routes {
		curDeliveryPoint := int32(route.DeliveryPoint)
		for _, delivery := range route.Deliveries {
			curBag, err := s.db.GetBag(ctx, delivery.Barcode)
			//if bag
			if err == nil {
				// handle bag load to vehicle
				if err := s.handleBagLoad(ctx, bagLoadParams{
					Bag:        curBag,
					Vehicle:    vehicle,
					DeliveryId: curDeliveryPoint,
				}); err != nil {
					s.l.Error.Printf("makeDelivery: Error while loading bag %s into vehicle %s: %s\n", curBag.Barcode, vehicle, errorResponse(err))
					continue
				}
				s.l.Info.Printf("makeDelivery: bag %s loaded into vehicle %s\n", curBag.Barcode, vehicle)
			} else { // if package
				// load package
				curPackage, err := s.db.GetPackage(ctx, delivery.Barcode)
				if err != nil {
					s.l.Error.Printf("makeDelivery: Error while loading package %s into vehicle %s: %s\n", curPackage.Barcode, vehicle, errorResponse(err))
					continue
				}

				if _, err := s.db.UpdatePackage(ctx, db.UpdatePackageParams{
					Barcode:       curPackage.Barcode,
					PackageStatus: db.PackagesStatusLoaded,
				}); err != nil {
					s.l.Error.Printf("makeDelivery: Error while loading package %s into vehicle %s: %s\n", curPackage.Barcode, vehicle, errorResponse(err))
					continue
				}
				s.l.Info.Printf("makeDelivery: package %s loaded into vehicle %s\n", curPackage.Barcode, vehicle)
			}
		}
	}
	// unload the shipments
	for _, route := range req.Routes {
		curDeliveryPoint := int32(route.DeliveryPoint)
		for _, delivery := range route.Deliveries {
			curBag, err := s.db.GetBag(ctx, delivery.Barcode)
			//if bag
			if err == nil {
				// handle bag unload from vehicle
				if err := s.handleBagUnload(ctx, bagLoadParams{
					Bag:        curBag,
					Vehicle:    vehicle,
					DeliveryId: int32(route.DeliveryPoint),
				}); err != nil {
					s.l.Error.Printf("makeDelivery: Error while unloading bag %s:  %s\n", curBag.Barcode, errorResponse(err))
					continue
				}
				s.l.Info.Printf("makeDelivery: bag %s and its packages unloaded\n", curBag.Barcode)
			} else { //if package
				curPackage, err := s.db.GetPackage(ctx, delivery.Barcode)
				if err != nil {
					s.l.Error.Printf("makeDelivery: Error while unloading package %s: %s\n", curPackage.Barcode, errorResponse(err))
					continue
				}
				if err := s.handlePackageUnload(ctx, packageLoadParams{
					Package:    curPackage,
					DeliveryId: (curDeliveryPoint),
				}); err != nil {
					s.l.Error.Printf("makeDelivery: Error while unloading package %s: %s\n", curPackage.Barcode, errorResponse(err))
				}

				s.l.Info.Printf("makeDelivery: package %s unloaded\n", curPackage.Barcode)
			}
		}
	}
}

type bagLoadParams struct {
	Bag        db.Bag
	Vehicle    string
	DeliveryId int32
}

func (s *Server) handleBagUnload(ctx *gin.Context, arg bagLoadParams) error {
	if arg.DeliveryId != arg.Bag.DeliveryID {
		s.l.Warn.Printf("handleBagUnload: delivery mismatch bag (%s->%d) - delivery (%d)", arg.Bag.Barcode, arg.Bag.DeliveryID, arg.DeliveryId)
		return nil
	}
	switch arg.DeliveryId {
	//if type Branch
	case 1:
		// No change.
	case 2: //if type Distribution Center
		// change bag status
		s.db.UpdateBag(ctx, db.UpdateBagParams{
			Barcode:   arg.Bag.Barcode,
			BagStatus: db.BagsStatusUnloaded})
		// change packages status
		packageBagList, err := s.db.ListPackagesInBag(ctx, arg.Bag.Barcode)
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
		if err = s.db.DeleteBagVehicle(ctx, arg.Bag.Barcode); err != nil {
			return err
		}
		return nil
	case 3: //if type transfer center
		s.db.UpdateBag(ctx, db.UpdateBagParams{
			Barcode:   arg.Bag.Barcode,
			BagStatus: db.BagsStatusUnloaded})
		// change packages status
		packageBagList, err := s.db.ListPackagesInBag(ctx, arg.Bag.Barcode)
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
		if err = s.db.DeleteBagVehicle(ctx, arg.Bag.Barcode); err != nil {
			return err
		}
		return nil

	}
	return nil
}

func (s *Server) handleBagLoad(ctx *gin.Context, arg bagLoadParams) error {

	s.db.AddBagVehicle(ctx, db.AddBagVehicleParams{
		BagBarcode:   arg.Bag.Barcode,
		VehiclePlate: arg.Vehicle,
	})

	s.db.UpdateBag(ctx, db.UpdateBagParams{
		Barcode:   arg.Bag.Barcode,
		BagStatus: db.BagsStatusLoaded})

	packageBagList, err := s.db.ListPackagesInBag(ctx, arg.Bag.Barcode)
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

type packageLoadParams struct {
	Package    db.Package
	DeliveryId int32
}

func (s *Server) handlePackageUnload(ctx *gin.Context, arg packageLoadParams) error {
	if arg.DeliveryId != arg.Package.DeliveryID {
		return nil
		// TODO: log
	}
	switch arg.DeliveryId {
	case 1: // if type branch
		packageBagList, err := s.db.ListPackageBag(ctx)
		if err != nil {
			//TODO: log
			fmt.Println("")
		}
		for _, relation := range packageBagList {
			if relation.PackageBarcode == arg.Package.Barcode {
				return nil
				// TODO: this problem, if loaded into bag, do not unload
			}
		}
		s.db.UpdatePackage(ctx, db.UpdatePackageParams{
			Barcode:       arg.Package.Barcode,
			PackageStatus: db.PackagesStatusUnloaded,
		})
	case 2: //if type Distribution Center
		s.db.UpdatePackage(ctx, db.UpdatePackageParams{
			Barcode:       arg.Package.Barcode,
			PackageStatus: db.PackagesStatusUnloaded,
		})
	case 3: //if type transfer center
		packageBagList, err := s.db.ListPackageBag(ctx)
		if err != nil {
			//TODO: log
			fmt.Println("")
		}
		for _, relation := range packageBagList {
			if relation.PackageBarcode == arg.Package.Barcode {
				s.db.UpdatePackage(ctx, db.UpdatePackageParams{
					Barcode:       arg.Package.Barcode,
					PackageStatus: db.PackagesStatusUnloaded,
				})
			}
		}

	}
	return nil

}
