package response

import "bbrombacher/cryptoautotrader/storage/models"

type GetCurrenciesResponse struct {
	Currencies []models.CurrencyEntry `json:"currencies"`
}

type GetCurrencyResponse struct {
	Currency *models.CurrencyEntry `json:"currency"`
}

type CreateCurrencyResponse struct {
	Currency *models.CurrencyEntry `json:"currency"`
}

type PatchCurrencyResponse struct {
	Currency *models.CurrencyEntry `json:"currency"`
}
