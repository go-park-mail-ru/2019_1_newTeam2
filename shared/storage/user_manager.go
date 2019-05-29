package storage

import (
	"fmt"

	"github.com/user/2019_1_newTeam2/shared/models"
)

func (db *Database) CheckUserByUsername(username string) (models.User, bool, error) {
	result := db.Conn.QueryRow(GetUserByUsernameQuery, username)

	user := new(models.User)

	err := result.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.LangID, &user.PronounceON, &user.Score, &user.AvatarPath)

	if err != nil {
		return models.User{}, false, nil
	}

	if user.Username == "" {
		return models.User{}, false, nil
	}
	return *user, true, nil
}

func (db *Database) GetUserByID(userID int) (models.User, bool, error) {
	row := db.Conn.QueryRow(GetUserByIDQuery, userID)
	user := new(models.User)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.LangID, &user.PronounceON, &user.Score, &user.AvatarPath)
	if err != nil {
		return models.User{}, false, err
	}
	return *user, true, nil
}

func (db *Database) DeleteUserById(userID int) (bool, error) {
	_, DeleteErr := db.Conn.Exec(
		DeleteUserQuery,
		userID,
	)

	if DeleteErr != nil {
		db.Logger.Log("user not create")
		return false, fmt.Errorf("user not create")
	}

	return true, nil
}

func (db *Database) UpdateUserById(userID int, username string, email string,
	/*password string,*/ langid int, pronounceOn int) (bool, error) {
	_, check, _ := db.GetUserByID(userID)
	if !check {
		db.Logger.Log("Такого пользователя не существует")
		return false, fmt.Errorf("Такого пользователя не существует")
	}

	_, UpdateErr := db.Conn.Exec(
		UpdateUserQuery,
		username,
		email,
		langid,
		pronounceOn,
		userID,
	)

	if UpdateErr != nil {
		db.Logger.Log("user not update")
		fmt.Println(UpdateErr)
		return false, fmt.Errorf("user not update")
	}

	return true, nil
}

func (db *Database) GetUsers(page int, rowsNum int) ([]models.UserTableElem, bool, error) {
	usersPage := make([]models.UserTableElem, 0)
	db.Logger.Log(page, rowsNum)
	offset := (page - 1) * rowsNum
	db.Logger.Log(offset)
	rows, err := db.Conn.Query(UsersPaginate, rowsNum, offset)
	if err != nil {
		return usersPage, false, err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		i++
		user := models.UserTableElem{}
		err := rows.Scan(&user.Username, &user.Score)
		if err != nil {
			return usersPage, false, err
		}
		usersPage = append(usersPage, user)
	}
	if i == 0 {
		return usersPage, false, nil
	}
	return usersPage, true, nil
}

func (db *Database) AddImage(path string, userID int) error {
	_, check, _ := db.GetUserByID(userID)
	if !check {
		db.Logger.Log("Такого пользователя не существует")
		return fmt.Errorf("Такого пользователя не существует")
	}

	_, _ = db.Conn.Exec("USE wordtrainer")
	_, CreateErr := db.Conn.Exec(
		UpdateImagePathUserQuery,
		path,
		userID,
	)

	if CreateErr != nil {
		fmt.Println(CreateErr)
		db.Logger.Log("image is not added")
		return fmt.Errorf("image is not added")
	}
	return nil
}

func (db *Database) UserRegistration(username string, email string,
	password string, langid int, pronounceOn int) (bool, error) {
	// here should possibly be transaction
	_, check, _ := db.CheckUserByUsername(username)
	if check {
		return false, fmt.Errorf("Такой пользователь уже существует")
	}
	hashPassword, err := HashPassword(password)
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
		"",
	)
	if CreateErr != nil {
		db.Logger.Log(CreateErr)
		return false, fmt.Errorf("CreateErr: user not create")
	}
	return true, nil
}
