package transforms

import (
	balanceRequest "bbrombacher/cryptoautotrader/be/controllers/balance/request"
	balanceResponse "bbrombacher/cryptoautotrader/be/controllers/balance/response"
	"bbrombacher/cryptoautotrader/be/storage/models"

	"github.com/shopspring/decimal"
)

func BuildUpdateBalanceEntryFromUpdateBalanceRequest(req balanceRequest.UpdateBalanceRequest) models.BalanceEntry {
	// string is validated earlier in the request
	amount, _ := decimal.NewFromString(req.Amount)

	return models.BalanceEntry{
		UserID:     req.UserID,
		CurrencyID: req.CurrencyID,
		Amount:     amount,
	}
}

func BuildBalanceResponseObjectFromBalanceEntry(entry *models.BalanceEntry) balanceResponse.BalanceResponse {
	return balanceResponse.BalanceResponse{
		Balance: balanceResponse.Balance{
			Amount:   entry.Amount.String(),
			Currency: entry.Currency,
		},
	}
}

func BuildBulkBalanceResponseObjectFromBalanceEntry(entries []models.BalanceEntry) balanceResponse.BulkBalanceResponse {
	balances := make([]balanceResponse.Balance, 0, len(entries))
	for _, entry := range entries {
		balance := balanceResponse.Balance{
			Amount:   entry.Amount.String(),
			Currency: entry.Currency,
		}
		balances = append(balances, balance)
	}

	return balanceResponse.BulkBalanceResponse{
		Balance: balances,
	}
}
