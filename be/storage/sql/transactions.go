package sql

import (
	"bbrombacher/cryptoautotrader/be/storage/models"
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
)

func (s *SqlClient) InsertTransaction(ctx context.Context, entry models.TransactionEntry) (*models.TransactionEntry, error) {
	valueMap, err := entry.RetrieveTagValues("db")
	if err != nil {
		return nil, err
	}

	query := sq.Insert("transactions").SetMap(valueMap).Suffix("RETURNING *")
	sql, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.TransactionEntry
	if err := s.db.GetContext(ctx, &result, sql, args...); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) SelectTransaction(ctx context.Context, id, userID string) (*models.TransactionEntry, error) {

	selectQuery := sq.Select("*").
		From("transactions").
		Where(sq.Eq{"id": id}).
		Where(sq.Eq{"user_id": userID})
	sqlQuery, args, err := selectQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.TransactionEntry
	if err = s.db.GetContext(ctx, &result, sqlQuery, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrTransactionDoesNotExist
		}
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) SelectTransactions(ctx context.Context, params models.GetTransactionsParams) ([]models.TransactionEntry, error) {

	limit := params.Limit
	if limit == 0 {
		limit = 100
	}

	selectQuery := sq.Select("*").
		From("transactions").
		Where(sq.GtOrEq{"cursor_id": params.Cursor}).
		Where(sq.Eq{"user_id": params.UserID}).
		Limit(uint64(limit))
	sqlQuery, args, err := selectQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result []models.TransactionEntry
	if err = s.db.SelectContext(ctx, &result, sqlQuery, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrTransactionDoesNotExist
		}
		return nil, err
	}

	return result, nil
}
