package storage

import (
	"bbrombacher/cryptoautotrader/be/storage/sql"

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
