package tradebot

import (
	"bbrombacher/cryptoautotrader/coinbase"
	"bbrombacher/cryptoautotrader/storage"
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/matryer/runner"
	"github.com/shopspring/decimal"
)

type Bot struct {
	Coinbase      coinbase.Client
	StorageClient *storage.StorageClient
	Tasks         *sync.Map
}

type currencies struct {
	UseCurrency string
	GetCurrency string
}

func (b Bot) StartTrading(userID string, duration int) (string, error) {

	// start ticker feed
	tickerID, err := b.Coinbase.StartTickerFeed(coinbase.StartTickerParams{
		Type:       "subscribe",
		ProductIDs: []string{"ETH-USD"},
		Channels: []coinbase.Channel{
			{
				Name:       "ticker",
				ProductIDs: []string{"ETH-USD"},
			},
		},
	})
	if err != nil {
		return "", err
	}

	var currencies models.Currencies
	currencies, err = b.StorageClient.GetCurrencies(context.Background(), models.GetCurrenciesParams{
		Cursor: 0,
		Limit:  100,
	})
	if err != nil {
		return "", err
	}

	useCurrencyID, err := currencies.GetCurrencyIDByName("usd")
	if err != nil {
		log.Println("get currency by name err", err)
	}

	getCurrencyID, err := currencies.GetCurrencyIDByName("eth")
	if err != nil {
		log.Println("get currency by name err", err)
	}

	task := runner.Go(func(shouldStop runner.S) error {
		err := b.startTrading(
			userID,
			tickerID,
			useCurrencyID,
			getCurrencyID,
		)
		if err != nil {
			return err
		}
		return nil
	})

	b.Tasks.Store(tickerID, task)

	return tickerID, nil
}

func (b Bot) startTrading(userID, tickerID, useCurrencyID, getCurrencyID string) error {

	count := 0
	for {
		msg, err := b.Coinbase.GetTickerMessages(tickerID)
		if err != nil {
			return err
		}
		if count == 10 {
			err = b.makeTrade(userID, useCurrencyID, getCurrencyID, msg)
			if err != nil {
				// if insufficient funds, we should stop bot.
			}
			count = 0
			continue
		}

		count++
	}
}

func (b Bot) makeTrade(userID, useCurrencyID, getCurrencyID string, msg map[string]interface{}) error {
	p, ok := msg["price"].(string)
	if !ok {
		return errors.New("unable to parse price")
	}

	price, err := decimal.NewFromString(p)
	if err != nil {
		return err
	}

	id := uuid.New()
	_, err = b.StorageClient.CreateTransaction(context.Background(), models.TransactionEntry{
		ID:              id.String(),
		UserID:          userID,
		UseCurrencyID:   useCurrencyID,
		GetCurrencyID:   getCurrencyID,
		TransactionType: "BUY",
		Amount:          decimal.NewFromInt(1),
		Price:           price,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b Bot) StopTrading(userID string, tickerID string) error {
	t, ok := b.Tasks.Load(tickerID)
	if !ok {
		return fmt.Errorf("trade session not found %v", tickerID)
	}

	task, ok := t.(*runner.Task)
	if !ok {
		return fmt.Errorf("could not cast task %v", task)
	}

	task.Stop()

	err := b.Coinbase.CloseTickerFeed(tickerID)
	if err != nil {
		return err
	}

	return nil
}
