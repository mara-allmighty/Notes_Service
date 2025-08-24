package service

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	User_id int    `json:"user_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
