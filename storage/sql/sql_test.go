package sql_test

import (
	goSql "database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func testCleanup() {
	if _, err := validDb.Exec("DELETE FROM balance"); err != nil {
		log.Fatal("could not clear test db - balance")
	}

	if _, err := validDb.Exec("DELETE FROM users"); err != nil {
		log.Fatal("could not clear test db - users")
	}

	if _, err := validDb.Exec(`ALTER TABLE users ALTER COLUMN cursor_id RESTART SET START 1`); err != nil {
		log.Fatal("could not reset curosr_id sequence - users")
	}

	if _, err := validDb.Exec("DELETE FROM currencies"); err != nil {
		log.Fatal("could not clear test db - currencies")
	}

	if _, err := validDb.Exec(`ALTER TABLE currencies ALTER COLUMN cursor_id RESTART SET START 1`); err != nil {
		log.Fatal("could not reset curosr_id sequence - currencies")
	}
}
