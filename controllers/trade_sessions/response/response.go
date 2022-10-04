package response

import "bbrombacher/cryptoautotrader/storage/models"

type GetTradeSessionsResponse struct {
	TradeSessions []models.TradeSessionEntry `json:"trade_sessions"`
}
