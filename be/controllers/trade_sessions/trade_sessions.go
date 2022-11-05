package trade_sessions

import (
	"bbrombacher/cryptoautotrader/be/controllers/helpers"
	tradeSessionsRequest "bbrombacher/cryptoautotrader/be/controllers/trade_sessions/request"
	tradeSessionsResponse "bbrombacher/cryptoautotrader/be/controllers/trade_sessions/response"
	"bbrombacher/cryptoautotrader/be/storage"
	storageModels "bbrombacher/cryptoautotrader/be/storage/models"
	"bbrombacher/cryptoautotrader/be/tradebot"
	"bbrombacher/cryptoautotrader/be/transforms"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	StorageClient *storage.StorageClient
	Bot           tradebot.Bot
}

func (c Controller) Register(r *mux.Router) *mux.Router {
	r.HandleFunc("/v1/trade-sessions", c.GetTradeSessions()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/v1/trade-sessions/{id}", c.GetTradeSession()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/v1/trade-sessions/start", c.StartSession()).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/v1/trade-sessions/stop/{trade-session-id}", c.StopSession()).Methods(http.MethodPost, http.MethodOptions)
	return r
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

func (c Controller) StartSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("x-user-id")
		if userID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a x-user-id")
			return
		}

		tradeID, err := c.Bot.StartTrading(userID, 30)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to start trading: %v", err.Error()))
			return
		}

		resp := map[string]interface{}{
			"trade_id": tradeID,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) StopSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("x-user-id")
		if userID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a x-user-id")
			return
		}

		vars := mux.Vars(r)
		tradeSessionID := vars["trade-session-id"]
		if tradeSessionID == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a trade session id")
			return
		}

		err := c.Bot.StopTrading(userID, tradeSessionID)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to stop trading: %v", err.Error()))
			return
		}

		resp := map[string]interface{}{
			"result": "success",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
