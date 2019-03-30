package storage

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/user/2019_1_newTeam2/models"
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

func (db *Database) UserRegistration(username string, email string,
	password string, langid int, pronounceOn int) (bool, error) {
	for _, i := range db.UserData {
		if i.Username == username {
			return false, fmt.Errorf("Такой пользователь уже существует")
		}
	}
	id := db.LastUserId
	db.Logger.Log(db.LastUserId)
	_, err := HashPassword(password)
	if err != nil {
		return false, fmt.Errorf("hash error")
	}
	db.UserData[id] = models.User{id, username, email, password, langid, pronounceOn, 0, "files/avatars/shrek.jpg"}
	return true, nil
}
