package service

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
