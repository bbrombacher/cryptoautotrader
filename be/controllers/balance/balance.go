package balance

import (
	balanceRequest "bbrombacher/cryptoautotrader/be/controllers/balance/request"
	"bbrombacher/cryptoautotrader/be/controllers/helpers"
	"bbrombacher/cryptoautotrader/be/storage"
	"bbrombacher/cryptoautotrader/be/transforms"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	StorageClient *storage.StorageClient
}

func (c Controller) Register(r *mux.Router) *mux.Router {
	r.HandleFunc("/v1/balance", c.GetBulkBalance()).Methods(http.MethodGet)
	r.HandleFunc("/v1/balance/{currency_id}", c.GetBalance()).Methods(http.MethodGet)
	r.HandleFunc("/v1/balance", c.UpdateBalance()).Methods(http.MethodPatch)
	return r
}

func (c Controller) GetBulkBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Header.Get("x-user-id")
		if userID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "must provide x-user-id")
			return
		}

		entries, err := c.StorageClient.GetBulkBalance(r.Context(), userID)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "could not get balance")
			return
		}

		resp := transforms.BuildBulkBalanceResponseObjectFromBalanceEntry(entries)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) GetBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Header.Get("x-user-id")
		if userID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "must provide x-user-id")
			return
		}

		vars := mux.Vars(r)
		currencyID := vars["currency_id"]
		if currencyID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a currency id")
			return
		}

		entry, err := c.StorageClient.GetBalance(r.Context(), userID, currencyID)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "could not get balance")
			return
		}

		resp := transforms.BuildBalanceResponseObjectFromBalanceEntry(entry)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) UpdateBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateBalanceRequest balanceRequest.UpdateBalanceRequest
		err := updateBalanceRequest.ParseRequest(r)
		if err != nil {
			helpers.ErrResponse(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		err = updateBalanceRequest.Validate()
		if err != nil {
			helpers.ErrResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		input := transforms.BuildUpdateBalanceEntryFromUpdateBalanceRequest(updateBalanceRequest)
		entry, err := c.StorageClient.UpdateBalance(r.Context(), input)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp := transforms.BuildBalanceResponseObjectFromBalanceEntry(entry)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
