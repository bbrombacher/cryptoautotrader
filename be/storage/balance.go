package storage

import (
	"bbrombacher/cryptoautotrader/be/storage/models"
	"context"
)

func (s *StorageClient) GetBalance(ctx context.Context, userID, currencyID string) (*models.BalanceEntry, error) {
	balanceEntry, err := s.SqlClient.SelectBalance(ctx, userID, currencyID)
	if err != nil {
		return nil, err
	}

	currencyEntry, err := s.SqlClient.SelectCurrency(ctx, currencyID)
	if err != nil {
		return nil, err
	}

	balanceEntry.Currency = *currencyEntry

	return balanceEntry, nil
}

func (s *StorageClient) GetBulkBalance(ctx context.Context, userID string) ([]models.BalanceEntry, error) {
	entries, err := s.SqlClient.SelectBulkBalance(ctx, userID)
	if err != nil {
		return nil, err
	}

	params := models.GetCurrenciesParams{Cursor: 0, Limit: 10000}
	currencies, err := s.SqlClient.SelectCurrencies(ctx, params)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(entries); i++ {
		for j := 0; j < len(currencies); j++ {
			if entries[i].CurrencyID == currencies[j].ID {
				entries[i].Currency = currencies[j]
			}
		}
	}

	return entries, nil
}

func (s *StorageClient) UpdateBalance(ctx context.Context, input models.BalanceEntry) (*models.BalanceEntry, error) {
	balanceEntry, err := s.SqlClient.UpsertBalance(ctx, input)
	if err != nil {
		return nil, err
	}

	currencyEntry, err := s.SqlClient.SelectCurrency(ctx, balanceEntry.CurrencyID)
	if err != nil {
		return nil, err
	}

	balanceEntry.Currency = *currencyEntry

	return balanceEntry, nil
}
