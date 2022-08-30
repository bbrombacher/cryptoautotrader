package response

import "bbrombacher/cryptoautotrader/storage/models"

type GetCurrencyResponse struct {
	Currency *models.CurrencyEntry `json:"currency,omitempty"`
}

type CreateCurrencyResponse struct {
	Currency *models.CurrencyEntry `json:"currency,omitempty"`
}

type PatchCurrencyResponse struct {
	Currency *models.CurrencyEntry `json:"currency,omitempty"`
}
