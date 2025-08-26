package jwttoken

import (
	"database/sql"
	"errors"
	"notes_service/middlewares"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Token struct{}

// Проверяет email & password
func (t *Token) LoginValidate(checkEmail, checkPswd string, db *sql.DB) bool {
	var userPassword string

	err := db.QueryRow(`SELECT hashed_password FROM users WHERE email = $1`, checkEmail).Scan(&userPassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(checkPswd))
	return err == nil
}

// Создает JWT-токен
func (t *Token) CreateToken(userId int) (string, error) {
	claims := &middlewares.JwtCustomClaims{
		User_id: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Create token with claims

	encodedToken, err := token.SignedString([]byte("secret")) // Generate encoded token and send it as response.
	if err != nil {
		return "", errors.New("error occured while encode the token")
	}

	return encodedToken, nil
}
