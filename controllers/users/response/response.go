package response

import "bbrombacher/cryptoautotrader/storage/models"

type GetUserResponse struct {
	User *models.UserEntry `json:"user,omitempty"`
}
