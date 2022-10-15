package sql_test

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"bbrombacher/cryptoautotrader/storage/sql"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSQLClient_InsertTransaction(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`

		INSERT INTO users (
			id,
			first_name,
			last_name
		)
		VALUES 
			('user_one', 'brandon', 'brombacher');

		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('currency_one', 'eth', 'eth coin');
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	id := uuid.New()
	entry := models.TransactionEntry{
		ID:              id.String(),
		UserID:          "user_one",
		UseCurrencyID:   "currency_one",
		GetCurrencyID:   "currency_two",
		TransactionType: "buy",
		Amount:          decimal.NewFromFloat(1.00),
		Price:           decimal.NewFromFloat(1.25),
	}
	result, err := sqlClient.InsertTransaction(context.Background(), entry)
	assert.Nil(t, err)

	assert.Equal(t, id.String(), result.ID)
	assert.Equal(t, "currency_one", result.UseCurrencyID)
	assert.Equal(t, "user_one", result.UserID)
	assert.Equal(t, "buy", result.TransactionType)

	expectedAmount := decimal.NewFromFloat32(1.00)
	assert.True(t, expectedAmount.Equal(result.Amount))

	expectedPrice := decimal.NewFromFloat32(1.25)
	assert.True(t, expectedPrice.Equal(result.Price))
}

func TestSQLClient_SelectTransaction(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`

		INSERT INTO users (
			id,
			first_name,
			last_name
		)
		VALUES 
			('user_one', 'brandon', 'brombacher');

		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('currency_one', 'eth', 'eth coin');

			INSERT INTO transactions (
				id,
				user_id,
				currency_id,
				transaction_type,
				amount,
				price
			)
			VALUES 
				('transaction_one', 'user_one', 'currency_one', 'buy', 1.00, 1.25);

	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	result, err := sqlClient.SelectTransaction(context.Background(), "transaction_one", "user_one")
	assert.Nil(t, err)

	assert.Equal(t, "transaction_one", result.ID)
	assert.Equal(t, "currency_one", result.UseCurrencyID)
	assert.Equal(t, "user_one", result.UserID)
	assert.Equal(t, "buy", result.TransactionType)

	expectedAmount := decimal.NewFromFloat32(1.00)
	assert.True(t, expectedAmount.Equal(result.Amount))

	expectedPrice := decimal.NewFromFloat32(1.25)
	assert.True(t, expectedPrice.Equal(result.Price))
}

func TestSQLClient_SelectTransactions(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`

		INSERT INTO users (
			id,
			first_name,
			last_name
		)
		VALUES 
			('user_one', 'brandon', 'brombacher');

		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('currency_one', 'eth', 'eth coin');

			INSERT INTO transactions (
				id,
				user_id,
				currency_id,
				transaction_type,
				amount,
				price
			)
			VALUES 
				('transaction_one', 'user_one', 'currency_one', 'buy', 1.00, 1.25),
				('transaction_two', 'user_one', 'currency_one', 'buy', 1.20, 1.35);

	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	params := models.GetTransactionsParams{UserID: "user_one"}
	result, err := sqlClient.SelectTransactions(context.Background(), params)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))

	transactionOne := result[0]
	assert.Equal(t, "transaction_one", transactionOne.ID)
	assert.Equal(t, "currency_one", transactionOne.UseCurrencyID)
	assert.Equal(t, "user_one", transactionOne.UserID)
	assert.Equal(t, "buy", transactionOne.TransactionType)

	expectedAmount := decimal.NewFromFloat32(1.00)
	assert.True(t, expectedAmount.Equal(transactionOne.Amount))

	expectedPrice := decimal.NewFromFloat32(1.25)
	assert.True(t, expectedPrice.Equal(transactionOne.Price))
}
