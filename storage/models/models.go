package models

import (
	"errors"
	"time"

	jsoninter "github.com/json-iterator/go"
)

var (
	ErrUserDoesNotExist     = errors.New("user does not exist")
	ErrCurrencyDoesNotExist = errors.New("currency does not exist")
)

type UserEntry struct {
	ID        string `db:"id" json:"id"`
	FirstName string `db:"first_name,omitempty" json:"first_name"`
	LastName  string `db:"last_name,omitempty" json:"last_name"`

	CursorID  int        `db:"cursor_id,omitempty" json:"cursor_id"`
	CreatedAt *time.Time `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at,omitempty" json:"updated_at"`
}

func (e *UserEntry) RetrieveTagValues(tag string) (map[string]interface{}, error) {
	tagMap := map[string]interface{}{}
	var rjson = jsoninter.Config{TagKey: tag}.Froze()
	data, err := rjson.Marshal(e)
	if err != nil {
		return nil, err
	}
	if err := rjson.Unmarshal(data, &tagMap); err != nil {
		return nil, err
	}
	return tagMap, nil
}

type CurrencyEntry struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name,omitempty" json:"name"`
	Description string `db:"description,omitempty" json:"description"`

	CursorID  int        `db:"cursor_id,omitempty" json:"cursor_id"`
	CreatedAt *time.Time `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at,omitempty" json:"updated_at"`
}

func (e *CurrencyEntry) RetrieveTagValues(tag string) (map[string]interface{}, error) {
	tagMap := map[string]interface{}{}
	var rjson = jsoninter.Config{TagKey: tag}.Froze()
	data, err := rjson.Marshal(e)
	if err != nil {
		return nil, err
	}
	if err := rjson.Unmarshal(data, &tagMap); err != nil {
		return nil, err
	}
	return tagMap, nil
}

type SelectCurrenciesParams struct {
	Cursor int
	Limit  int
}
