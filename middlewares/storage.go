package middlewares

import "github.com/golang-jwt/jwt/v5"

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	User_id int `json:"user_id"`
	jwt.RegisteredClaims
}
