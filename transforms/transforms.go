package transforms

import "bbrombacher/cryptoautotrader/coinbase"

func MakeStartTickerParams(productIDs []string) coinbase.StartTickerParams {
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
