package service

import (
	"errors"
	"fmt"
	"net/http"

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
