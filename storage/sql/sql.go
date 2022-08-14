package sql

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"

	sq "github.com/Masterminds/squirrel"

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

func (s *sqlClient) InsertUser(ctx context.Context, entry models.UserEntry) (*models.UserEntry, error) {
	valueMap, err := entry.RetrieveTagValues("db")
	if err != nil {
		return nil, err
	}

	query := sq.Insert("users").SetMap(valueMap).Suffix("RETURNING *")
	sql, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.UserEntry
	if err := s.db.GetContext(ctx, &result, sql, args...); err != nil {
		return nil, err
	}

	return &result, nil
}
