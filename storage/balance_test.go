package storage_test

import (
	mockstorage "bbrombacher/cryptoautotrader/mocks/storage"
	"bbrombacher/cryptoautotrader/storage"
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_GetBulkBalance(t *testing.T) {
	// setup gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock sql calls
	sqlClient := mockstorage.NewMockSQLClient(ctrl)

	amnt := decimal.NewFromFloat32(1.25)
	expectEntries := []models.BalanceEntry{
		{
			UserID: "U", CurrencyID: "C", Amount: amnt,
		},
	}
	sqlClient.EXPECT().SelectBulkBalance(gomock.Any(), "U").Return(expectEntries, nil)

	// make storage client and execute test
	storageClient := storage.NewStorageClient(sqlClient)
	result, err := storageClient.GetBulkBalance(context.Background(), "U")
	assert.Nil(t, err)
	assert.Equal(t, "U", result[0].UserID)
	assert.Equal(t, "C", result[0].CurrencyID)
}

func Test_GetBalance(t *testing.T) {
	// setup gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock sql calls
	sqlClient := mockstorage.NewMockSQLClient(ctrl)

	amnt := decimal.NewFromFloat32(1.25)
	expectEntries := models.BalanceEntry{UserID: "U", CurrencyID: "C", Amount: amnt}
	sqlClient.EXPECT().SelectBalance(gomock.Any(), "U", "C").Return(&expectEntries, nil)

	// make storage client and execute test
	storageClient := storage.NewStorageClient(sqlClient)
	result, err := storageClient.GetBalance(context.Background(), "U", "C")
	assert.Nil(t, err)
	assert.Equal(t, "U", result.UserID)
	assert.Equal(t, "C", result.CurrencyID)
}

func Test_UpdateBalance(t *testing.T) {
	// setup gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock sql calls
	sqlClient := mockstorage.NewMockSQLClient(ctrl)

	amnt := decimal.NewFromFloat32(1.25)
	entry := models.BalanceEntry{UserID: "U", CurrencyID: "C", Amount: amnt}
	sqlClient.EXPECT().UpsertBalance(gomock.Any(), entry).Return(&entry, nil)

	// make storage client and execute test
	storageClient := storage.NewStorageClient(sqlClient)
	result, err := storageClient.UpdateBalance(context.Background(), entry)
	assert.Nil(t, err)
	assert.Equal(t, "U", result.UserID)
	assert.Equal(t, "C", result.CurrencyID)
}
