CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE monitored_ips (
    id SERIAL PRIMARY KEY,
    address TEXT UNIQUE NOT NULL
);

CREATE TABLE log_events (
    id SERIAL PRIMARY KEY,
    ip_address TEXT NOT NULL,
    event TEXT NOT NULL CHECK (event IN ('online', 'offline')),
    timestamp TIMESTAMP DEFAULT NOW()
);
