package users

type User struct {
	Id       int `json:"id"`
	Email    string
	Password string
}
