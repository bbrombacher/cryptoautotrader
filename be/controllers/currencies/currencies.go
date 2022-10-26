package currencies

import (
	currencyRequest "bbrombacher/cryptoautotrader/be/controllers/currencies/request"
	currencyResponse "bbrombacher/cryptoautotrader/be/controllers/currencies/response"
	"bbrombacher/cryptoautotrader/be/controllers/helpers"
	"bbrombacher/cryptoautotrader/be/storage"
	storageModels "bbrombacher/cryptoautotrader/be/storage/models"
	"bbrombacher/cryptoautotrader/be/transforms"
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

	v1Router.HandleFunc("/currencies/{id}", c.GetCurrency()).Methods(http.MethodGet)
	v1Router.HandleFunc("/currencies", c.CreateCurrency()).Methods(http.MethodPost)
	v1Router.HandleFunc("/currencies", c.GetCurrencies()).Methods(http.MethodGet)
	v1Router.HandleFunc("/currencies/{id}", c.DeleteCurrency()).Methods(http.MethodDelete)
	v1Router.HandleFunc("/currencies/{id}", c.UpdateCurrency()).Methods(http.MethodPatch)
	return v1Router
}

func (c Controller) GetCurrencies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req currencyRequest.GetCurrenciesRequest
		err := req.ParseRequest(r)
		if err != nil {
			helpers.ErrResponse(w, http.StatusUnprocessableEntity, "error parsing request query parameters")
			return
		}

		params := transforms.BuildGetCurrenciesParamsFromGetCurrenciesRequest(req)
		entries, err := c.StorageClient.GetCurrencies(r.Context(), params)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "error getting currencies")
			return
		}

		resp := currencyResponse.GetCurrenciesResponse{
			Currencies: entries,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) GetCurrency() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a currency id")
			return
		}

		currencyEntry, err := c.StorageClient.GetCurrency(r.Context(), id)
		if err != nil {
			if errors.Is(err, storageModels.ErrCurrencyDoesNotExist) {
				helpers.ErrResponse(w, http.StatusNotFound, "could not find currency")
				return
			}
			helpers.ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
			return
		}

		resp := currencyResponse.GetCurrencyResponse{
			Currency: currencyEntry,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) CreateCurrency() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req currencyRequest.PostCurrencyRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			helpers.ErrResponse(w, http.StatusBadRequest, "could not unmarshal request body")
			return
		}

		err = req.Validate()
		if err != nil {
			helpers.ErrResponse(w, http.StatusBadRequest, "request body fails validation")
			return
		}

		currencyEntry := transforms.BuildCurrencyEntryFromPostRequest(req)
		newCurrencyEntry, err := c.StorageClient.CreateCurrency(r.Context(), currencyEntry)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
			return
		}

		resp := currencyResponse.CreateCurrencyResponse{
			Currency: newCurrencyEntry,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) UpdateCurrency() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req currencyRequest.PatchCurrencyRequest
		err := req.ParseRequest(r)
		if err != nil {
			helpers.ErrResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		entry := transforms.BuildCurrencyEntryFromPatchRequest(req)
		updatedEntry, err := c.StorageClient.UpdateCurrency(r.Context(), entry, req.SuppliedFields.Array())
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "failed to update the currency")
			return
		}

		resp := currencyResponse.PatchCurrencyResponse{
			Currency: updatedEntry,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) DeleteCurrency() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a currency id")
			return
		}

		err := c.StorageClient.DeleteCurrency(r.Context(), id)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "failed to delete the currency")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
