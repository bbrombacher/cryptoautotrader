package storage_test

import (
	mockstorage "bbrombacher/cryptoautotrader/mocks/storage"
	"bbrombacher/cryptoautotrader/storage"
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStorageClient_GetCurrency(t *testing.T) {
	// setup gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock sql calls
	sqlClient := mockstorage.NewMockSQLClient(ctrl)
	sqlEntry := models.CurrencyEntry{ID: "one", Name: "eth", Description: "eth coin"}
	sqlClient.EXPECT().SelectCurrency(gomock.Any(), "one").Return(&sqlEntry, nil)

	// make storage client and execute test
	storageClient := storage.NewStorageClient(sqlClient)
	result, err := storageClient.GetCurrency(context.Background(), "one")
	assert.Nil(t, err)
	assert.Equal(t, "eth", result.Name)
	assert.Equal(t, "eth coin", result.Description)
}
