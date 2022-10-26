package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

var (
	validate = validator.New()
)

type UpdateBalanceRequest struct {
	UserID     string `json:"-"`
	CurrencyID string `json:"currency_id" validate:"required"`
	Amount     string `json:"amount" validate:"required"`
}

func (r *UpdateBalanceRequest) ParseRequest(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	defer request.Body.Close()

	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	userID := request.Header.Get("x-user-id")
	if userID == "" {
		return errors.New("x-user-id is required")
	}

	r.UserID = userID

	return nil
}

func (r *UpdateBalanceRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}

	_, err := decimal.NewFromString(r.Amount)
	if err != nil {
		return err
	}

	return nil
}
