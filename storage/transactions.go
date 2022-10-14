package storage

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"errors"
	"fmt"
)

func (s *StorageClient) CreateTransaction(ctx context.Context, input models.TransactionEntry) (*models.TransactionEntry, error) {

	// get balance of purchase with currency and to purchase currency
	// validate purchase with currency has enough to continue
	// insert into transaction table
	// update balance of purchase with currency and purchased currency.

	currentBalance, err := s.GetBalance(ctx, input.UserID, input.CurrencyID)
	if err != nil {
		return nil, fmt.Errorf("error getting balance %w", err)
	}
	if currentBalance.Amount.LessThan(input.Price) {
		return nil, errors.New("insufficient funds for transaction")
	}

	entry, err := s.SqlClient.InsertTransaction(ctx, input)
	if err != nil {
		return nil, err
	}

	balanceInput := models.BalanceEntry{
		UserID:     input.UserID,
		CurrencyID: input.CurrencyID,
		Amount:     currentBalance.Amount.Sub(input.Price),
	}
	_, err = s.UpdateBalance(ctx, balanceInput)
	if err != nil {
		return nil, fmt.Errorf("error updating balance %w", err)
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
