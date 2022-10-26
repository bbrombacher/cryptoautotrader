package sql

import (
	"bbrombacher/cryptoautotrader/be/storage/models"
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (s *SqlClient) InsertTransactionSessionMapEntry(ctx context.Context, entry models.TransactionSessionMapEntry) (*models.TransactionSessionMapEntry, error) {
	valueMap, err := entry.RetrieveTagValues("db")
	if err != nil {
		return nil, err
	}

	query := sq.Insert("transaction_sessions_map").SetMap(valueMap).Suffix("RETURNING *")
	sql, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.TransactionSessionMapEntry
	if err := s.db.GetContext(ctx, &result, sql, args...); err != nil {
		return nil, err
	}

	return &result, nil
}
