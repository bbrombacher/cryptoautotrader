package storage

import (
	"bbrombacher/cryptoautotrader/be/storage/models"
	"context"
)

func (s *StorageClient) StartStopTradeSession(ctx context.Context, input models.TradeSessionEntry) (*models.TradeSessionEntry, error) {
	// add logic to get starting and ending balance?
	entry, err := s.SqlClient.UpsertTradeSession(ctx, input)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *StorageClient) GetTradeSession(ctx context.Context, userID, sessionID string) (*models.TradeSessionEntry, error) {
	entry, err := s.SqlClient.SelectTradeSession(ctx, userID, sessionID)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *StorageClient) GetTradeSessions(ctx context.Context, params models.GetTradeSessionsParams) ([]models.TradeSessionEntry, error) {
	entries, err := s.SqlClient.SelectTradeSessions(ctx, params)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
