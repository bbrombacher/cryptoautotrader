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

type GetCurrenciesRequest struct {
	Limit  int
	Cursor int
}

func (r *GetCurrenciesRequest) ParseRequest(request *http.Request) error {

	err := request.ParseForm()
	if err != nil {
		return errors.New("error parsing query parameters")
	}

	err = decoder.Decode(r, request.Form)
	if err != nil {
		return errors.New("error parsing query parameters")
	}

	return nil
}

type PostCurrencyRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (r *PostCurrencyRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}

type PatchCurrencyRequest struct {
	ID             string         `json:"id"`
	Name           *string        `json:"name"`
	Description    *string        `json:"description"`
	SuppliedFields SuppliedFields `json:"-" validate:"required,gte=1"`
}

func (r *PatchCurrencyRequest) ParseRequest(request *http.Request) error {
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
