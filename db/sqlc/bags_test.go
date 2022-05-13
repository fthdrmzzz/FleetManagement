package db

import (
	"context"
	"database/sql"
	"sort"
	"testing"

	"github.com/DevelopmentHiring/FatihDurmaz/util"
	"github.com/stretchr/testify/require"
)

func TestCreateBag(t *testing.T) {
	bag := createRandomBag(t)
	deliveryPointId := bag.DeliveryID
	deleteBagWithBarcode(t, bag.Barcode)
	deleteDeliveryPointWithId(t, deliveryPointId)
}

func TestGetBag(t *testing.T) {
	bag := createRandomBag(t)

	testBag, err := testQueries.GetBag(context.Background(), bag.Barcode)
	require.NoError(t, err)
	require.NotEmpty(t, testBag)
	require.Equal(t, bag.Barcode, testBag.Barcode)
	require.Equal(t, bag.BagStatus, testBag.BagStatus)
	require.Equal(t, bag.DeliveryID, testBag.DeliveryID)

	deliveryPointId := bag.DeliveryID
	deleteBagWithBarcode(t, bag.Barcode)
	deleteDeliveryPointWithId(t, deliveryPointId)
}

func TestDeleteBag(t *testing.T) {
	bag := createRandomBag(t)
	deliveryPointId := bag.DeliveryID
	deleteBagWithBarcode(t, bag.Barcode)
	deleteDeliveryPointWithId(t, deliveryPointId)
}

func TestListBags(t *testing.T) {
	bags := createListBag(t)
	testBags, err := testQueries.ListBags(context.Background())
	require.NoError(t, err)
	sort.Slice(testBags, func(i, j int) bool {
		return testBags[i].Barcode < testBags[j].Barcode
	})
	sort.Slice(bags, func(i, j int) bool {
		return bags[i].Barcode < bags[j].Barcode
	})
	require.NotEmpty(t, testBags)
	require.Equal(t, bags, testBags)
	for i := 0; i < 10; i++ {
		deleteBagWithBarcode(t, testBags[i].Barcode)
	}
	deleteDeliveryPointWithId(t, bags[0].DeliveryID)
}

func TestUpdateBag(t *testing.T) {
	bag := createRandomBag(t)
	arg := UpdateBagParams{
		Barcode:   bag.Barcode,
		BagStatus: "unloaded",
	}
	testBag, err := testQueries.UpdateBag(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, testBag.BagStatus, arg.BagStatus)
	deleteBagWithBarcode(t, bag.Barcode)
	deleteDeliveryPointWithId(t, bag.DeliveryID)
}

func deleteBagWithBarcode(t *testing.T, barcode string) {
	err := testQueries.DeleteBag(context.Background(), barcode)
	require.NoError(t, err)

	testBag, err := testQueries.GetBag(context.Background(), barcode)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, testBag)
}

func createRandomBag(t *testing.T) Bag {
	deliveryPoint := createRandomDeliveryPoint(t)
	return createRandomBagWithDelivery(t, deliveryPoint.ID)
}

func createRandomBagWithDelivery(t *testing.T, deliveryID int32) Bag {
	arg := CreateBagParams{
		Barcode:    util.RandomBarcode(),
		BagStatus:  BagsStatusCreated,
		DeliveryID: deliveryID,
	}
	testBag, err := testQueries.CreateBag(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, testBag)
	require.Equal(t, arg.Barcode, testBag.Barcode)
	require.Equal(t, arg.BagStatus, testBag.BagStatus)
	require.Equal(t, arg.DeliveryID, testBag.DeliveryID)
	return testBag
}
func createListBag(t *testing.T) []Bag {
	deliveryPoint := createRandomDeliveryPoint(t)
	var bags []Bag
	for i := 0; i < 10; i++ {
		testBag := createRandomBagWithDelivery(t, deliveryPoint.ID)
		bags = append(bags, testBag)
	}
	return bags
}
