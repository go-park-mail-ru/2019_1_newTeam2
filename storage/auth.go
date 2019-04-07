package storage

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

func (db *Database) Login(username string, password string, secret []byte) (string, string, error) {
	for _, i := range db.UserData {
		if i.Username != username {
			continue
		}
		_, err := HashPassword(password)
		if err != nil {
			return "", "", fmt.Errorf("hash error")
		}
		if password == i.Password {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": username,
				"id":       int64(i.ID),
			})
			str, _ := token.SignedString(secret)
			return str, strconv.Itoa(i.ID), nil
		} else {
			return "", "", fmt.Errorf("Неверный пароль")
		}

	}
	return "", "", fmt.Errorf("Неверное имя пользователя")
}
