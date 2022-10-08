package tradebot

import (
	"bbrombacher/cryptoautotrader/coinbase"
	"bbrombacher/cryptoautotrader/storage"
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"fmt"
	"sync"

	"github.com/matryer/runner"
)

type Bot struct {
	Coinbase      coinbase.Client
	StorageClient *storage.StorageClient
	Tasks         *sync.Map
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

	task := runner.Go(func(shouldStop runner.S) error {
		err := b.startTrading(userID, tickerID)
		if err != nil {
			return err
		}
		return nil
	})

	b.Tasks.Store(tickerID, task)

	return tickerID, nil
}

func (b Bot) startTrading(userID, tickerID string) error {

	count := 0
	for {
		msg, err := b.Coinbase.GetTickerMessages(tickerID)
		if err != nil {
			return err
		}
		if count == 10 {
			b.makeTrade(userID, msg)
			count = 0
		}

		count++
	}
}

func (b Bot) makeTrade(userID string, msg map[string]interface{}) error {
	// TODO: Fix create transaction function.
	_, err := b.StorageClient.CreateTransaction(context.Background(), models.TransactionEntry{
		ID: "t",
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
