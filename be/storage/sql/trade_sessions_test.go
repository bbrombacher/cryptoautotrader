package sql_test

import (
	"bbrombacher/cryptoautotrader/be/storage/models"
	"bbrombacher/cryptoautotrader/be/storage/sql"
	"context"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSQLClient_UpsertTradeSession_Start(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`
		INSERT INTO users (
			id,
			first_name,
			last_name
		)
		VALUES 
			('U', 'brandon', 'brombacher');

		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('C', 'eth', 'eth coin');
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	id := uuid.New()
	now := time.Now()
	entry := models.TradeSessionEntry{
		ID:              id.String(),
		UserID:          "U",
		Algorithm:       "basic",
		CurrencyID:      "C",
		StartingBalance: decimal.NewFromFloat32(1.25),
		StartedAt:       &now,
	}

	// upset entry
	_, err = sqlClient.UpsertTradeSession(context.Background(), entry)
	assert.Nil(t, err)
}

func TestSQLClient_UpsertTradeSession_End(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`
		INSERT INTO users (
			id,
			first_name,
			last_name
		)
		VALUES 
			('U', 'brandon', 'brombacher');

		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('C', 'eth', 'eth coin');
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	id := uuid.New()
	startNow := time.Now()
	start := models.TradeSessionEntry{
		ID:              id.String(),
		UserID:          "U",
		Algorithm:       "basic",
		CurrencyID:      "C",
		StartingBalance: decimal.NewFromFloat32(1.25),
		StartedAt:       &startNow,
	}

	// start trade session
	_, err = sqlClient.UpsertTradeSession(context.Background(), start)
	assert.Nil(t, err)

	// end trade session
	endNow := time.Now()
	end := models.TradeSessionEntry{
		ID:              id.String(),
		UserID:          "U",
		Algorithm:       "basic",
		CurrencyID:      "C",
		StartingBalance: decimal.NewFromFloat32(1.25),
		EndingBalance:   decimal.NewFromFloat32(1.50),
		StartedAt:       &startNow,
		EndedAt:         &endNow,
	}

	_, err = sqlClient.UpsertTradeSession(context.Background(), end)
	assert.Nil(t, err)
}

func TestSQLClient_SelectTradeSession(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`
		INSERT INTO users (
			id,
			first_name,
			last_name
		)
		VALUES 
			('U', 'brandon', 'brombacher');

		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('C', 'eth', 'eth coin');

		INSERT INTO trade_sessions (
			id, 
			user_id,
			algorithm,
			currency_id,
			starting_balance,
			ending_balance,
			started_at,
			ended_at
		)
		VALUES
			('T', 'U', 'BASIC', 'C', 1.00, 1.25, '2022-08-27 04:08:08.35889+00', '2022-08-27 04:08:08.361871+00'),		
			('TT', 'U', 'BASIC', 'C', 1.25, 1.50, '2022-09-27 04:08:08.35889+00', '2022-09-27 04:08:08.361871+00');
			
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	// upset entry
	_, err = sqlClient.SelectTradeSession(context.Background(), "U", "TT")
	assert.Nil(t, err)
}

func TestSQLClient_SelectTradeSessions(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`
		INSERT INTO users (
			id,
			first_name,
			last_name
		)
		VALUES 
			('U', 'brandon', 'brombacher');

		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('C', 'eth', 'eth coin');

		INSERT INTO trade_sessions (
			id, 
			user_id,
			algorithm,
			currency_id,
			starting_balance,
			ending_balance,
			started_at,
			ended_at
		)
		VALUES
			('T', 'U', 'BASIC', 'C', 1.00, 1.25, '2022-08-27 04:08:08.35889+00', '2022-08-27 04:08:08.361871+00'),		
			('TT', 'U', 'BASIC', 'C', 1.25, 1.50, '2022-09-27 04:08:08.35889+00', '2022-09-27 04:08:08.361871+00');
			
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	// upset entry
	params := models.GetTradeSessionsParams{
		UserID: "U",
	}
	results, err := sqlClient.SelectTradeSessions(context.Background(), params)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(results))
}
