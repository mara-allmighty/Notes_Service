package users

import (
	"database/sql"

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

	user.Email = email
	user.Password = ur.createUserHashedPswd(password)

	_, err := ur.db.Exec(`INSERT INTO users (email, hashed_password) VALUES ($1, $2)`, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// Log in
func (ur *UsersRepo) LogIn(checkEmail, checkPswd string) bool {
	var userPassword string

	err := ur.db.QueryRow(`SELECT hashed_password FROM users WHERE email = $1`, checkEmail).Scan(&userPassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(checkPswd))
	return err == nil
}

// Хеширует пароль
func (ur *UsersRepo) createUserHashedPswd(unhashedPassword string) string {
	var user User

	byteHashedPassword, err := bcrypt.GenerateFromPassword([]byte(unhashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	user.Password = string(byteHashedPassword)
	return user.Password
}

// Для получения user_id во время логина, т.к. из токена Id получить нельзя из-за отсутствия токена
func (ur *UsersRepo) GetUserId(email string) (int, error) {
	var user User

	err := ur.db.QueryRow(`SELECT id FROM users WHERE email = $1`, email).Scan(&user.Id)
	if err != nil {
		return -1, err
	}

	return user.Id, nil
}
