package request

import (
	"errors"
	"net/http"

	"github.com/gorilla/schema"
)

var (
	decoder = schema.NewDecoder()
)

type GetTradeSessionsRequest struct {
	UserID string
	Cursor int
	Limit  int
}

func (r *GetTradeSessionsRequest) ParseRequest(request *http.Request) error {
	err := request.ParseForm()
	if err != nil {
		return errors.New("error parsing query parameters")
	}

	err = decoder.Decode(r, request.Form)
	if err != nil {
		return errors.New("error parsing query parameters")
	}

	userID := request.Header.Get("x-user-id")
	if userID == "" {
		return errors.New("must supply x-user-id")
	}

	r.UserID = userID

	return nil
}
