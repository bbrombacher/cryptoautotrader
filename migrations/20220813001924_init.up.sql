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
    user_id text not null,
    currency_id TEXT not null,
    amount DECIMAL,
    updated_at TIMESTAMPTZ default now(),
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id),
    CONSTRAINT fk_currency_id
        FOREIGN KEY(currency_id)
        REFERENCES currencies(id),
    CONSTRAINT "balance_pk" PRIMARY KEY ("user_id", "currency_id")
);

CREATE TABLE IF NOT EXISTS transactions(
    id TEXT PRIMARY KEY,
    user_id TEXT,
    use_currency_id TEXT,
    get_currency_id TEXT,
    transaction_type TEXT,
    amount DECIMAL,
    price DECIMAL,
    cursor_id INT GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ default now(),
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id),
    CONSTRAINT fk_use_currency_id
        FOREIGN KEY(use_currency_id)
        REFERENCES currencies(id),
    CONSTRAINT fk_get_currency_id
        FOREIGN KEY(get_currency_id)
        REFERENCES currencies(id)
);

CREATE TABLE IF NOT EXISTS trade_sessions(
    id TEXT PRIMARY KEY,
    user_id TEXT,
    algorithm TEXT,
    currency_id TEXT,
    starting_balance decimal,
    ending_balance decimal,
    cursor_id INT GENERATED ALWAYS AS IDENTITY,
    started_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(id),
    CONSTRAINT fk_currency_id
        FOREIGN KEY(currency_id)
        REFERENCES currencies(id)
);

CREATE TABLE IF NOT EXISTS transaction_sessions_map(
    trade_session_id TEXT,
    transaction_id TEXT,
    CONSTRAINT "trade_session_map_pk" PRIMARY KEY ("trade_session_id", "transaction_id")
);