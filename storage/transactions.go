package storage

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
)

func (s *StorageClient) CreateTransaction(ctx context.Context, input models.TransactionEntry) (*models.TransactionEntry, error) {
	entry, err := s.SqlClient.InsertTransaction(ctx, input)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *StorageClient) GetTransaction(ctx context.Context, transactionID, userID string) (*models.TransactionEntry, error) {
	entry, err := s.SqlClient.SelectTransaction(ctx, transactionID, userID)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *StorageClient) GetTransactions(ctx context.Context, params models.GetTransactionsParams) ([]models.TransactionEntry, error) {
	entries, err := s.SqlClient.SelectTransactions(ctx, params)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
