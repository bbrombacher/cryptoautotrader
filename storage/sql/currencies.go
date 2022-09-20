package sql

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (s *SqlClient) SelectCurrency(ctx context.Context, id string) (*models.CurrencyEntry, error) {

	selectQuery := sq.Select("*").
		From("currencies").
		Where(sq.Eq{"id": id})
	sqlQuery, args, err := selectQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.CurrencyEntry
	if err = s.db.GetContext(ctx, &result, sqlQuery, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrCurrencyDoesNotExist
		}
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) SelectCurrencies(ctx context.Context, params models.GetCurrenciesParams) ([]models.CurrencyEntry, error) {

	limit := params.Limit
	if limit == 0 {
		limit = 100
	}

	selectQuery := sq.Select("*").
		From("currencies").
		Where(sq.GtOrEq{"cursor_id": params.Cursor}).
		Limit(uint64(limit))
	sqlQuery, args, err := selectQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result []models.CurrencyEntry
	if err = s.db.SelectContext(ctx, &result, sqlQuery, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrCurrencyDoesNotExist
		}
		return nil, err
	}

	return result, nil
}

func (s *SqlClient) InsertCurrency(ctx context.Context, entry models.CurrencyEntry) (*models.CurrencyEntry, error) {
	valueMap, err := entry.RetrieveTagValues("db")
	if err != nil {
		return nil, err
	}

	query := sq.Insert("currencies").SetMap(valueMap).Suffix("RETURNING *")
	sql, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.CurrencyEntry
	if err := s.db.GetContext(ctx, &result, sql, args...); err != nil {
		// TODO: find error code to check if unique key already exists
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) UpdateCurrency(ctx context.Context, entry models.CurrencyEntry, updateColumns []string) (*models.CurrencyEntry, error) {

	valueMap, err := entry.RetrieveTagValues("db")
	if err != nil {
		return nil, err
	}

	updateQuery := sq.Update("currencies").
		Where(sq.Eq{"id": entry.ID}).
		Set("updated_at", time.Now()).
		Suffix("RETURNING *")

	for _, tag := range updateColumns {
		if val, ok := valueMap[tag]; ok {
			updateQuery = updateQuery.Set(tag, val)
		}
	}

	updateSql, updateArgs, err := updateQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.CurrencyEntry
	if err = s.db.GetContext(ctx, &result, updateSql, updateArgs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrCurrencyDoesNotExist
		}
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) DeleteCurrency(ctx context.Context, id string) error {

	query := sq.Delete("currencies").Where(sq.Eq{"id": id})
	sql, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	sqlResult, err := s.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	if rows, _ := sqlResult.RowsAffected(); rows == 0 {
		return models.ErrCurrencyDoesNotExist
	}

	return nil
}
