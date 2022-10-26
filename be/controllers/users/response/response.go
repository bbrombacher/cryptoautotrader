package response

import "bbrombacher/cryptoautotrader/be/storage/models"

type GetUserResponse struct {
	User *models.UserEntry `json:"user,omitempty"`
}

type CreateUserResponse struct {
	User *models.UserEntry `json:"user,omitempty"`
}

type PatchUserResponse struct {
	User *models.UserEntry `json:"user,omitempty"`
}
