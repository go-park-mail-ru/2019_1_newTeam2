package storage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/user/2019_1_newTeam2/models"
)

func (db *Database) CheckUserByUsername(username string) (models.User, bool, error) {
	results, err := db.Conn.Query(GetUserByUsernameQuery, username)

	if err != nil {
		return models.User{}, false, err
	}

	user := new(models.User)
	for results.Next() {
		err = results.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.LangID, &user.PronounceON, &user.Score, &user.AvatarPath)
		if err != nil {
			return models.User{}, false, nil
		}
	}
	if user.Username == "" {
		return models.User{}, false, nil
	}
	return *user, true, nil
}

func (db *Database) GetUserByID(userID int) (models.User, bool, error) {
	results, err := db.Conn.Query(GetUserByIDQuery, userID)

	if err != nil {
		return models.User{}, false, err
	}

	user := new(models.User)
	for results.Next() {
		err = results.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.LangID, &user.PronounceON, &user.Score, &user.AvatarPath)
		if err != nil {
			return models.User{}, false, nil
		}
	}
	return *user, true, nil
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

func (db *Database) UserRegistration(username string, email string,
	password string, langid int, pronounceOn int) (bool, error) {

	_, check, _ := db.CheckUserByUsername(username)
	if check {
		return false, fmt.Errorf("Такой пользователь уже существует")
	}

	db.Logger.Log(db.LastUserId)
	hashPassword, err := HashPassword(password)
	kekPassword, err := HashPassword(password)
	fmt.Println("-- password: ", password)
	fmt.Println("-- hash password: ", hashPassword)
	fmt.Println("-- kek password: ", kekPassword)

	if err != nil {
		return false, fmt.Errorf("hash error")
	}

	_, CreateErr := db.Conn.Exec(
		AddUserQuery,
		username,
		email,
		hashPassword,
		langid,
		pronounceOn,
		0,
		"files/avatars/shrek.jpg",
	)

	if CreateErr != nil {
		return false, fmt.Errorf("user not create")
	}

	return true, nil
}
