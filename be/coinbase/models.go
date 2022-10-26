package coinbase

import (
	"time"

	"github.com/shopspring/decimal"
)

type StartTickerParams struct {
	Type       string    `json:"type"`
	ProductIDs []string  `json:"product_ids"`
	Channels   []Channel `json:"channels"`
}

type Channel struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}

type TickerMsg struct {
	BestAsk     string          `json:"best_ask"`
	BestAskSize string          `json:"best_ask_size"`
	BestBid     string          `json:"best_bid"`
	BestBidSize string          `json:"best_bid_size"`
	High24h     string          `json:"high_24h"`
	LastSize    string          `json:"last_size"`
	Low24h      string          `json:"low_24h"`
	Open24h     string          `json:"open_24h"`
	Price       string          `json:"price"`
	ProductID   string          `json:"product_id"`
	Sequence    decimal.Decimal `json:"sequence"`
	Side        string          `json:"side"`
	Time        time.Time       `json:"time"`
	TradeID     decimal.Decimal `json:"trade_id"`
	Type        string          `json:"type"`
	Volumn24h   string          `json:"volume_24h"`
	Volume30d   string          `json:"volume_30d"`
}

func (t TickerMsg) GetPriceAsDecimal() (decimal.Decimal, error) {
	price, err := decimal.NewFromString(t.Price)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return price, nil
}
