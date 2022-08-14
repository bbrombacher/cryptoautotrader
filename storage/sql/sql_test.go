package sql_test

import (
	"bbrombacher/cryptoautotrader/storage/models"
	"bbrombacher/cryptoautotrader/storage/sql"
	"context"
	goSql "database/sql"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var validDb = GetTestDB()

func GetTestDB() *sqlx.DB {
	goSql, _ := goSql.Open(
		"postgres",
		"postgres://pguser:pgpass@localhost:9001/robot-transact?sslmode=disable",
	)
	db := sqlx.NewDb(goSql, "postgres")
	return db
}

func userTestCleanup() {
	if _, err := validDb.Exec("DELETE FROM users"); err != nil {
		log.Fatal("could not clear test db")
	}

	if _, err := validDb.Exec(`ALTER TABLE users ALTER COLUMN cursor_id RESTART SET START 1`); err != nil {
		log.Fatal("could not reset curosr_id sequence")
	}
}

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
