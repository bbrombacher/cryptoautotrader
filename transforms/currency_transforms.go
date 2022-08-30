package transforms

import (
	currencyRequest "bbrombacher/cryptoautotrader/controllers/currencies/request"
	storageModels "bbrombacher/cryptoautotrader/storage/models"
)

func BuildCurrencyEntryFromPostRequest(request currencyRequest.PostCurrencyRequest) storageModels.CurrencyEntry {
	return storageModels.CurrencyEntry{
		Name:        request.Name,
		Description: request.Description,
	}
}

func BuildCurrencyEntryFromPatchRequest(request currencyRequest.PatchCurrencyRequest) storageModels.CurrencyEntry {
	return storageModels.CurrencyEntry{
		ID:          request.ID,
		Name:        safeDereferenceString(request.Name),
		Description: safeDereferenceString(request.Description),
	}
}

func BuildGetCurrenciesParamsFromGetCurrenciesRequest(request currencyRequest.GetCurrenciesRequest) storageModels.GetCurrenciesParams {
	return storageModels.GetCurrenciesParams{
		Cursor: request.Cursor,
		Limit:  request.Limit,
	}
}
