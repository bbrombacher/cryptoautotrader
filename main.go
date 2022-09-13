package main

import (
	"bbrombacher/cryptoautotrader/controllers/currencies"
	users "bbrombacher/cryptoautotrader/controllers/users"
	"bbrombacher/cryptoautotrader/storage"
	"bbrombacher/cryptoautotrader/storage/sql"
	"context"
	goSql "database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var localDB = "postgres://pguser:pgpass@localhost:9001/robot-transact?sslmode=disable"

func main() {

	dbURL := localDB
	if os.Getenv("ENV") == "server" {
		dbURL = os.Getenv("DB_URL")
	}

	// set up db
	sqldb, err := goSql.Open(
		"postgres",
		dbURL,
	)
	if err != nil {
		log.Println("error opening sql", err.Error())
	}
	db := sqlx.NewDb(sqldb, "postgres")
	sqlClient, err := sql.NewSQLClient(context.Background(), db)
	if err != nil {
		log.Println("could not start sql client: ", err.Error())
	}
	storageClient := storage.NewStorageClient(sqlClient)

	// setup controllers
	userController := users.Controller{
		StorageClient: storageClient,
	}

	currencyController := currencies.Controller{
		StorageClient: storageClient,
	}

	// setup router
	r := mux.NewRouter()
	userController.Register(r)
	currencyController.Register(r)
	http.ListenAndServe(":8080", r)
}
