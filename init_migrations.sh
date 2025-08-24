#!/bin/bash

# Настройки подключения к базе данных
DB_NAME="Notes_Service"
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_HOST="localhost"
DB_PORT="5432"

# Функция для выполнения SQL-файла
execute_sql_file() {
    local sql_file=$1
    echo "Executing $sql_file..."
    PGPASSWORD=$DB_PASSWORD psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -d "$DB_NAME" -f "$sql_file"
    if [ $? -ne 0 ]; then
        echo "Error executing $sql_file. Exiting..."
        exit 1
    fi
    echo "$sql_file executed successfully."
}

# Создание таблиц
execute_sql_file "migrations/drop.sql"
execute_sql_file "migrations/000_migrations.sql"
execute_sql_file "migrations/001_created_users.sql"
execute_sql_file "migrations/002_created_notes.sql"

echo "All SQL files executed successfully."