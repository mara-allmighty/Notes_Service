package middlewares

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	User_id int `json:"user_id"`
	jwt.RegisteredClaims
}
