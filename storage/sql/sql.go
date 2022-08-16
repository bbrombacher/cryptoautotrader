package sql

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type SQLClient interface{}

type sqlClient struct {
	db *sqlx.DB
}

func NewSQLClient(ctx context.Context, db *sqlx.DB) (*sqlClient, error) {
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return &sqlClient{db: db}, nil
}
