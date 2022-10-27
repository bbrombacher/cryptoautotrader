package users

import (
	"bbrombacher/cryptoautotrader/be/controllers/helpers"
	userRequest "bbrombacher/cryptoautotrader/be/controllers/users/request"
	userResponse "bbrombacher/cryptoautotrader/be/controllers/users/response"
	"bbrombacher/cryptoautotrader/be/storage"
	storageModels "bbrombacher/cryptoautotrader/be/storage/models"
	"bbrombacher/cryptoautotrader/be/transforms"
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
	v1Router := r.PathPrefix("/v1").Subrouter()

	v1Router.HandleFunc("/users/{id}", c.GetUser()).Methods(http.MethodGet)
	v1Router.HandleFunc("/users", c.CreateUser()).Methods(http.MethodPost)
	v1Router.HandleFunc("/users/{id}", c.DeleteUser()).Methods(http.MethodDelete)
	v1Router.HandleFunc("/users/{id}", c.UpdateUser()).Methods(http.MethodPatch)
	return v1Router
}

func (c Controller) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("cors?", w.Header().Get("Access-Control-Allow-Origin"))

		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a user id")
			return
		}

		userEntry, err := c.StorageClient.GetUser(r.Context(), id)
		if err != nil {
			if errors.Is(err, storageModels.ErrUserDoesNotExist) {
				helpers.ErrResponse(w, http.StatusNotFound, "could not find user")
				return
			}
			helpers.ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
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
			helpers.ErrResponse(w, http.StatusBadRequest, "could not unmarshal request body")
			return
		}

		err = req.Validate()
		if err != nil {
			helpers.ErrResponse(w, http.StatusBadRequest, "request body fails validation")
			return
		}

		userEntry := transforms.BuildUserEntryFromPostRequest(req)
		newUserEntry, err := c.StorageClient.CreateUser(r.Context(), userEntry)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "an unexpected error occurred")
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
			helpers.ErrResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		entry := transforms.BuildUserEntryFromPatchRequest(req)
		updatedEntry, err := c.StorageClient.UpdateUser(r.Context(), entry, req.SuppliedFields.Array())
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "failed to update the user")
			return
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
			helpers.ErrResponse(w, http.StatusBadRequest, "you must provide a user id")
			return
		}

		err := c.StorageClient.DeleteUser(r.Context(), id)
		if err != nil {
			helpers.ErrResponse(w, http.StatusInternalServerError, "failed to delete the user")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
