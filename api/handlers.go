package api

import (
	"fmt"
	"net/http"

	db "github.com/DevelopmentHiring/FatihDurmaz/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (s *Server) createVehicle(ctx *gin.Context) {
	var req createVehicleRequest
	// Format Check
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("createVehicle: %s\n", errorResponse(err))
		return
	}

	// DB query
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
	// Format Check
	var req createDeliveryPointsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("createDeliveryPoints: %s\n", errorResponse(err))
		return
	}

	// for each deliveryPoint in request
	response := createDeliveryPointsRequest{}
	for _, deliveryPoint := range req.DeliveryPoints {
		// make DB Query
		arg := db.CreateDeliveryPointParams(deliveryPoint)
		createdDeliveryPoint, err := s.db.CreateDeliveryPoint(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			s.l.Error.Printf("createDeliveryPoints: %s\n", errorResponse(err))
			return
		}
		fmt.Println(createdDeliveryPoint)
		s.l.Info.Printf("createDeliveryPoints: delivery point %d, %s created\n", createdDeliveryPoint.ID, createdDeliveryPoint.Name)
		response.DeliveryPoints = append(response.DeliveryPoints, deliveryPoint)
	}
	ctx.JSON(http.StatusOK, response)
}

func (s *Server) createBags(ctx *gin.Context) {
	// format check
	var req createBagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("createBags: %s\n", errorResponse(err))
		return
	}

	// for each bag in request
	var bags []db.Bag
	for _, bag := range req.Bags {
		arg := db.CreateBagParams{
			Barcode:    bag.Barcode,
			BagStatus:  db.BagsStatusCreated,
			DeliveryID: bag.DeliveryID,
		}
		// make DB query
		bag, err := s.db.CreateBag(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			s.l.Error.Printf("createBags: %s\n", errorResponse(err))
			return
		}
		bags = append(bags, bag)
		s.l.Info.Printf("createBags: bag %s created\n", bag.Barcode)
	}
	ctx.JSON(http.StatusOK, bags)
}

func (s *Server) createPackages(ctx *gin.Context) {
	// format check
	var req createPackageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("createPackages: %s\n", errorResponse(err))
		return
	}

	// for each package in request
	var pkSlice []db.Package
	for _, pk := range req.Packages {
		arg := db.CreatePackageParams{
			Barcode:       pk.Barcode,
			PackageStatus: db.PackagesStatusCreated,
			PackageWeight: pk.PackageWeight,
			DeliveryID:    pk.DeliveryID,
		}
		// make DB query
		pk, err := s.db.CreatePackage(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			s.l.Error.Printf("createPackages: %s\n", errorResponse(err))
			return
		}
		pkSlice = append(pkSlice, pk)
		s.l.Info.Printf("createPackages: package %s created\n", pk.Barcode)
	}
	ctx.JSON(http.StatusOK, pkSlice)
}

func (s *Server) addPackageBag(ctx *gin.Context) {
	// format check
	var req addPackagesToBagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("addPackageBag: %s\n", errorResponse(err))
		return
	}

	// for each assignment in query
	var pkSlice []db.PackageBag
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

		// make DB query
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
		pkSlice = append(pkSlice, pk)
		s.l.Info.Printf("addPackageBag: package %s added to bag %s \n", pk.PackageBarcode, pk.BagBarcode)

	}
	ctx.JSON(http.StatusOK, pkSlice)

}

func (s *Server) makeDelivery(ctx *gin.Context) {
	// format check
	var req makeDeliveryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		s.l.Error.Printf("makeDelivery: %s \n", errorResponse(err))
		return
	}

	vehicle := req.Plate

	// For each Delivery Route
	for i, route := range req.Routes {
		curDeliveryPoint := int32(route.DeliveryPoint)
		// For each shipment in route
		for j, delivery := range route.Deliveries {
			curBag, err := s.db.GetBag(ctx, delivery.Barcode)
			// if the shipment is bag
			if err == nil {
				// handle bag load to vehicle
				if state, err := s.handleBagLoad(ctx, bagLoadParams{
					Bag:        curBag,
					Vehicle:    vehicle,
					DeliveryId: curDeliveryPoint,
				}); err != nil {
					s.l.Error.Printf("makeDelivery: Error while loading bag %s into vehicle %s: %s\n", curBag.Barcode, vehicle, errorResponse(err))
					ctx.JSON(http.StatusInternalServerError, err)
					return
				} else {
					req.Routes[i].Deliveries[j].State = state
					s.l.Info.Printf("makeDelivery: bag %s loaded into vehicle %s\n", curBag.Barcode, vehicle)
				}

			} else { // if the shipment is package

				curPackage, err := s.db.GetPackage(ctx, delivery.Barcode)
				if err != nil {
					s.l.Error.Printf("makeDelivery: Error while loading package %s into vehicle %s: %s\n", curPackage.Barcode, vehicle, errorResponse(err))
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
				// handle package load to vehicle
				if _, err := s.db.UpdatePackage(ctx, db.UpdatePackageParams{
					Barcode:       curPackage.Barcode,
					PackageStatus: db.PackagesStatusLoaded,
				}); err != nil {
					s.l.Error.Printf("makeDelivery: Error while loading package %s into vehicle %s: %s\n", curPackage.Barcode, vehicle, errorResponse(err))
					ctx.JSON(http.StatusInternalServerError, err)
					return
				} else {
					req.Routes[i].Deliveries[j].State = string(db.PackagesStatusLoaded)
					s.l.Info.Printf("makeDelivery: package %s loaded into vehicle %s\n", curPackage.Barcode, vehicle)
				}
			}
		}

	}
	// For Each Delivery Point
	for i, route := range req.Routes {
		curDeliveryPoint := int32(route.DeliveryPoint)
		// For each shipment
		for j, delivery := range route.Deliveries {
			curBag, err := s.db.GetBag(ctx, delivery.Barcode)
			// if the shipment is bag
			if err == nil {
				// handle bag unload from vehicle
				if state, err := s.handleBagUnload(ctx, bagLoadParams{
					Bag:        curBag,
					Vehicle:    vehicle,
					DeliveryId: int32(route.DeliveryPoint),
				}); err != nil {
					s.l.Error.Printf("makeDelivery: Error while unloading bag %s:  %s\n", curBag.Barcode, errorResponse(err))
					ctx.JSON(http.StatusInternalServerError, err)
					return
				} else {
					req.Routes[i].Deliveries[j].State = state
				}
			} else { //if the shipment is package
				curPackage, err := s.db.GetPackage(ctx, delivery.Barcode)
				if err != nil {
					s.l.Error.Printf("makeDelivery: Error while unloading package %s: %s\n", curPackage.Barcode, errorResponse(err))
					ctx.JSON(http.StatusInternalServerError, err)
					return
				}
				// handle package unload from vehicle
				if state, err := s.handlePackageUnload(ctx, packageLoadParams{
					Package:    curPackage,
					DeliveryId: (curDeliveryPoint),
				}); err != nil {
					s.l.Error.Printf("makeDelivery: Error while unloading package %s: %s\n", curPackage.Barcode, errorResponse(err))
					ctx.JSON(http.StatusInternalServerError, err)
					return
				} else {
					req.Routes[i].Deliveries[j].State = state
				}
			}
		}
	}
	ctx.JSON(http.StatusOK, req)
}

type bagLoadParams struct {
	Bag        db.Bag
	Vehicle    string
	DeliveryId int32
}

func (s *Server) handleBagUnload(ctx *gin.Context, arg bagLoadParams) (string, error) {
	if arg.DeliveryId != arg.Bag.DeliveryID {
		s.l.Warn.Printf("handleBagUnload: delivery mismatch bag (%s->%d) - delivery (%d)\n", arg.Bag.Barcode, arg.Bag.DeliveryID, arg.DeliveryId)
		return string(db.BagsStatusLoaded), nil
	}
	switch arg.DeliveryId {

	case 1: //if type Branch
		// No change.
		s.l.Warn.Printf("handleBagUnload: Can not unload bag %s in branch \n", arg.Bag.Barcode)
		return string(db.BagsStatusLoaded), nil
	case 2: //if type Distribution Center

		// change bag status
		if _, err := s.db.UpdateBag(ctx, db.UpdateBagParams{
			Barcode:   arg.Bag.Barcode,
			BagStatus: db.BagsStatusUnloaded}); err != nil {
			return "", err
		}
		// change packages status
		if packageBagList, err := s.db.ListPackagesInBag(ctx, arg.Bag.Barcode); err != nil {
			return "", err
		} else {
			for _, pkBag := range packageBagList {
				if _, err := s.db.UpdatePackage(ctx, db.UpdatePackageParams{
					Barcode:       pkBag.PackageBarcode,
					PackageStatus: db.PackagesStatusUnloaded,
				}); err != nil {
					return "", err
				}
			}
			// delete bag vehicle
			if err = s.db.DeleteBagVehicle(ctx, arg.Bag.Barcode); err != nil {
				return "", err
			}
		}
	case 3: //if type transfer center
		if _, err := s.db.UpdateBag(ctx, db.UpdateBagParams{
			Barcode:   arg.Bag.Barcode,
			BagStatus: db.BagsStatusUnloaded}); err != nil {
			return "", err
		}
		// change packages status
		if packageBagList, err := s.db.ListPackagesInBag(ctx, arg.Bag.Barcode); err != nil {
			return "", err
		} else {
			for _, pkBag := range packageBagList {
				if _, err := s.db.UpdatePackage(ctx, db.UpdatePackageParams{
					Barcode:       pkBag.PackageBarcode,
					PackageStatus: db.PackagesStatusUnloaded,
				}); err != nil {
					return "", err
				}
			}
			// delete bag vehicle
			if err = s.db.DeleteBagVehicle(ctx, arg.Bag.Barcode); err != nil {
				return "", err
			}
		}
	}
	s.l.Info.Printf("makeDelivery: bag %s and its packages unloaded\n", arg.Bag.Barcode)
	return string(db.BagsStatusUnloaded), nil
}

func (s *Server) handleBagLoad(ctx *gin.Context, arg bagLoadParams) (string, error) {

	s.db.AddBagVehicle(ctx, db.AddBagVehicleParams{
		BagBarcode:   arg.Bag.Barcode,
		VehiclePlate: arg.Vehicle,
	})

	s.db.UpdateBag(ctx, db.UpdateBagParams{
		Barcode:   arg.Bag.Barcode,
		BagStatus: db.BagsStatusLoaded})

	packageBagList, err := s.db.ListPackagesInBag(ctx, arg.Bag.Barcode)
	if err != nil {
		return "", err
	}
	for _, pkBag := range packageBagList {
		if _, err := s.db.UpdatePackage(ctx, db.UpdatePackageParams{
			Barcode:       pkBag.PackageBarcode,
			PackageStatus: db.PackagesStatusLoaded,
		}); err != nil {
			return "", err
		}
	}

	return string(db.BagsStatusLoaded), nil
}

type packageLoadParams struct {
	Package    db.Package
	DeliveryId int32
}

func (s *Server) handlePackageUnload(ctx *gin.Context, arg packageLoadParams) (string, error) {
	if arg.DeliveryId != arg.Package.DeliveryID {
		s.l.Warn.Printf("handlePackageUnload: delivery mismatch package (%s->%d) - delivery (%d)\n", arg.Package.Barcode, arg.Package.DeliveryID, arg.DeliveryId)
		return string(db.PackagesStatusLoaded), nil
	}
	switch arg.DeliveryId {
	case 1: // if type branch
		if packageBagList, err := s.db.ListPackageBag(ctx); err != nil {
			return "", err
		} else {
			for _, relation := range packageBagList {
				if relation.PackageBarcode == arg.Package.Barcode {
					s.l.Warn.Printf("handlePackageUnload: Cannot unload package loaded into bag %s\n", arg.Package.Barcode)
					return string(db.PackagesStatusLoaded), nil
				}
			}
			if _, err := s.db.UpdatePackage(ctx, db.UpdatePackageParams{
				Barcode:       arg.Package.Barcode,
				PackageStatus: db.PackagesStatusUnloaded,
			}); err != nil {
				return "", err
			} else {
				s.l.Info.Printf("handlePackageUnload: package %s unloaded\n", arg.Package.Barcode)
				return string(db.PackagesStatusUnloaded), nil
			}
		}
	case 2: //if type Distribution Center
		if _, err := s.db.UpdatePackage(ctx, db.UpdatePackageParams{
			Barcode:       arg.Package.Barcode,
			PackageStatus: db.PackagesStatusUnloaded,
		}); err != nil {
			return "", err
		} else {
			s.l.Info.Printf("handlePackageUnload: package %s unloaded\n", arg.Package.Barcode)
			return string(db.PackagesStatusUnloaded), nil
		}
	case 3: //if type transfer center
		if packageBagList, err := s.db.ListPackageBag(ctx); err != nil {
			return "", err
		} else {
			for _, relation := range packageBagList {
				if relation.PackageBarcode == arg.Package.Barcode {
					if _, err := s.db.UpdatePackage(ctx, db.UpdatePackageParams{
						Barcode:       arg.Package.Barcode,
						PackageStatus: db.PackagesStatusUnloaded,
					}); err != nil {
						return "", err
					} else {
						s.l.Info.Printf("handlePackageUnload: package %s unloaded\n", arg.Package.Barcode)
						return string(db.PackagesStatusUnloaded), nil
					}
				}
			}
			// if package is not in bag,
			s.l.Warn.Printf("handlePackageUnload: package %s not in bag. cannot be unloaded in transfer center.\n", arg.Package.Barcode)
			return string(db.PackagesStatusLoaded), nil
		}
	}
	s.l.Info.Printf("handlePackageUnload: package %s unloaded\n", arg.Package.Barcode)
	return string(db.PackagesStatusUnloaded), nil
}
