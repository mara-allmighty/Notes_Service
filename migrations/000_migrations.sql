CREATE TABLE IF NOT EXISTS migrations (
    id SERIAL PRIMARY KEY,                  -- уникальный ID
    filename TEXT NOT NULL UNIQUE,          -- имя файла миграции (001_create_users.sql и т.п.)
    applied_at TIMESTAMP DEFAULT NOW()      -- когда применили
);
