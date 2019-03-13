package database

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"github.com/user/2019_1_newTeam2/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (db *Database) IncUserLastID() {
	db.LastUserId++
}

func (db *Database) Login(username string, password string, secret []byte) (string, string, error) {
	for _, i := range db.UserData {
		if i.Username == username {
			// h := sha256.New()
			// h.Write([]byte(password))
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
	}
	return "", "", fmt.Errorf("Неверное имя пользователя")
}

func (db *Database) GetUserByID(userID int) (models.User, bool, error) {
	for _, i := range db.UserData {
		if i.ID == userID {
			return i, true, nil
		}
	}
	return models.User{}, false, nil
}

func (db *Database) UserRegistration(username string, email string,
	password string, langid int, pronounceOn int) (bool, error) {
	for _, i := range db.UserData {
		if i.Username == username {
			return false, fmt.Errorf("Такой пользователь уже существует")
		}
	}
	id := db.LastUserId
	fmt.Println(db.LastUserId)
	// h := sha256.New()
	// h.Write([]byte(password))
	_, err := HashPassword(password)
	if err != nil {
		return false, fmt.Errorf("hash error")
	}
	db.UserData[id] = models.User{id, username, email, password, langid, pronounceOn, 0, "uploads/avatars/1.jpg"}
	return true, nil
}

func (db *Database) DeleteUserById(userID int) (bool, error) {
	delete(db.UserData, userID)
	return true, nil
}

func (db *Database) UpdateUserById(userID int, username string, email string,
	password string, langid int, pronounceOn int) (bool, error) {
	db.UserData[userID] = models.User{userID, username, email, password, langid, pronounceOn, db.UserData[userID].Score, "uploads/avatars/1.jpg"}
	return true, nil
}

func (db *Database) GetUsers(page int, rowsNum int) ([]models.UserTableElem, error) {
	usersPage := make([]models.UserTableElem, 0)
	fmt.Println(page, rowsNum)
	offset := (page - 1) * rowsNum
	fmt.Println(offset)
	// get data from db, if null is returned
	if false {
		return nil, fmt.Errorf("No such users")
	}
	j := 0
	for _, i := range db.UserData {
		j++
		usersPage = append(usersPage, models.UserTableElem{i.Username, i.Score})
		if j == rowsNum {
			break
		}
	}
	return usersPage, nil
}

func (db *Database) AddImage(path string, userID int) error {
	_, ok := db.UserData[userID]
	if !ok {
		return fmt.Errorf("no such user")
	}
	user := db.UserData[userID]
	user.AvatarPath = path
	fmt.Println(path)
	db.UserData[userID] = user
	return nil
}
