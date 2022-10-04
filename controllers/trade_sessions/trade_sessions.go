package trade_sessions

import (
	"bbrombacher/cryptoautotrader/controllers/helpers"
	tradeSessionsRequest "bbrombacher/cryptoautotrader/controllers/trade_sessions/request"
	tradeSessionsResponse "bbrombacher/cryptoautotrader/controllers/trade_sessions/response"
	"bbrombacher/cryptoautotrader/storage"
	storageModels "bbrombacher/cryptoautotrader/storage/models"
	"bbrombacher/cryptoautotrader/transforms"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	StorageClient *storage.StorageClient
}

func (c Controller) Register(r *mux.Router) *mux.Router {
	v1Router := r.PathPrefix("/v1").Subrouter()

	v1Router.HandleFunc("/trade-sessions", c.GetTradeSessions()).Methods(http.MethodGet)
	v1Router.HandleFunc("/trade-sessions/{id}", c.GetTradeSession()).Methods(http.MethodGet)

	return v1Router
}

func (c Controller) GetTradeSessions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req tradeSessionsRequest.GetTradeSessionsRequest

		err := req.ParseRequest(r)
		if err != nil {
			helpers.ErrResponse(w, http.StatusUnprocessableEntity, "error parsing request query parameters")
			return
		}

		params := transforms.BuildGetTradeSessionsFromGetRequest(req)
		result, err := c.StorageClient.GetTradeSessions(r.Context(), params)
		if err != nil {
			return
		}

		resp := tradeSessionsResponse.GetTradeSessionsResponse{
			TradeSessions: result,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) GetTradeSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sessionID := vars["id"]
		if sessionID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a trade session id")
			return
		}

		userID := r.Header.Get("x-user-id")
		if userID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a x-user-id")
			return
		}

		tradeSessionEntry, err := c.StorageClient.GetTradeSession(r.Context(), userID, sessionID)
		if err != nil {
			if errors.Is(storageModels.ErrTradeSessionDoesNotExist, err) {
				helpers.ErrResponse(w, http.StatusNotFound, "could not find transaction")
				return
			}
			helpers.ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
			return
		}

		tradeSessions := []storageModels.TradeSessionEntry{*tradeSessionEntry}
		resp := tradeSessionsResponse.GetTradeSessionsResponse{
			TradeSessions: tradeSessions,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
