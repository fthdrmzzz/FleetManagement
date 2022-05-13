package db

import (
	"context"
	"database/sql"
	"sort"
	"testing"

	"github.com/DevelopmentHiring/FatihDurmaz/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePackage(t *testing.T) {
	testPackage := createRandomPackage(t)
	deletePackageWithBarcode(t, testPackage.Barcode)
	deleteDeliveryPointWithId(t, testPackage.DeliveryID)
}
func TestGetPackage(t *testing.T) {
	pk := createRandomPackage(t)

	testPackage, err := testQueries.GetPackage(context.Background(), pk.Barcode)
	require.NoError(t, err)
	require.NotEmpty(t, testPackage)
	require.Equal(t, pk.DeliveryID, testPackage.DeliveryID)
	require.Equal(t, pk.Barcode, testPackage.Barcode)
	require.Equal(t, pk.PackageStatus, testPackage.PackageStatus)
	require.Equal(t, pk.PackageWeight, testPackage.PackageWeight)

	deletePackageWithBarcode(t, testPackage.Barcode)
	deleteDeliveryPointWithId(t, testPackage.DeliveryID)
}
func TestDeletePackage(t *testing.T) {
	testPackage := createRandomPackage(t)
	deletePackageWithBarcode(t, testPackage.Barcode)
	deleteDeliveryPointWithId(t, testPackage.DeliveryID)
}

func TestListPackages(t *testing.T) {
	packages := createListPackages(t)
	testPackages, err := testQueries.ListPackages(context.Background())
	require.NoError(t, err)
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].Barcode < packages[j].Barcode
	})
	sort.Slice(testPackages, func(i, j int) bool {
		return testPackages[i].Barcode < testPackages[j].Barcode
	})
	require.NotEmpty(t, testPackages)
	require.Equal(t, packages, testPackages)
	for i := 0; i < 10; i++ {
		deletePackageWithBarcode(t, testPackages[i].Barcode)
	}
	deleteDeliveryPointWithId(t, packages[0].DeliveryID)
}

func TestUpdatePackage(t *testing.T) {
	pk := createRandomPackage(t)
	arg := UpdatePackageParams{
		Barcode:       pk.Barcode,
		PackageStatus: "unloaded",
	}
	testPackage, err := testQueries.UpdatePackage(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, testPackage.PackageStatus, arg.PackageStatus)
	deletePackageWithBarcode(t, pk.Barcode)
	deleteDeliveryPointWithId(t, pk.DeliveryID)
}
func createRandomPackage(t *testing.T) Package {
	deliveryPoint := createRandomDeliveryPoint(t)
	return createRandomPackageWithDelivery(t, deliveryPoint.ID)
}

func createRandomPackageWithDelivery(t *testing.T, deliveryId int32) Package {
	arg := CreatePackageParams{
		Barcode:       util.RandomBarcode(),
		PackageStatus: PackagesStatusCreated,
		PackageWeight: util.RandomWeight(),
		DeliveryID:    deliveryId,
	}
	testPackage, err := testQueries.CreatePackage(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.DeliveryID, testPackage.DeliveryID)
	require.Equal(t, arg.Barcode, testPackage.Barcode)
	require.Equal(t, arg.PackageStatus, testPackage.PackageStatus)
	require.Equal(t, arg.PackageWeight, testPackage.PackageWeight)

	return testPackage
}

func deletePackageWithBarcode(t *testing.T, barcode string) {
	err := testQueries.DeletePackage(context.Background(), barcode)
	require.NoError(t, err)

	testPackage, err := testQueries.GetPackage(context.Background(), barcode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, testPackage)
}

func createListPackages(t *testing.T) []Package {
	deliveryPoint := createRandomDeliveryPoint(t)
	var packages []Package
	for i := 0; i < 10; i++ {
		testPackage := createRandomPackageWithDelivery(t, deliveryPoint.ID)
		packages = append(packages, testPackage)
	}
	return packages
}
