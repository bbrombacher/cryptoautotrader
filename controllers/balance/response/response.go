package response

import "bbrombacher/cryptoautotrader/storage/models"

type BulkBalanceResponse struct {
	Balance []Balance `json:"balance"`
}

type BalanceResponse struct {
	Balance Balance `json:"balance"`
}

type Balance struct {
	Amount   string               `json:"amount"`
	Currency models.CurrencyEntry `json:"currency"`
}
