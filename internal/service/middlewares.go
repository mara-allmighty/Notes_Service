package service

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// Middleware для проверки токена
func AuthMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims { return new(JwtCustomClaims) },
		SigningKey:    []byte("secret"),
	}
	return echojwt.WithConfig(config)
}
