package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddPackageBag(t *testing.T) {
	bag := createRandomBag(t)
	pk := createRandomPackageWithDelivery(t, bag.DeliveryID)

	arg := AddPackageBagParams{
		BagBarcode:     bag.Barcode,
		PackageBarcode: pk.Barcode,
	}

	testPackageBag, err := testQueries.AddPackageBag(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, testPackageBag)
	require.Equal(t, arg.BagBarcode, testPackageBag.BagBarcode)
	require.Equal(t, arg.PackageBarcode, testPackageBag.PackageBarcode)

	testDeletePackageBag(t, testPackageBag.PackageBarcode)
	deletePackageWithBarcode(t, arg.PackageBarcode)
	deleteBagWithBarcode(t, bag.Barcode)
	deleteDeliveryPointWithId(t, bag.DeliveryID)
}

func TestAddPackageVehicle(t *testing.T) {
	pk := createRandomPackage(t)
	vehicle := createRandomVehicle(t)
	arg := AddPackageVehicleParams{
		PackageBarcode: pk.Barcode,
		VehiclePlate:   vehicle,
	}
	testPackageVehicle, err := testQueries.AddPackageVehicle(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, testPackageVehicle)
	require.Equal(t, arg.PackageBarcode, testPackageVehicle.PackageBarcode)
	require.Equal(t, arg.VehiclePlate, testPackageVehicle.VehiclePlate)

	testDeletePackageVehicle(t, testPackageVehicle.PackageBarcode)
	deletePackageWithBarcode(t, arg.PackageBarcode)
	deleteVehicleWithPlate(t, vehicle)
	deleteDeliveryPointWithId(t, pk.DeliveryID)
}

func TestAddBagVehicle(t *testing.T) {
	bag := createRandomBag(t)
	vehicle := createRandomVehicle(t)
	arg := AddBagVehicleParams{
		BagBarcode:   bag.Barcode,
		VehiclePlate: vehicle,
	}
	testBagVehicle, err := testQueries.AddBagVehicle(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, testBagVehicle)
	require.Equal(t, arg.BagBarcode, testBagVehicle.BagBarcode)
	require.Equal(t, arg.VehiclePlate, testBagVehicle.VehiclePlate)

	testDeleteBagVehicle(t, testBagVehicle.BagBarcode)
	deleteBagWithBarcode(t, arg.BagBarcode)
	deleteVehicleWithPlate(t, vehicle)
	deleteDeliveryPointWithId(t, bag.DeliveryID)
}

func testDeletePackageBag(t *testing.T, packageBarcode string) {
	err := testQueries.DeletePackageBag(context.Background(), packageBarcode)
	require.NoError(t, err)

	testBag, err := testQueries.GetPackageBag(context.Background(), packageBarcode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, testBag)
}
func testDeleteBagVehicle(t *testing.T, bagBarcode string) {
	err := testQueries.DeleteBagVehicle(context.Background(), bagBarcode)
	require.NoError(t, err)

	testBag, err := testQueries.GetBagVehicle(context.Background(), bagBarcode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, testBag)
}
func testDeletePackageVehicle(t *testing.T, packageBarcode string) {
	err := testQueries.DeletePackageVehicle(context.Background(), packageBarcode)
	require.NoError(t, err)

	testBag, err := testQueries.GetPackageVehicle(context.Background(), packageBarcode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, testBag)
}
