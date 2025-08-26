package service

import (
	"errors"
	"fmt"
	"net/http"
	"notes_service/middlewares"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// POST /signup
func (s *Service) SignUp(c echo.Context) error {
	var user User

	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")

	err := s.usersRepo.SignUp(user.Email, user.Password)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusInternalServerError, "user already exist")
	}

	user.Id, err = s.usersRepo.GetUserId(user.Email)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("user %s created successfully!", user.Email),
		"info":    user,
	})
}

// GET /login
func (s *Service) LogIn(c echo.Context) error {
	var user User

	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")

	if ok := s.usersRepo.IsUserAuthSuccessful(user.Email, user.Password); !ok {
		s.logger.Error("unauthorized")
		return echo.ErrUnauthorized
	}

	id, err := s.usersRepo.GetUserId(user.Email)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.New("user not found"))
	}
	user.Id = id

	claims := &middlewares.JwtCustomClaims{ // Set custom claims
		User_id: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Create token with claims

	t, err := token.SignedString([]byte("secret")) // Generate encoded token and send it as response.
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user":  user,
		"token": t,
	})
}
