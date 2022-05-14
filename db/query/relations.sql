-- name: AddPackageBag :one
INSERT INTO package_bag (
    bag_barcode,
    package_barcode
) VALUES (
    $1,$2
) RETURNING *;

-- name: AddBagVehicle :one
INSERT INTO bag_vehicle (
    bag_barcode,
    vehicle_plate
) VALUES (
    $1,$2
) RETURNING *;

-- name: AddPackageVehicle :one
INSERT INTO package_vehicle (
    package_barcode,
    vehicle_plate
) VALUES (
    $1,$2
) RETURNING *;

-- name: GetPackageBag :one
SELECT * FROM package_bag
WHERE package_barcode = $1  LIMIT 1;

-- name: ListPackagesInBag :many
SELECT * FROM package_bag
WHERE bag_barcode =$1 ;

-- name: ListPackageBag :many
SELECT * FROM package_bag;

-- name: GetBagVehicle :one
SELECT * FROM bag_vehicle
WHERE bag_barcode = $1  LIMIT 1;

-- name: GetPackageVehicle :one
SELECT * FROM package_vehicle
WHERE package_barcode = $1  LIMIT 1;

-- name: DeletePackageBag :exec
DELETE FROM package_bag WHERE package_barcode =$1;

-- name: DeleteBagVehicle :exec
DELETE FROM bag_vehicle WHERE bag_barcode =$1;

-- name: DeletePackageVehicle :exec
DELETE FROM package_vehicle WHERE package_barcode =$1;






