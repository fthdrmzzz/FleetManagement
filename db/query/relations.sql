-- name: AddPackageBag :one
INSERT INTO package_bag (
    bag_barcode,
    package_barcode
) VALUES (
    $1,$2
) RETURNING *;

-- name: AddBagVehicle :one
INSERT INTO vehicle_bag (
    bag_barcode,
    vehicle_plate
) VALUES (
    $1,$2
) RETURNING *;

-- name: AddPackageVehicle :one
INSERT INTO vehicle_package (
    package_barcode,
    vehicle_plate
) VALUES (
    $1,$2
) RETURNING *;

-- name: GetPackageBag :one
SELECT * FROM package_bag
WHERE package_barcode = $1  LIMIT 1;

-- name: GetBagVehicle :one
SELECT * FROM vehicle_bag
WHERE bag_barcode = $1  LIMIT 1;

-- name: GetPackageVehicle :one
SELECT * FROM vehicle_package
WHERE package_barcode = $1  LIMIT 1;

-- name: DeletePackageBag :exec
DELETE FROM package_bag WHERE package_barcode =$1;

-- name: DeleteBagVehicle :exec
DELETE FROM vehicle_bag WHERE bag_barcode =$1;

-- name: DeletePackageVehicle :exec
DELETE FROM vehicle_package WHERE package_barcode =$1;






