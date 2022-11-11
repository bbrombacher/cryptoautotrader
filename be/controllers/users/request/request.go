package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

type PostLoginRequest struct {
	First string `json:"first" validate:"required"`
	Last  string `json:"last" validate:"required"`
}

func (r *PostLoginRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
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

type PatchUserRequest struct {
	ID             string         `json:"id"`
	FirstName      *string        `json:"first_name"`
	LastName       *string        `json:"last_name"`
	SuppliedFields SuppliedFields `json:"-" validate:"required,gte=1"`
}

func (r *PatchUserRequest) ParseRequest(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	defer request.Body.Close()

	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &r.SuppliedFields)
	if err != nil {
		return err
	}

	vars := mux.Vars(request)
	id := vars["id"]
	r.ID = id

	return nil
}

type SuppliedFields map[string]struct{}

// UnmarshalJSON fills NullFields with the names of keys that are set to null in a JSON string
func (s *SuppliedFields) UnmarshalJSON(data []byte) error {
	valueMap := make(map[string]interface{})
	if err := json.Unmarshal(data, &valueMap); err != nil {
		return err
	}

	suppliedFields := make(SuppliedFields, len(valueMap))
	for key := range valueMap {
		suppliedFields[key] = struct{}{}
	}

	*s = suppliedFields
	return nil
}

func (s SuppliedFields) Array() []string {
	result := make([]string, 0, len(s))

	for key := range s {
		result = append(result, key)
	}

	sort.Strings(result)
	return result
}
