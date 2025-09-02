package externalapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetQuote() (map[string]string, error) {
	// URL API
	url := "https://favqs.com/api/qotd"

	// Отправляем GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Парсим JSON
	var quoteResp QuoteResponse
	err = json.Unmarshal(body, &quoteResp)
	if err != nil {
		return nil, err
	}

	// Выводим цитату
	quoteData := map[string]string{
		"author": quoteResp.Quote.Author,
		"quote":  quoteResp.Quote.Body,
	}

	return quoteData, nil
}
