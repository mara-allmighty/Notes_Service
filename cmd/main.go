package main

import (
	"fmt"
	"log"
)

func main() {
	db, err := PostgresConnection()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()
	fmt.Println("hello")
}
