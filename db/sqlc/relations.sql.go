// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: relations.sql

package db

import (
	"context"
)

const addBagVehicle = `-- name: AddBagVehicle :one
INSERT INTO bag_vehicle (
    bag_barcode,
    vehicle_plate
) VALUES (
    $1,$2
) RETURNING vehicle_plate, bag_barcode
`

type AddBagVehicleParams struct {
	BagBarcode   string `json:"bag_barcode"`
	VehiclePlate string `json:"vehicle_plate"`
}

func (q *Queries) AddBagVehicle(ctx context.Context, arg AddBagVehicleParams) (BagVehicle, error) {
	row := q.db.QueryRowContext(ctx, addBagVehicle, arg.BagBarcode, arg.VehiclePlate)
	var i BagVehicle
	err := row.Scan(&i.VehiclePlate, &i.BagBarcode)
	return i, err
}

const addPackageBag = `-- name: AddPackageBag :one
INSERT INTO package_bag (
    bag_barcode,
    package_barcode
) VALUES (
    $1,$2
) RETURNING bag_barcode, package_barcode
`

type AddPackageBagParams struct {
	BagBarcode     string `json:"bag_barcode"`
	PackageBarcode string `json:"package_barcode"`
}

func (q *Queries) AddPackageBag(ctx context.Context, arg AddPackageBagParams) (PackageBag, error) {
	row := q.db.QueryRowContext(ctx, addPackageBag, arg.BagBarcode, arg.PackageBarcode)
	var i PackageBag
	err := row.Scan(&i.BagBarcode, &i.PackageBarcode)
	return i, err
}

const addPackageVehicle = `-- name: AddPackageVehicle :one
INSERT INTO package_vehicle (
    package_barcode,
    vehicle_plate
) VALUES (
    $1,$2
) RETURNING vehicle_plate, package_barcode
`

type AddPackageVehicleParams struct {
	PackageBarcode string `json:"package_barcode"`
	VehiclePlate   string `json:"vehicle_plate"`
}

func (q *Queries) AddPackageVehicle(ctx context.Context, arg AddPackageVehicleParams) (PackageVehicle, error) {
	row := q.db.QueryRowContext(ctx, addPackageVehicle, arg.PackageBarcode, arg.VehiclePlate)
	var i PackageVehicle
	err := row.Scan(&i.VehiclePlate, &i.PackageBarcode)
	return i, err
}

const deleteBagVehicle = `-- name: DeleteBagVehicle :exec
DELETE FROM bag_vehicle WHERE bag_barcode =$1
`

func (q *Queries) DeleteBagVehicle(ctx context.Context, bagBarcode string) error {
	_, err := q.db.ExecContext(ctx, deleteBagVehicle, bagBarcode)
	return err
}

const deletePackageBag = `-- name: DeletePackageBag :exec
DELETE FROM package_bag WHERE package_barcode =$1
`

func (q *Queries) DeletePackageBag(ctx context.Context, packageBarcode string) error {
	_, err := q.db.ExecContext(ctx, deletePackageBag, packageBarcode)
	return err
}

const deletePackageVehicle = `-- name: DeletePackageVehicle :exec
DELETE FROM package_vehicle WHERE package_barcode =$1
`

func (q *Queries) DeletePackageVehicle(ctx context.Context, packageBarcode string) error {
	_, err := q.db.ExecContext(ctx, deletePackageVehicle, packageBarcode)
	return err
}

const getBagVehicle = `-- name: GetBagVehicle :one
SELECT vehicle_plate, bag_barcode FROM bag_vehicle
WHERE bag_barcode = $1  LIMIT 1
`

func (q *Queries) GetBagVehicle(ctx context.Context, bagBarcode string) (BagVehicle, error) {
	row := q.db.QueryRowContext(ctx, getBagVehicle, bagBarcode)
	var i BagVehicle
	err := row.Scan(&i.VehiclePlate, &i.BagBarcode)
	return i, err
}

const getPackageBag = `-- name: GetPackageBag :one
SELECT bag_barcode, package_barcode FROM package_bag
WHERE package_barcode = $1  LIMIT 1
`

func (q *Queries) GetPackageBag(ctx context.Context, packageBarcode string) (PackageBag, error) {
	row := q.db.QueryRowContext(ctx, getPackageBag, packageBarcode)
	var i PackageBag
	err := row.Scan(&i.BagBarcode, &i.PackageBarcode)
	return i, err
}

const getPackageVehicle = `-- name: GetPackageVehicle :one
SELECT vehicle_plate, package_barcode FROM package_vehicle
WHERE package_barcode = $1  LIMIT 1
`

func (q *Queries) GetPackageVehicle(ctx context.Context, packageBarcode string) (PackageVehicle, error) {
	row := q.db.QueryRowContext(ctx, getPackageVehicle, packageBarcode)
	var i PackageVehicle
	err := row.Scan(&i.VehiclePlate, &i.PackageBarcode)
	return i, err
}

const listPackagesInBag = `-- name: ListPackagesInBag :many
SELECT bag_barcode, package_barcode FROM package_bag
ORDER BY bag_barcode =$1
`

func (q *Queries) ListPackagesInBag(ctx context.Context, bagBarcode string) ([]PackageBag, error) {
	rows, err := q.db.QueryContext(ctx, listPackagesInBag, bagBarcode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PackageBag
	for rows.Next() {
		var i PackageBag
		if err := rows.Scan(&i.BagBarcode, &i.PackageBarcode); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
