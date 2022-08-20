package request

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

var (
	decoder        = schema.NewDecoder()
	ErrParseParams = errors.New("unable to parse parameters")
	validate       = validator.New()
)

func init() {
	decoder.SetAliasTag("json")
	decoder.IgnoreUnknownKeys(true)
}

type PostUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (r *PostUserRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}
