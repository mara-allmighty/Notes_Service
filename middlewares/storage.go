package middlewares

import "github.com/golang-jwt/jwt/v5"

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Username string `json:"name"`
	jwt.RegisteredClaims
}
