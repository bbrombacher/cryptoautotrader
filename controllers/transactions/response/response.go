package response

import "bbrombacher/cryptoautotrader/storage/models"

type GetTransactionsResponse struct {
	Transactions []models.TransactionEntry `json:"transactions"`
}
