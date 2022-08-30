CREATE TABLE IF NOT EXISTS users(
    id TEXT PRIMARY KEY,
    first_name TEXT, 
    last_name TEXT,
    cursor_id INT GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now()
);

CREATE TABLE IF NOT EXISTS currencies(
    id TEXT PRIMARY KEY,
    name text NOT NULL UNIQUE,
    description TEXT,
    cursor_id INT GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now()
);

CREATE TABLE IF NOT EXISTS balance(
    id TEXT PRIMARY KEY,
    user_id text not null,
    currency_id TEXT,
    amount DECIMAL,
    updated_at TIMESTAMPTZ default now(),
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id),
    CONSTRAINT fk_currency
        FOREIGN KEY(currency_id)
        REFERENCES currencies(id)
);

CREATE TABLE IF NOT EXISTS trade_sessions(
    id TEXT PRIMARY KEY,
    user_id TEXT,
    algorithm TEXT,
    currency TEXT,
    starting_balance decimal,
    ending_balance decimal,
    cursor_id INT GENERATED ALWAYS AS IDENTITY,
    started_at TIMESTAMPTZ default now(),
    ended_at TIMESTAMPTZ default now(),
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS transactions(
    id TEXT PRIMARY KEY,
    trade_session_id TEXT,
    user_id TEXT,
    currency TEXT,
    transaction_type TEXT,
    amount DECIMAL,
    price DECIMAL,
    cursor_id INT GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ default now(),
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id),
    CONSTRAINT fk_trade_session_id
        FOREIGN KEY(trade_session_id)
        REFERENCES trade_sessions(id)    
);