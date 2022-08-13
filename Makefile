LOCAL_PSQL_URL=postgres://pguser:pgpass@localhost:9001/robot-transact?sslmode=disable

db.setup:
	docker-compose up -d db
	sleep 5
	make migrate.up
db.reset:
	make db.delete
	make db.setup
db.delete:
	docker stop robot-transactions
	docker rm robot-transactions
migrate.up:
	migrate -path migrations -database $(LOCAL_PSQL_URL) -verbose up