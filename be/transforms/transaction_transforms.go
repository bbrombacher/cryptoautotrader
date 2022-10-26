package transforms

import (
	transactionRequest "bbrombacher/cryptoautotrader/be/controllers/transactions/request"
	storageModels "bbrombacher/cryptoautotrader/be/storage/models"
)

func BuildGetTransactionsParamsFromGetTransactionsRequest(request transactionRequest.GetTransactionsRequest) storageModels.GetTransactionsParams {
	return storageModels.GetTransactionsParams{
		UserID: request.UserID,
		Cursor: request.Cursor,
		Limit:  request.Limit,
	}
}
