package db

import (
	"context"
	"database/sql"
	"sort"
	"testing"

	"github.com/DevelopmentHiring/FatihDurmaz/util"
	"github.com/stretchr/testify/require"
)

func TestCreateVehicle(t *testing.T) {
	testVehicle := createRandomVehicle(t)
	deleteVehicleWithPlate(t, testVehicle)
}

func TestGetVehicle(t *testing.T) {
	vehicle := createRandomVehicle(t)
	testVehicle, err := testQueries.GetVehicle(context.Background(), vehicle)
	require.NoError(t, err)
	require.NotEmpty(t, testVehicle)
	require.Equal(t, vehicle, testVehicle)
	deleteVehicleWithPlate(t, testVehicle)
}

func TestListVehicles(t *testing.T) {
	vehicles := createListVehicles(t)
	testVehicles, err := testQueries.ListVehicles(context.Background())
	require.NoError(t, err)
	sort.Strings(vehicles)
	sort.Strings(testVehicles)
	require.NotEmpty(t, testVehicles)
	require.Equal(t, vehicles, testVehicles)
	for i := 0; i < 10; i++ {
		deleteVehicleWithPlate(t, testVehicles[i])
	}
}
func createListVehicles(t *testing.T) []string {
	var vehicles []string
	for i := 0; i < 10; i++ {
		testVehicle := createRandomVehicle(t)
		vehicles = append(vehicles, testVehicle)
	}
	return vehicles
}
func createRandomVehicle(t *testing.T) string {
	arg := util.RandomName()
	testVehicle, err := testQueries.CreateVehicle(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg, testVehicle)
	return testVehicle
}

func deleteVehicleWithPlate(t *testing.T, plate string) {
	err := testQueries.DeleteVehicle(context.Background(), plate)
	require.NoError(t, err)

	_, err = testQueries.GetVehicle(context.Background(), plate)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
