package sql

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (s *SqlClient) UpsertBalance(ctx context.Context, entry models.BalanceEntry) (*models.BalanceEntry, error) {

	valueMap, err := entry.RetrieveTagValues("db")
	if err != nil {
		return nil, err
	}

	insertQuery := sq.Insert("balance").
		SetMap(valueMap).
		Suffix(`ON CONFLICT ON CONSTRAINT balance_pk DO`)

	insertSQL, insertArgs, err := insertQuery.ToSql()
	if err != nil {
		return nil, err
	}

	updateQuery := sq.Update(" ").
		Where(sq.Eq{"balance.user_id": entry.UserID}).
		Where(sq.Eq{"balance.currency_id": entry.CurrencyID}).
		Set("updated_at", time.Now()).
		Suffix(`RETURNING *`)

	updateColumns := []string{"amount"}
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

	var result models.BalanceEntry
	if err = s.db.GetContext(ctx, &result, query, args...); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) SelectBalance(ctx context.Context, userID, currencyID string) (*models.BalanceEntry, error) {

	selectQuery := sq.Select("*").
		From("balance").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"currency_id": currencyID})

	sqlQuery, args, err := selectQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.BalanceEntry
	err = s.db.GetContext(ctx, &result, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) SelectBulkBalance(ctx context.Context, userID string) ([]models.BalanceEntry, error) {

	selectQuery := sq.Select("*").
		From("balance").
		Where(sq.Eq{"user_id": userID})

	sqlQuery, args, err := selectQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result []models.BalanceEntry
	err = s.db.SelectContext(ctx, &result, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
