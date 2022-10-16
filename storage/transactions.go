package storage

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"errors"
	"fmt"
)

// buy crypto with usd
func (s *StorageClient) buy(ctx context.Context, tradeSessionID string, input models.TransactionEntry) (*models.TransactionEntry, error) {

	cost := input.Price.Mul(input.Amount)

	useCurrencyBalance, err := s.GetBalance(ctx, input.UserID, input.UseCurrencyID)
	if err != nil {
		return nil, fmt.Errorf("error getting balance for use currency %w", err)
	}
	getCurrencyBalance, err := s.GetBalance(ctx, input.UserID, input.GetCurrencyID)
	if err != nil {
		return nil, fmt.Errorf("error getting balance for get currency %w", err)
	}
	if useCurrencyBalance.Amount.LessThan(cost) {
		return nil, errors.New("insufficient funds for transaction")
	}

	entry, err := s.SqlClient.InsertTransaction(ctx, input)
	if err != nil {
		return nil, err
	}

	_, err = s.SqlClient.InsertTransactionSessionMapEntry(ctx, models.TransactionSessionMapEntry{
		TradeSessionID: tradeSessionID,
		TransactionID:  input.ID,
	})
	if err != nil {
		return nil, err
	}

	useBalanceUpdate := models.BalanceEntry{
		UserID:     input.UserID,
		CurrencyID: input.UseCurrencyID,
		Amount:     useCurrencyBalance.Amount.Sub(cost),
	}
	_, err = s.UpdateBalance(ctx, useBalanceUpdate)
	if err != nil {
		return nil, fmt.Errorf("error updating balance %w", err)
	}

	getBalanceUpdate := models.BalanceEntry{
		UserID:     input.UserID,
		CurrencyID: input.GetCurrencyID,
		Amount:     getCurrencyBalance.Amount.Add(input.Amount),
	}
	_, err = s.UpdateBalance(ctx, getBalanceUpdate)
	if err != nil {
		return nil, fmt.Errorf("error updating balance %w", err)
	}

	return entry, nil
}

// sell crypto with usd
func (s *StorageClient) sell(ctx context.Context, tradeSessionID string, input models.TransactionEntry) (*models.TransactionEntry, error) {

	useCurrencyBalance, err := s.GetBalance(ctx, input.UserID, input.UseCurrencyID)
	if err != nil {
		return nil, fmt.Errorf("error getting balance for use currency %w", err)
	}
	getCurrencyBalance, err := s.GetBalance(ctx, input.UserID, input.GetCurrencyID)
	if err != nil {
		return nil, fmt.Errorf("error getting balance for get currency %w", err)
	}

	// do we have enough crypto to sell?
	if useCurrencyBalance.Amount.LessThan(input.Amount) {
		return nil, errors.New("insufficient funds for transaction")
	}

	entry, err := s.SqlClient.InsertTransaction(ctx, input)
	if err != nil {
		return nil, err
	}

	_, err = s.SqlClient.InsertTransactionSessionMapEntry(ctx, models.TransactionSessionMapEntry{
		TradeSessionID: tradeSessionID,
		TransactionID:  input.ID,
	})
	if err != nil {
		return nil, err
	}

	useBalanceUpdate := models.BalanceEntry{
		UserID:     input.UserID,
		CurrencyID: input.UseCurrencyID,
		Amount:     useCurrencyBalance.Amount.Sub(input.Amount),
	}
	_, err = s.UpdateBalance(ctx, useBalanceUpdate)
	if err != nil {
		return nil, fmt.Errorf("error updating balance %w", err)
	}

	received := input.Price.Mul(input.Amount)
	getBalanceUpdate := models.BalanceEntry{
		UserID:     input.UserID,
		CurrencyID: input.GetCurrencyID,
		Amount:     getCurrencyBalance.Amount.Add(received),
	}
	_, err = s.UpdateBalance(ctx, getBalanceUpdate)
	if err != nil {
		return nil, fmt.Errorf("error updating balance %w", err)
	}

	return entry, nil
}

func (s *StorageClient) CreateTransaction(ctx context.Context, tradeSessionID string, input models.TransactionEntry) (*models.TransactionEntry, error) {
	var entry *models.TransactionEntry
	var err error

	switch input.TransactionType {
	case "BUY":
		entry, err = s.buy(ctx, tradeSessionID, input)
	case "SELL":
		entry, err = s.sell(ctx, tradeSessionID, input)
	default:
		return nil, fmt.Errorf("invalid transaction type %v", input.TransactionType)
	}
	return entry, err
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
