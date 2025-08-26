package service

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Log in
func (s *Service) GetToken(c echo.Context) error {
	var user User

	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")

	if ok := s.loginValidate(user.Email, user.Password); !ok {
		s.logger.Error("unauthorized: wrong email or password")
		return echo.ErrUnauthorized
	}

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

// Проверяет email & password
func (s *Service) loginValidate(checkEmail, checkPswd string) bool {
	flag := s.token.LoginValidate(checkEmail, checkPswd, s.db)

	return flag
}

// Создает JWT-Token
func (s *Service) createToken(user User) (string, error) {
	id, err := s.usersRepo.GetUserId(user.Email)
	if err != nil {
		s.logger.Error(err)
		return "", err
	}
	user.Id = id

	token, err := s.token.CreateToken(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}
