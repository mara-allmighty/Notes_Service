package users

import (
	"database/sql"
	"errors"
	"notes_service/middlewares"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

// Sign up
func (ur *UsersRepo) SignUp(email, password string) error {
	var user User

	err := ur.db.QueryRow(`SELECT email FROM users WHERE email = $1`, email).Scan(&user.Email)
	if err == nil {
		return errors.New("user already exist")
	}

	user.Email = email
	user.Password = ur.createUserHashedPswd(password)

	_, err = ur.db.Exec(`INSERT INTO users (email, hashed_password) VALUES ($1, $2)`, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// Хеширует пароль [protected method].
func (ur *UsersRepo) createUserHashedPswd(unhashed_pswd string) string {
	var user User

	byte_hashed_pswd, err := bcrypt.GenerateFromPassword([]byte(unhashed_pswd), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	user.Password = string(byte_hashed_pswd)
	return user.Password
}

// Возвращает идентификатор текущего пользователя. Обеспечивает доступ только к своим заметкам.
func (ur *UsersRepo) GetCurrentUser(c echo.Context) int {
	// Получаем токен из контекста
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JwtCustomClaims) // Преобразуем в собственный тип
	user_id := claims.User_id

	return user_id
}

// Только для получения user_id во время логина.
func (ur *UsersRepo) GetUserId(email string) (int, error) {
	var user User

	err := ur.db.QueryRow(`SELECT id FROM users WHERE email = $1`, email).Scan(&user.Id)
	if err != nil {
		return -1, err
	}

	return user.Id, nil
}
