package transforms

import "bbrombacher/cryptoautotrader/coinbase"

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
