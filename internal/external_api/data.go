package externalapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetQuote() map[string]string {
	// URL API
	url := "https://favqs.com/api/qotd"

	// Отправляем GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка при запросе: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка HTTP: %d\n", resp.StatusCode)
		return nil
	}

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка чтения ответа: %v\n", err)
		return nil
	}

	// Парсим JSON
	var quoteResp QuoteResponse
	err = json.Unmarshal(body, &quoteResp)
	if err != nil {
		fmt.Printf("Ошибка парсинга JSON: %v\n", err)
		return nil
	}

	// Выводим цитату
	quoteData := map[string]string{
		"author": quoteResp.Quote.Author,
		"quote":  quoteResp.Quote.Body,
	}

	return quoteData
}
