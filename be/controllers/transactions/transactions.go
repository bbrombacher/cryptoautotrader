package transactions

import (
	"bbrombacher/cryptoautotrader/be/controllers/helpers"
	"bbrombacher/cryptoautotrader/be/storage"
	"bbrombacher/cryptoautotrader/be/transforms"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	transactionRequest "bbrombacher/cryptoautotrader/be/controllers/transactions/request"
	transactionResponse "bbrombacher/cryptoautotrader/be/controllers/transactions/response"
	storageModels "bbrombacher/cryptoautotrader/be/storage/models"

	"github.com/gorilla/mux"
)

type Controller struct {
	StorageClient *storage.StorageClient
}

func (c Controller) Register(r *mux.Router) *mux.Router {
	v1Router := r.PathPrefix("/v1").Subrouter()

	v1Router.HandleFunc("/transactions/{id}", c.GetTransaction()).Methods(http.MethodGet)
	v1Router.HandleFunc("/transactions", c.GetTransactions()).Methods(http.MethodGet, http.MethodOptions)

	return v1Router
}

func (c Controller) GetTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		transactionID := vars["id"]
		if transactionID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a transaction id")
			return
		}

		userID := r.Header.Get("x-user-id")
		if userID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a x-user-id")
			return
		}

		transactionEntry, err := c.StorageClient.GetTransaction(r.Context(), transactionID, userID)
		if err != nil {
			if errors.Is(err, storageModels.ErrTransactionDoesNotExist) {
				helpers.ErrResponse(w, http.StatusNotFound, "could not find transaction")
				return
			}
			helpers.ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
			return
		}

		transactions := []storageModels.TransactionEntry{*transactionEntry}
		resp := transactionResponse.GetTransactionsResponse{
			Transactions: transactions,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) GetTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("GetTransactions: Access-Control-Allow-Origin: ", w.Header().Get("Access-Control-Allow-Origin"))

		var req transactionRequest.GetTransactionsRequest

		err := req.ParseRequest(r)
		if err != nil {
			helpers.ErrResponse(w, http.StatusUnprocessableEntity, "error parsing request query parameters")
			return
		}

		params := transforms.BuildGetTransactionsParamsFromGetTransactionsRequest(req)
		result, err := c.StorageClient.GetTransactions(r.Context(), params)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "error getting transactions")
			return
		}

		resp := transactionResponse.GetTransactionsResponse{
			Transactions: result,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
