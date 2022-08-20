package models

import "bbrombacher/cryptoautotrader/storage/models"

type GetUserResponse struct {
	User *models.UserEntry `json:"user,omitempty"`
}

type ErrorResponse struct {
	Error ErrorMessage `json:"error,omitempty"`
}
type ErrorMessage struct {
	Message string `json:"message"`
}
