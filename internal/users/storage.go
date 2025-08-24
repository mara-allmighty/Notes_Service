package users

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Email    string
	Password string
}

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
