CREATE TABLE IF NOT EXISTS users(
    id TEXT PRIMARY KEY,
    first_name TEXT, 
    last_name TEXT,
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now(),
    CONSTRAINT "PK_user_id" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS balance(
    id TEXT PRIMARY KEY,
    user_id text not null,
    currency TEXT,
    amount DECIMAL,
    updated_at TIMESTAMPTZ default now(),
    CONSTRAINT fk_user
      FOREIGN KEY(user_id)
);

CREATE TABLE IF NOT EXISTS trade_sessions(
    id TEXT PRIMARY KEY,
    user_id TEXT,
    algorithm TEXT,
    currency TEXT,
    starting_balance decimal,
    ending_balance decimal,
    started_at TIMESTAMPTZ default now(),
    ended_at TIMESTAMPTZ default now(),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
);

CREATE TALBE IF NOT EXISTS transactions(
    id TEXT PRIMARY KEY,
    trade_session_id TEXT,
    user_id TEXT,
    currency TEXT,
    transaction_type TEXT,
    amount DECIMAL,
    price DECIMAL,
    created_at TIMESTAMPTZ default now(),
    CONSTRAINT fk_user_id 
        FOREIGN KEY(user_id)
        FOREIGN KEY(trade_session_id)
);