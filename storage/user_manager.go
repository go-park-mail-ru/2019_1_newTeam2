package storage

import (
	"fmt"

	"github.com/user/2019_1_newTeam2/models"
)

func (db *Database) IncUserLastID() {
	db.LastUserId++
}

func (db *Database) GetUserByID(userID int) (models.User, bool, error) {
	for _, i := range db.UserData {
		if i.ID == userID {
			return i, true, nil
		}
	}
	return models.User{}, false, nil
}

func (db *Database) DeleteUserById(userID int) (bool, error) {
	delete(db.UserData, userID)
	return true, nil
}

func (db *Database) UpdateUserById(userID int, username string, email string,
	password string, langid int, pronounceOn int) (bool, error) {
	db.UserData[userID] = models.User{userID, username, email, password, langid, pronounceOn, db.UserData[userID].Score, db.UserData[userID].AvatarPath}
	return true, nil
}

func (db *Database) GetUsers(page int, rowsNum int) ([]models.UserTableElem, error) {
	usersPage := make([]models.UserTableElem, 0)
	db.Logger.Log(page, rowsNum)
	offset := (page - 1) * rowsNum
	db.Logger.Log(offset)
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
	db.Logger.Log(path)
	db.UserData[userID] = user
	return nil
}
