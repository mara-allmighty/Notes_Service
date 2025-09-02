package service

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type JwtCustomClaims struct {
	User_id int `json:"user_id"`
	jwt.RegisteredClaims
}
