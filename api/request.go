package api

import db "github.com/DevelopmentHiring/FatihDurmaz/db/sqlc"

type createBagRequest struct {
	Bags []db.Bag `json:"bags"`
}

type createDeliveryPointsRequest struct {
	DeliveryPoints []db.DeliveryPoint `json:"delivery_points"`
}

type createPackageRequest struct {
	Packages []db.Package `json:"packages"`
}
type createVehicleRequest struct {
	Plate string `json:"plate"`
}

type addPackagesToBagRequest struct {
	Assignments []assignment `json:"assignments"`
}
type assignment struct {
	BagBarcode     string `json:"bag_barcode"`
	PackageBarcode string `json:"package_barcode"`
}

type addPackageVehicleRequest struct {
	VehiclePlate   string `json:"vehicle_plate"`
	PackageBarcode string `json:"package_barcode"`
}

type addBagVehicleRequest struct {
	VehiclePlate string `json:"vehicle_plate"`
	BagBarcode   string `json:"bag_barcode"`
}

type makeDeliveryRequest struct {
	Plate  string  `json:"plate"`
	Routes []Route `json:"route"`
}

type Route struct {
	DeliveryPoint int        `json:"deliveryPoint"`
	Deliveries    []Delivery `json:"deliveries"`
}
type Delivery struct {
	Barcode string `json:"barcode"`
}
