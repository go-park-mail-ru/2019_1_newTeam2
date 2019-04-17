package storage

import (
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

func (db *Database) Login(username string, password string, secret []byte) (string, string, error) {
	user, check, _ := db.CheckUserByUsername(username)
	if !check {
		return "", "", fmt.Errorf("Неверное имя пользователя")
	}

	kek := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if kek != nil {
		return "", "", fmt.Errorf("Неверный пароль")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       int64(user.ID),
	})
	str, _ := token.SignedString(secret)
	return str, strconv.Itoa(user.ID), nil
}
