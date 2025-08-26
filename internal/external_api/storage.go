package externalapi

// Структура для парсинга ответа от API
type QuoteResponse struct {
	Quote struct {
		ID        int      `json:"id"`
		Author    string   `json:"author"`
		Body      string   `json:"body"`
		CreatedAt string   `json:"created_at"`
		UpdatedAt string   `json:"updated_at"`
		Tags      []string `json:"tags"`
	} `json:"quote"`
}
