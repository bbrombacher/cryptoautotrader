package response

import "bbrombacher/cryptoautotrader/be/storage/models"

type GetTradeSessionsResponse struct {
	TradeSessions []models.TradeSessionEntry `json:"trade_sessions"`
}
