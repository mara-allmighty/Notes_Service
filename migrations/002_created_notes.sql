CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY, 
    user_id INT NOT NULL REFERENCES users(id),
    title TEXT NOT NULL UNIQUE,
    body TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO migrations (filename) VALUES ('002_created_notes.sql');