CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY, 
    user_id INT REFERENCES users(id),
    title TEXT NOT NULL,
    body TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO migrations (filename) VALUES ('002_created_notes.sql');