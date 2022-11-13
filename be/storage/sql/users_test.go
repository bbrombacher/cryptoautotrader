package sql_test

import (
	"bbrombacher/cryptoautotrader/be/storage/models"
	"bbrombacher/cryptoautotrader/be/storage/sql"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSQLClient_SelectUser(t *testing.T) {
	testCleanup()
	defer testCleanup()

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

	result, err := sqlClient.SelectUser(context.Background(), models.GetUserParams{ID: "one"})
	assert.Nil(t, err)

	assert.Equal(t, "brandon", result.FirstName)
	assert.Equal(t, "brombacher", result.LastName)
}

func TestSQLClient_InsertUser(t *testing.T) {
	testCleanup()
	defer testCleanup()

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

func TestSQLClient_UpdateUser(t *testing.T) {
	testCleanup()
	defer testCleanup()

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

	newEntry := models.UserEntry{
		ID:        "one",
		FirstName: "not brandon",
		LastName:  "brombacher",
	}

	updateColumns := []string{"first_name", "last_name"}

	// insert entry
	result, err := sqlClient.UpdateUser(context.Background(), newEntry, updateColumns)
	assert.Nil(t, err)

	assert.Equal(t, "not brandon", result.FirstName)
	assert.Equal(t, "brombacher", result.LastName)
}

func TestSQLClient_DeleteUser(t *testing.T) {
	testCleanup()
	defer testCleanup()

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
