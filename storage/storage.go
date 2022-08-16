package storage

import (
	"bbrombacher/cryptoautotrader/storage/sql"

	"github.com/google/uuid"
)

type StorageClient struct {
	SqlClient sql.SQLClient
}

func NewStorageClient(sqlClient sql.SQLClient) *StorageClient {
	return &StorageClient{SqlClient: sqlClient}
}

func generateUUID() string {
	return uuid.New().String()
}
