package transforms

import (
	"bbrombacher/cryptoautotrader/coinbase"
	userRequest "bbrombacher/cryptoautotrader/controllers/users/request"
	storageModels "bbrombacher/cryptoautotrader/storage/models"
)

func BuildStartTickerParams(productIDs []string) coinbase.StartTickerParams {
	return coinbase.StartTickerParams{
		Type:       "subscribe",
		ProductIDs: productIDs,
		Channels: []coinbase.Channel{
			{
				Name:       "ticker",
				ProductIDs: productIDs,
			},
		},
	}
}

func BuildUserEntryFromPostRequest(request userRequest.PostUserRequest) storageModels.UserEntry {
	return storageModels.UserEntry{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
}
