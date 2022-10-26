package storage_test

import (
	mockstorage "bbrombacher/cryptoautotrader/be/mocks/storage"
	"bbrombacher/cryptoautotrader/be/storage"
	"bbrombacher/cryptoautotrader/be/storage/models"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStorageClient_GetUser(t *testing.T) {
	// setup gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock sql calls
	sqlClient := mockstorage.NewMockSQLClient(ctrl)
	sqlEntry := models.UserEntry{ID: "one", FirstName: "brandon", LastName: "brombacher"}
	sqlClient.EXPECT().SelectUser(gomock.Any(), "one").Return(&sqlEntry, nil)

	// make storage client and execute test
	storageClient := storage.NewStorageClient(sqlClient)
	result, err := storageClient.GetUser(context.Background(), "one")
	assert.Nil(t, err)
	assert.Equal(t, "brandon", result.FirstName)
	assert.Equal(t, "brombacher", result.LastName)
}
