package models

import (
	"errors"
	"fmt"
	"time"

	jsoninter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

var (
	ErrUserDoesNotExist         = errors.New("user does not exist")
	ErrTransactionDoesNotExist  = errors.New("transaction does not exist")
	ErrCurrencyDoesNotExist     = errors.New("currency does not exist")
	ErrTradeSessionDoesNotExist = errors.New("trade session does not exist")
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

type Currencies []CurrencyEntry

func (c Currencies) GetCurrencyIDByName(name string) (string, error) {
	for _, value := range c {
		if value.Name == name {
			return value.ID, nil
		}
	}

	return "", fmt.Errorf("could not find currency with name %s", name)
}

type GetCurrenciesParams struct {
	Cursor int
	Limit  int
}

type BalanceEntry struct {
	UserID     string          `db:"user_id" json:"user_id"`
	CurrencyID string          `db:"currency_id" json:"currency_id"`
	Amount     decimal.Decimal `db:"amount" json:"amount"`
	UpdatedAt  *time.Time      `db:"updated_at,omitempty" json:"updated_at"`

	Currency CurrencyEntry `db:"-" json:"-"`
}

func (e *BalanceEntry) RetrieveTagValues(tag string) (map[string]interface{}, error) {
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

type TransactionEntry struct {
	ID              string          `db:"id" json:"id"`
	UserID          string          `db:"user_id" json:"user_id"`
	UseCurrencyID   string          `db:"use_currency_id" json:"use_currency_id"`
	GetCurrencyID   string          `db:"get_currency_id" json:"get_currency_id"`
	TransactionType string          `db:"transaction_type" json:"transaction_type"` // buy/sell
	Amount          decimal.Decimal `db:"amount" json:"amount"`
	Price           decimal.Decimal `db:"price" json:"price"`

	CursorID  int        `db:"cursor_id,omitempty" json:"cursor_id"`
	CreatedAt *time.Time `db:"created_at,omitempty" json:"created_at"`
}

func (e *TransactionEntry) RetrieveTagValues(tag string) (map[string]interface{}, error) {
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

type GetTransactionsParams struct {
	UserID string
	Cursor int
	Limit  int
}

type TradeSessionEntry struct {
	ID              string          `db:"id" json:"id"`
	UserID          string          `db:"user_id" json:"user_id"`
	Algorithm       string          `db:"algorithm" json:"algorithm"`
	CurrencyID      string          `db:"currency_id" json:"currency_id"`
	StartingBalance decimal.Decimal `db:"starting_balance" json:"starting_balance"`
	EndingBalance   decimal.Decimal `db:"ending_balance" json:"ending_balance"`

	CursorID  int        `db:"cursor_id,omitempty" json:"cursor_id"`
	StartedAt *time.Time `db:"started_at,omitempty" json:"started_at"`
	EndedAt   *time.Time `db:"ended_at,omitempty" json:"ended_at"`
}

func (e *TradeSessionEntry) RetrieveTagValues(tag string) (map[string]interface{}, error) {
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

type GetTradeSessionsParams struct {
	UserID string
	Cursor int
	Limit  int
}
