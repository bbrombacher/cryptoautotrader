package helpers

import (
	"bbrombacher/cryptoautotrader/models"
	"encoding/json"
	"net/http"
)

func ErrResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error: models.ErrorMessage{
			Message: message,
		},
	})
}
