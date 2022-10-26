package response

import "bbrombacher/cryptoautotrader/be/storage/models"

type GetTransactionsResponse struct {
	Transactions []models.TransactionEntry `json:"transactions"`
}
