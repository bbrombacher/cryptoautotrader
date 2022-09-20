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

	if _, err := validDb.Exec("DELETE FROM transactions"); err != nil {
		log.Fatalln("could not clear test db - transactions", err)
	}

	if _, err := validDb.Exec(`ALTER TABLE currencies ALTER COLUMN cursor_id RESTART SET START 1`); err != nil {
		log.Fatalln("could not reset curosr_id sequence - transactions", err)
	}

	if _, err := validDb.Exec("DELETE FROM balance"); err != nil {
		log.Fatalln("could not clear test db - balance", err)
	}

	if _, err := validDb.Exec("DELETE FROM users"); err != nil {
		log.Fatalln("could not clear test db - users", err)
	}

	if _, err := validDb.Exec(`ALTER TABLE users ALTER COLUMN cursor_id RESTART SET START 1`); err != nil {
		log.Fatalln("could not reset curosr_id sequence - users", err)
	}

	if _, err := validDb.Exec("DELETE FROM currencies"); err != nil {
		log.Fatalln("could not clear test db - currencies", err)
	}

	if _, err := validDb.Exec(`ALTER TABLE currencies ALTER COLUMN cursor_id RESTART SET START 1`); err != nil {
		log.Fatalln("could not reset curosr_id sequence - currencies", err)
	}

}
