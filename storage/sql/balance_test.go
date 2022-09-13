package sql_test

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"bbrombacher/cryptoautotrader/storage/sql"
	"context"
	"log"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSQLClient_UpsertBalance(t *testing.T) {
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

	entry := models.BalanceEntry{
		UserID:     "U",
		CurrencyID: "C",
		Amount:     decimal.NewFromFloat32(1.25),
	}

	// upset entry
	result, err := sqlClient.UpsertBalance(context.Background(), entry)
	assert.Nil(t, err)
	assert.Equal(t, decimal.NewFromFloat32(1.25), result.Amount)
	assert.Equal(t, "U", result.UserID)
	assert.Equal(t, "C", result.CurrencyID)
}

func TestSQLClient_SelectBalance(t *testing.T) {
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

		INSERT INTO balance (
				user_id,
				currency_id,
				amount
			)
			VALUES 
				('U', 'C', 1.00);
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	// upset entry
	result, err := sqlClient.SelectBalance(context.Background(), "U", "C")
	assert.Nil(t, err)
	expectedDec := decimal.NewFromFloat32(1.00)
	assert.True(t, expectedDec.Equal(result.Amount))
	assert.Equal(t, "U", result.UserID)
	assert.Equal(t, "C", result.CurrencyID)
}

func TestSQLClient_SelectBulkBalance(t *testing.T) {
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
			('C', 'eth', 'eth coin'), 
			('B', 'usd', 'usd coin');

		INSERT INTO balance (
				user_id,
				currency_id,
				amount
			)
			VALUES 
				('U', 'C', 1.00), 
				('U', 'B', 1.23);
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	// upset entry
	result, err := sqlClient.SelectBulkBalance(context.Background(), "U")
	assert.Nil(t, err)

	assert.Equal(t, 2, len(result))

	expectedDec := decimal.NewFromFloat32(1.00)
	assert.True(t, expectedDec.Equal(result[0].Amount))
	assert.Equal(t, "U", result[0].UserID)
	assert.Equal(t, "C", result[0].CurrencyID)
}
