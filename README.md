# cryptoautotrader
auto trader bot beep boop beep

postgres://pguser:pgpass@localhost:9001/robot-transact?sslmode=disable

TODO:
- Trade Sessions
- - For reach transaction, insert entry to transaction_sessions_map (with transaction_id and trade_session_id)
- - Create logic of actual trade session (i.e. bot go brrrr)
- - - bot package should get the storage client so it can perform trades
- - Start/Stop Trade Session API
