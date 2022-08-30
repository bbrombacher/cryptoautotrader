package currencies

import (
	currencyRequest "bbrombacher/cryptoautotrader/controllers/currencies/request"
	currencyResponse "bbrombacher/cryptoautotrader/controllers/currencies/response"
	"bbrombacher/cryptoautotrader/models"
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

	v1Router.HandleFunc("/currencies/{id}", c.GetCurrency()).Methods(http.MethodGet)
	v1Router.HandleFunc("/currencies", c.CreateCurrency()).Methods(http.MethodPost)
	v1Router.HandleFunc("/currencies/{id}", c.DeleteCurrency()).Methods(http.MethodDelete)
	v1Router.HandleFunc("/currencies/{id}", c.UpdateCurrency()).Methods(http.MethodPatch)
	return v1Router
}

func ErrResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error: models.ErrorMessage{
			Message: message,
		},
	})
}

func (c Controller) GetCurrency() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			ErrResponse(w, http.StatusBadRequest, "you must provide a currency id")
			return
		}

		currencyEntry, err := c.StorageClient.GetCurrency(r.Context(), id)
		if err != nil {
			if errors.Is(err, storageModels.ErrCurrencyDoesNotExist) {
				ErrResponse(w, http.StatusNotFound, "could not find currency")
				return
			}
			ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
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
			ErrResponse(w, http.StatusBadRequest, "could not unmarshal request body")
			return
		}

		err = req.Validate()
		if err != nil {
			ErrResponse(w, http.StatusBadRequest, "request body fails validation")
			return
		}

		currencyEntry := transforms.BuildCurrencyEntryFromPostRequest(req)
		newCurrencyEntry, err := c.StorageClient.CreateCurrency(r.Context(), currencyEntry)
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
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
			ErrResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		entry := transforms.BuildCurrencyEntryFromPatchRequest(req)
		updatedEntry, err := c.StorageClient.UpdateCurrency(r.Context(), entry, req.SuppliedFields.Array())
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, "failed to update the currency")
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
			ErrResponse(w, http.StatusBadRequest, "you must provide a currency id")
			return
		}

		err := c.StorageClient.DeleteCurrency(r.Context(), id)
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, "failed to delete the currency")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
