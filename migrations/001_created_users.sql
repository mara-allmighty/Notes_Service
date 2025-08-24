CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL
);

INSERT INTO migrations (filename) VALUES ('001_created_users.sql');