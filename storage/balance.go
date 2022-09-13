package storage

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
)

func (s *StorageClient) GetBalance(ctx context.Context, userID, currencyID string) (*models.BalanceEntry, error) {
	entry, err := s.SqlClient.SelectBalance(ctx, userID, currencyID)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *StorageClient) GetBulkBalance(ctx context.Context, userID string) ([]models.BalanceEntry, error) {
	entry, err := s.SqlClient.SelectBulkBalance(ctx, userID)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *StorageClient) UpdateBalance(ctx context.Context, input models.BalanceEntry) (*models.BalanceEntry, error) {
	entry, err := s.SqlClient.UpsertBalance(ctx, input)
	if err != nil {
		return nil, err
	}

	return entry, nil
}
