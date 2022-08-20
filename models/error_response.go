package models

type ErrorResponse struct {
	Error ErrorMessage `json:"error,omitempty"`
}
type ErrorMessage struct {
	Message string `json:"message"`
}
