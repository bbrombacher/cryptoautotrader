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

func TestSQLClient_InsertUser(t *testing.T) {
	userTestCleanup()
	defer userTestCleanup()

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	// set up entry
	id := uuid.New()
	entry := models.UserEntry{
		ID:        id.String(),
		FirstName: "brandon",
		LastName:  "brombacher",
	}

	// setup expected result
	expectedResult := models.UserEntry{
		ID:        id.String(),
		FirstName: "brandon",
		LastName:  "brombacher",
		CursorID:  1,
	}

	// insert entry
	result, err := sqlClient.InsertUser(context.Background(), entry)
	assert.Nil(t, err)

	assert.Equal(t, expectedResult.FirstName, result.FirstName)
	assert.Equal(t, expectedResult.LastName, result.LastName)
	assert.Equal(t, expectedResult.CursorID, result.CursorID)
}

func TestSQLClient_DeleteUser(t *testing.T) {
	userTestCleanup()
	defer userTestCleanup()

	if _, err := validDb.Exec(`
		INSERT INTO users (
			id,
			first_name,
			last_name
		)
		VALUES 
			('one', 'brandon', 'brombacher');
	
	`); err != nil {
		log.Fatal("could not seed test db ", err)
	}

	// setup db
	sqlClient, err := sql.NewSQLClient(context.Background(), validDb)
	assert.Nil(t, err)

	// insert entry
	err = sqlClient.DeleteUser(context.Background(), "one")
	assert.Nil(t, err)
}