package transforms

import (
	tradeSessionsRequest "bbrombacher/cryptoautotrader/be/controllers/trade_sessions/request"
	storageModels "bbrombacher/cryptoautotrader/be/storage/models"
)

func BuildGetTradeSessionsFromGetRequest(req tradeSessionsRequest.GetTradeSessionsRequest) storageModels.GetTradeSessionsParams {
	return storageModels.GetTradeSessionsParams{
		UserID: req.UserID,
		Cursor: req.Cursor,
		Limit:  req.Limit,
	}
}
