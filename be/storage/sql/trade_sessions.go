package sql

import (
	"bbrombacher/cryptoautotrader/be/storage/models"
	"context"
	"database/sql"
	"errors"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

func (s *SqlClient) UpsertTradeSession(ctx context.Context, entry models.TradeSessionEntry) (*models.TradeSessionEntry, error) {

	valueMap, err := entry.RetrieveTagValues("db")
	if err != nil {
		return nil, err
	}

	insertQuery := sq.Insert("trade_sessions").
		SetMap(valueMap).
		Suffix(`ON CONFLICT (id) DO`)

	insertSQL, insertArgs, err := insertQuery.ToSql()
	if err != nil {
		return nil, err
	}

	updateQuery := sq.Update(" ").
		Where(sq.Eq{"trade_sessions.id": entry.ID}).
		Where(sq.Eq{"trade_sessions.user_id": entry.UserID}).
		Where(sq.Eq{"trade_sessions.currency_id": entry.CurrencyID}).
		Suffix(`RETURNING *`)

	updateColumns := []string{"ended_at", "ending_balance"}
	for _, tag := range updateColumns {
		if val, ok := valueMap[tag]; ok {
			updateQuery = updateQuery.Set(tag, val)
		}
	}

	updateSql, updateArgs, err := updateQuery.ToSql()
	if err != nil {
		return nil, err
	}

	sql := strings.Join([]string{insertSQL, updateSql}, " ")
	args := append(insertArgs, updateArgs...)

	query := enumeratePlaceholders(sql, "?", args)

	var result models.TradeSessionEntry
	if err = s.db.GetContext(ctx, &result, query, args...); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) SelectTradeSession(ctx context.Context, userID, sessionID string) (*models.TradeSessionEntry, error) {

	selectQuery := sq.Select("*").
		From("trade_sessions").
		Where(sq.Eq{"id": sessionID}).
		Where(sq.Eq{"user_id": userID})
	sqlQuery, args, err := selectQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.TradeSessionEntry
	if err = s.db.GetContext(ctx, &result, sqlQuery, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrCurrencyDoesNotExist
		}
		return nil, err
	}
	return &result, nil
}

func (s *SqlClient) SelectTradeSessions(ctx context.Context, params models.GetTradeSessionsParams) ([]models.TradeSessionEntry, error) {

	limit := params.Limit
	if limit == 0 {
		limit = 100
	}

	selectQuery := sq.Select("*").
		From("trade_sessions").
		Where(sq.GtOrEq{"cursor_id": params.Cursor}).
		OrderBy("cursor_id desc").
		Limit(uint64(limit))

	if params.UserID != "" {
		selectQuery = selectQuery.Where(sq.Eq{"user_id": params.UserID})
	}

	if params.OrphanedSessions {
		selectQuery = selectQuery.Where(sq.Eq{"ended_at": nil})
		// TODO: instead query by started_at is within some amount of time?
	}

	sqlQuery, args, err := selectQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result []models.TradeSessionEntry
	if err = s.db.SelectContext(ctx, &result, sqlQuery, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrCurrencyDoesNotExist
		}
		return nil, err
	}

	return result, nil
}
