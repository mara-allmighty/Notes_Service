package users

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

// CREATE user
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

// GET JWT-token
func (ur *UsersRepo) IsUserAuthSuccessful(checkEmail, checkPswd string) bool {
	// ~Log in
	var user User

	// достаем хеш пароля
	err := ur.db.QueryRow(`SELECT hashed_password FROM users WHERE email = $1`, checkEmail).Scan(&user.Password)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(checkPswd))
	return err == nil
}

// Хешировать -- protected method
func (ur *UsersRepo) createUserHashedPswd(unhashed_pswd string) string {
	var user User

	byte_hashed_pswd, err := bcrypt.GenerateFromPassword([]byte(unhashed_pswd), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	user.Password = string(byte_hashed_pswd)
	return user.Password
}
