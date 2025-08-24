package notes

import "time"

type Note struct {
	User_id    int       `json:"user_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Created_at time.Time `json:"created_at"`
}
