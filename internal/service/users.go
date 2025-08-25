package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// POST localhost:8000/signup
func (s *Service) SignUp(c echo.Context) error {
	var user User

	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")

	err := s.usersRepo.SignUp(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "user already exist")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("user %s created successfully!", user.Email),
		"info":    user,
	})
}

// GET localhost:8000/login
func (s *Service) LogIn(c echo.Context) error {
	var user User

	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")

	if ok := s.usersRepo.IsUserAuthSuccessful(user.Email, user.Password); !ok {
		return echo.ErrUnauthorized
	}

	claims := &JwtCustomClaims{ // Set custom claims
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Create token with claims

	t, err := token.SignedString([]byte("secret")) // Generate encoded token and send it as response.
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user":  user,
		"token": t,
	})
}
