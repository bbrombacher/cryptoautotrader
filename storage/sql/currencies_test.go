package sql_test

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"bbrombacher/cryptoautotrader/storage/sql"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSQLClient_SelectCurrency(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`
		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('one', 'eth', 'eth coin');
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	result, err := sqlClient.SelectCurrency(context.Background(), "one")
	assert.Nil(t, err)

	assert.Equal(t, "eth", result.Name)
	assert.Equal(t, "eth coin", result.Description)
}

func TestSQLClient_SelectCurrencies(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`
		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('one', 'eth', 'eth coin'),
			('two', 'bitcoin', 'bit coin');
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	params := models.GetCurrenciesParams{}
	result, err := sqlClient.SelectCurrencies(context.Background(), params)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(result))
}

func TestSQLClient_InsertCurrency(t *testing.T) {
	testCleanup()
	defer testCleanup()

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	// set up entry
	id := uuid.New()
	entry := models.CurrencyEntry{
		ID:          id.String(),
		Name:        "eth",
		Description: "ethcoin",
	}

	// setup expected result
	expectedResult := models.CurrencyEntry{
		ID:          id.String(),
		Name:        "eth",
		Description: "ethcoin",
		CursorID:    1,
	}

	// insert entry
	result, err := sqlClient.InsertCurrency(context.Background(), entry)
	assert.Nil(t, err)

	assert.Equal(t, expectedResult.Name, result.Name)
	assert.Equal(t, expectedResult.Description, result.Description)
	assert.Equal(t, expectedResult.CursorID, result.CursorID)
}

func TestSQLClient_UpdateCurrency(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`
		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('one', 'eth', 'eth coin');
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	newEntry := models.CurrencyEntry{
		ID:          "one",
		Name:        "not eth",
		Description: "not eth coin",
	}

	updateColumns := []string{"name", "description"}

	// insert entry
	result, err := sqlClient.UpdateCurrency(context.Background(), newEntry, updateColumns)
	assert.Nil(t, err)

	assert.Equal(t, "not eth", result.Name)
	assert.Equal(t, "not eth coin", result.Description)
}

func TestSQLClient_DeleteCurrency(t *testing.T) {
	testCleanup()
	defer testCleanup()

	if _, err := validDb.Exec(`
		INSERT INTO currencies (
			id,
			name,
			description
		)
		VALUES 
			('one', 'eth', 'eth coin');
	
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	// insert entry
	err = sqlClient.DeleteCurrency(context.Background(), "one")
	assert.Nil(t, err)
}
