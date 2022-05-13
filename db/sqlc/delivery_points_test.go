package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DevelopmentHiring/FatihDurmaz/util"
	"github.com/stretchr/testify/require"
)

func TestCreateDeliveryPoint(t *testing.T) {
	point := createRandomDeliveryPoint(t)
	deleteDeliveryPointWithId(t, point.ID)

}

func TestGetDeliveryPoint(t *testing.T) {
	point := createRandomDeliveryPoint(t)

	testPoint, err := testQueries.GetDeliveryPoint(context.Background(), point.ID)
	require.NoError(t, err)
	require.NotEmpty(t, testPoint)
	require.Equal(t, point.ID, testPoint.ID)
	require.Equal(t, point.Name, testPoint.Name)

	deleteDeliveryPointWithId(t, point.ID)
}

func TestDeleteDeliveryPoint(t *testing.T) {
	point := createRandomDeliveryPoint(t)
	deleteDeliveryPointWithId(t, point.ID)
}

func deleteDeliveryPointWithId(t *testing.T, id int32) {
	err := testQueries.DeleteDeliveryPoint(context.Background(), id)
	require.NoError(t, err)

	testPoint, err := testQueries.GetDeliveryPoint(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, testPoint)
}

func createRandomDeliveryPoint(t *testing.T) DeliveryPoint {
	arg := CreateDeliveryPointParams{
		ID:   util.RandomInt(1, 4),
		Name: util.RandomName(),
	}
	testPoint, err := testQueries.CreateDeliveryPoint(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, testPoint)
	require.Equal(t, arg.ID, testPoint.ID)
	require.Equal(t, arg.Name, testPoint.Name)

	return testPoint
}
