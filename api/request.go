package api

type createBagRequest struct {
	Barcode    string `json:"barcode"`
	DeliveryID int32  `json:"delivery_id"`
}

type createDeliveryPointRequest struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type createPackageRequest struct {
	Barcode       string `json:"barcode"`
	PackageWeight int32  `json:"package_weight"`
	DeliveryID    int32  `json:"delivery_id"`
}
type createVehicleRequest struct {
	Plate string `json:"plate"`
}

type addPackageBagRequest struct {
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
