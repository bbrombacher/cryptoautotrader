package transforms

import (
	tradeSessionsRequest "bbrombacher/cryptoautotrader/controllers/trade_sessions/request"
	storageModels "bbrombacher/cryptoautotrader/storage/models"
)

func BuildGetTradeSessionsFromGetRequest(req tradeSessionsRequest.GetTradeSessionsRequest) storageModels.GetTradeSessionsParams {
	return storageModels.GetTradeSessionsParams{
		UserID: req.UserID,
		Cursor: req.Cursor,
		Limit:  req.Limit,
	}
}
