package storage

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
)

func (s *StorageClient) GetCurrency(ctx context.Context, id string) (*models.CurrencyEntry, error) {
	entry, err := s.SqlClient.SelectCurrency(ctx, id)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *StorageClient) GetCurrencies(ctx context.Context, params models.GetCurrenciesParams) ([]models.CurrencyEntry, error) {
	entry, err := s.SqlClient.SelectCurrencies(ctx, params)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (s *StorageClient) CreateCurrency(ctx context.Context, entry models.CurrencyEntry) (*models.CurrencyEntry, error) {
	if entry.ID == "" {
		entry.ID = generateUUID()
	}

	result, err := s.SqlClient.InsertCurrency(ctx, entry)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *StorageClient) UpdateCurrency(ctx context.Context, entry models.CurrencyEntry, updateColumns []string) (*models.CurrencyEntry, error) {
	result, err := s.SqlClient.UpdateCurrency(ctx, entry, updateColumns)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *StorageClient) DeleteCurrency(ctx context.Context, id string) error {
	err := s.SqlClient.DeleteCurrency(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
