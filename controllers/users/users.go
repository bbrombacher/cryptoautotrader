package users

import (
	"bbrombacher/cryptoautotrader/storage"
	"bbrombacher/cryptoautotrader/storage/models"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	StorageClient *storage.StorageClient
}

func (c Controller) Register(r *mux.Router) *mux.Router {
	r.HandleFunc("/users/{id}", c.GetUser())
	return r
}

func (c Controller) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		userEntry, err := c.StorageClient.GetUser(r.Context(), id)
		if err != nil {
			if errors.Is(err, models.ErrUserDoesNotExist) {
				http.Error(w, "could not find user", http.StatusNotFound)
			}
			http.Error(w, "an unexpected error occurred", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(userEntry)
		if err != nil {
			http.Error(w, "failed to marshal", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (c Controller) UpdateUser(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (c Controller) DeleteUser(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
