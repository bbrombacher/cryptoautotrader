package tradebot

import (
	"bbrombacher/cryptoautotrader/coinbase"
	"bbrombacher/cryptoautotrader/storage"
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
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

	currencyOne, err := currencies.GetCurrencyIDByName("usd")
	if err != nil {
		log.Println("get currency by name err", err)
	}

	currencyTwo, err := currencies.GetCurrencyIDByName("eth")
	if err != nil {
		log.Println("get currency by name err", err)
	}

	task := runner.Go(func(shouldStop runner.S) error {
		err := b.startTrading(
			userID,
			tickerID,
			currencyOne,
			currencyTwo,
		)
		if err != nil {
			log.Println("error:", err)
			return err
		}
		return nil
	})

	b.Tasks.Store(tickerID, task)

	return tickerID, nil
}

func (b Bot) startTrading(userID, tickerID, currencyOne, currencyTwo string) error {

	buyCount := 0
	sellCount := 0
	for {
		msg, err := b.Coinbase.GetTickerMessages(tickerID)
		if err != nil {
			return err
		}
		price, _ := msg.GetPriceAsDecimal()

		if buyCount == 10 {
			err = b.makeTrade(TransactionParams{
				UserID:         userID,
				UseCurrencyID:  currencyOne,
				GetCurrencyID:  currencyTwo,
				TradeSessionID: tickerID,
				Type:           "BUY",
				Amount:         decimal.NewFromFloat32(1.25),
				Price:          price,
			})
			if err != nil {
				b.StopTrading(userID, tickerID)
			}
			buyCount = 0
			continue
		}

		if sellCount == 15 {
			err = b.makeTrade(TransactionParams{
				UserID:         userID,
				UseCurrencyID:  currencyTwo,
				GetCurrencyID:  currencyOne,
				TradeSessionID: tickerID,
				Type:           "SELL",
				Amount:         decimal.NewFromFloat32(1.25),
				Price:          price,
			})
			if err != nil {
				b.StopTrading(userID, tickerID)
			}
			sellCount = 0
			continue
		}

		sellCount++
		buyCount++
	}
}

type TransactionParams struct {
	UserID         string
	UseCurrencyID  string
	GetCurrencyID  string
	TradeSessionID string
	Type           string
	Amount         decimal.Decimal
	Price          decimal.Decimal
}

func (b Bot) makeTrade(params TransactionParams) error {
	id := uuid.New()
	_, err := b.StorageClient.CreateTransaction(
		context.Background(),
		params.Type,
		models.TransactionEntry{
			ID:              id.String(),
			UserID:          params.UserID,
			UseCurrencyID:   params.UseCurrencyID,
			GetCurrencyID:   params.GetCurrencyID,
			TransactionType: params.Type,
			Amount:          params.Amount,
			Price:           params.Price,
		},
	)
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