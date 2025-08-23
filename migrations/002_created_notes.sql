CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY, 
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    body TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO migrations (filename) VALUES ('002_created_notes.sql');