package users

import (
	userRequest "bbrombacher/cryptoautotrader/controllers/users/request"
	userResponse "bbrombacher/cryptoautotrader/controllers/users/response"
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

	v1Router.HandleFunc("/users/{id}", c.GetUser()).Methods(http.MethodGet)
	v1Router.HandleFunc("/users", c.CreateUser()).Methods(http.MethodPost)
	v1Router.HandleFunc("/users/{id}", c.DeleteUser()).Methods(http.MethodDelete)
	v1Router.HandleFunc("/users/{id}", c.UpdateUser()).Methods(http.MethodPatch)
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

		resp := userResponse.GetUserResponse{
			User: userEntry,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userRequest.PostUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			ErrResponse(w, http.StatusBadRequest, "could not unmarshal request body")
		}

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

		resp := userResponse.CreateUserResponse{
			User: newUserEntry,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req userRequest.PatchUserRequest
		err := req.ParseRequest(r)
		if err != nil {
			ErrResponse(w, http.StatusBadRequest, "failed to parse request")
		}

		entry := transforms.BuildUserEntryFromPatchRequest(req)
		updatedEntry, err := c.StorageClient.UpdateUser(r.Context(), entry, req.SuppliedFields.Array())
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, "failed to update the user")
		}

		resp := userResponse.PatchUserResponse{
			User: updatedEntry,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func (c Controller) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			ErrResponse(w, http.StatusBadRequest, "you must provide a user id")
			return
		}

		err := c.StorageClient.DeleteUser(r.Context(), id)
		if err != nil {
			ErrResponse(w, http.StatusInternalServerError, "failed to delete the user")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
