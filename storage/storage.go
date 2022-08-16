package storage

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"bbrombacher/cryptoautotrader/storage/sql"
	"context"

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

func (s *StorageClient) GetUser(ctx context.Context, id string) (*models.UserEntry, error) {
	entry, err := s.SqlClient.SelectUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *StorageClient) CreateUser(ctx context.Context, entry models.UserEntry) (*models.UserEntry, error) {
	if entry.ID == "" {
		entry.ID = generateUUID()
	}

	result, err := s.SqlClient.InsertUser(ctx, entry)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *StorageClient) UpdateUser(ctx context.Context, entry models.UserEntry, updateColumns []string) (*models.UserEntry, error) {
	result, err := s.SqlClient.UpdateUser(ctx, entry, updateColumns)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *StorageClient) DeleteUser(ctx context.Context, id string) error {
	err := s.SqlClient.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
