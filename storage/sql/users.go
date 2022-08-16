package sql

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

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

func (s *sqlClient) DeleteUser(ctx context.Context, id string) error {

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
		return fmt.Errorf("user does not exist")
	}

	return nil
}
