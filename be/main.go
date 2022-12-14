package main

import (
	"bbrombacher/cryptoautotrader/be/coinbase"
	"bbrombacher/cryptoautotrader/be/controllers/balance"
	"bbrombacher/cryptoautotrader/be/controllers/currencies"
	"bbrombacher/cryptoautotrader/be/controllers/trade_sessions"
	"bbrombacher/cryptoautotrader/be/controllers/transactions"
	users "bbrombacher/cryptoautotrader/be/controllers/users"
	"bbrombacher/cryptoautotrader/be/storage"
	"bbrombacher/cryptoautotrader/be/storage/sql"
	"bbrombacher/cryptoautotrader/be/tradebot"
	"context"
	goSql "database/sql"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var localDB = "postgres://pguser:pgpass@localhost:9001/robot-transact?sslmode=disable"
var coinbaseURL = "wss://ws-feed.pro.coinbase.com"

func main() {

	dbURL := localDB
	port := "8080"
	if os.Getenv("ENV") == "server" {
		dbURL = os.Getenv("DB_URL")
		port = os.Getenv("PORT")
	}

	log.Println("ENV", os.Getenv("ENV"))
	log.Println("DB_URL", dbURL)
	log.Println("PORT", port)

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
	taskMap := &sync.Map{}

	// setup controllers
	userController := users.Controller{
		StorageClient: storageClient,
	}

	currencyController := currencies.Controller{
		StorageClient: storageClient,
	}

	balanceController := balance.Controller{
		StorageClient: storageClient,
	}

	transactionController := transactions.Controller{
		StorageClient: storageClient,
	}

	tradeBot := tradebot.Bot{
		Coinbase:      *coinbase.New(coinbaseURL),
		StorageClient: storageClient,
		Tasks:         taskMap,
	}

	err = tradeBot.StartOrphanedTradeSessions()
	if err != nil {
		log.Fatalln(err.Error())
	}

	tradeSessionsController := trade_sessions.Controller{
		StorageClient: storageClient,
		Bot:           tradeBot,
	}

	// setup router
	r := mux.NewRouter()

	// register routes
	userController.Register(r)
	currencyController.Register(r)
	balanceController.Register(r)
	transactionController.Register(r)
	tradeSessionsController.Register(r)

	// set middleware
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(accessControlAllowOriginsMW)

	// listen
	http.ListenAndServe(":"+port, r)
}

func accessControlAllowOriginsMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "x-user-id, Content-Type")

		log.Println("allowed", w.Header().Get("Access-Control-Allow-Methods"))
		if r.Method == http.MethodOptions {
			log.Println("options method")
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
