package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Sign up
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

// Log in
func (s *Service) LogIn(c echo.Context) error {
	var user User

	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")

	// валидация User-а
	if ok := s.usersRepo.LogIn(user.Email, user.Password); !ok {
		s.logger.Error("unauthorized: wrong email or password")
		return echo.ErrUnauthorized
	}

	// создаем jwt-token
	token, err := s.createToken(user)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.New("token create error"))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user":  user,
		"token": token,
	})
}

// Обеспечивает доступ только к своим заметкам
func (s *Service) GetCurrentUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)

	claims := user.Claims.(*JwtCustomClaims)
	user_id := claims.User_id

	return user_id
}

func (s *Service) createToken(user User) (string, error) {
	id, err := s.usersRepo.GetUserId(user.Email)
	if err != nil {
		s.logger.Error(err)
		return "", err
	}
	user.Id = id

	claims := &JwtCustomClaims{
		User_id: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", errors.New("error occured while encode the token")
	}

	return encodedToken, nil
}
