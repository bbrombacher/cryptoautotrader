package users

import (
	userRequest "bbrombacher/cryptoautotrader/controllers/users/request"
	"bbrombacher/cryptoautotrader/models"
	"bbrombacher/cryptoautotrader/storage"
	storageModels "bbrombacher/cryptoautotrader/storage/models"
	"bbrombacher/cryptoautotrader/transforms"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	StorageClient *storage.StorageClient
}

func (c Controller) Register(r *mux.Router) *mux.Router {
	r.HandleFunc("/users/{id}", c.GetUser()).Methods("GET")
	r.HandleFunc("/users", c.CreateUser()).Methods("POST")
	return r
}

func ErrResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error: models.ErrorMessage{
			Message: message,
		},
	})
}

func (c Controller) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			ErrResponse(w, http.StatusBadRequest, "you must provide a user id")
			return
		}

		userEntry, err := c.StorageClient.GetUser(r.Context(), id)
		if err != nil {
			if errors.Is(err, storageModels.ErrUserDoesNotExist) {
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

func (c Controller) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userRequest.PostUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			ErrResponse(w, http.StatusBadRequest, "could not unmarshal request body")
		}

		log.Printf("REQ %#v", req)

		err = req.Validate()
		if err != nil {
			ErrResponse(w, http.StatusBadRequest, "request body fails validation")
			return
		}

		userEntry := transforms.BuildUserEntryFromPostRequest(req)
		newUserEntry, err := c.StorageClient.CreateUser(r.Context(), userEntry)
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newUserEntry)
	}
}

func (c Controller) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (c Controller) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
