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
	r.HandleFunc("/users/{id}", c.GetUser()).Methods("GET")
	r.HandleFunc("/users", c.GetUser()).Methods("POST")
	return r
}

func ErrResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: ErrorMessage{
			Message: "could not find",
		},
	})
}

func (c Controller) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		userEntry, err := c.StorageClient.GetUser(r.Context(), id)
		if err != nil {
			if errors.Is(err, models.ErrUserDoesNotExist) {
				ErrResponse(w, http.StatusNotFound, "could not find user")
				return
			}
			ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userEntry)
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
