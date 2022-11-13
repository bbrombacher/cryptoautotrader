package sql

import (
	"bbrombacher/cryptoautotrader/be/storage/models"
	"context"
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (s *SqlClient) SelectUser(ctx context.Context, params models.GetUserParams) (*models.UserEntry, error) {

	selectQuery := sq.Select("*").From("users")

	if params.ID != "" {
		selectQuery = selectQuery.Where(sq.Eq{"id": params.ID})
	} else if params.FirstName != "" && params.LastName != "" {
		selectQuery = selectQuery.Where(sq.Eq{"first_name": params.FirstName}).
			Where(sq.Eq{"last_name": params.LastName})
	}

	sqlQuery, args, err := selectQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var result models.UserEntry
	if err = s.db.GetContext(ctx, &result, sqlQuery, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserDoesNotExist
		}
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) InsertUser(ctx context.Context, entry models.UserEntry) (*models.UserEntry, error) {
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

func (s *SqlClient) UpdateUser(ctx context.Context, entry models.UserEntry, updateColumns []string) (*models.UserEntry, error) {
	valueMap, err := entry.RetrieveTagValues("db")
	if err != nil {
		return nil, err
	}

	updateQuery := sq.Update("users").
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

	var result models.UserEntry
	if err = s.db.GetContext(ctx, &result, updateSql, updateArgs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserDoesNotExist
		}
		return nil, err
	}

	return &result, nil
}

func (s *SqlClient) DeleteUser(ctx context.Context, id string) error {

	query := sq.Delete("users").Where(sq.Eq{"id": id})
	sql, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	sqlResult, err := s.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	if rows, _ := sqlResult.RowsAffected(); rows == 0 {
		return models.ErrUserDoesNotExist
	}

	return nil
}
