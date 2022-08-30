package sql

//go:generate mockgen -source=sql.go -destination=../../mocks/storage/sql.go -package=mocksql github.com/bbrombacher/cryptoautotrader/storage/sql

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"

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
