package sql

//go:generate mockgen -source=sql.go -destination=../../mocks/storage/sql.go -package=mocksql github.com/bbrombacher/cryptoautotrader/storage/sql

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type SQLClient interface {
	SelectUser(ctx context.Context, id string) (*models.UserEntry, error)
	InsertUser(ctx context.Context, entry models.UserEntry) (*models.UserEntry, error)
	UpdateUser(ctx context.Context, entry models.UserEntry, updateColumns []string) (*models.UserEntry, error)
	DeleteUser(ctx context.Context, id string) error

	SelectCurrencies(ctx context.Context, params models.GetCurrenciesParams) ([]models.CurrencyEntry, error)
	SelectCurrency(ctx context.Context, id string) (*models.CurrencyEntry, error)
	InsertCurrency(ctx context.Context, entry models.CurrencyEntry) (*models.CurrencyEntry, error)
	UpdateCurrency(ctx context.Context, entry models.CurrencyEntry, updateColumns []string) (*models.CurrencyEntry, error)
	DeleteCurrency(ctx context.Context, id string) error

	UpsertBalance(ctx context.Context, entry models.BalanceEntry) (*models.BalanceEntry, error)
	SelectBalance(ctx context.Context, userID, currencyID string) (*models.BalanceEntry, error)
	SelectBulkBalance(ctx context.Context, userID string) ([]models.BalanceEntry, error)

	InsertTransaction(ctx context.Context, entry models.TransactionEntry) (*models.TransactionEntry, error)
	SelectTransaction(ctx context.Context, id, userID string) (*models.TransactionEntry, error)
	SelectTransactions(ctx context.Context, params models.GetTransactionsParams) ([]models.TransactionEntry, error)

	UpsertTradeSession(ctx context.Context, entry models.TradeSessionEntry) (*models.TradeSessionEntry, error)
	SelectTradeSession(ctx context.Context, userID, sessionID string) (*models.TradeSessionEntry, error)
	SelectTradeSessions(ctx context.Context, params models.GetTradeSessionsParams) ([]models.TradeSessionEntry, error)
}

type SqlClient struct {
	db *sqlx.DB
}

func NewSQLClient(ctx context.Context, db *sqlx.DB) (*SqlClient, error) {
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return &SqlClient{db: db}, nil
}

func enumeratePlaceholders(input string, existingPlaceholder string, args []interface{}) string {
	var placeholderNumeration []interface{}
	for i := range args {
		placeholderNumeration = append(placeholderNumeration, i+1)
	}

	return fmt.Sprintf(strings.ReplaceAll(input, existingPlaceholder, "$%v"), placeholderNumeration...)
}
