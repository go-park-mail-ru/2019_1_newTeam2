package database

import (
	"github.com/user/2019_1_newTeam2/models"
)

type UserManager interface {
	Login(username string, password string, secret []byte) (string, string, error)
	GetUserByID(userID int) (models.User, bool, error)
	UserRegistration(username string, email string,
		password string, langid int, pronounceOn int) (bool, error)
	DeleteUserById(userID int) (bool, error)
	GetUsers(page int, rowsNum int) ([]models.UserTableElem, error)
	AddImage(path string, userID int) error
	UpdateUserById(userID int, username string, email string,
		password string, langid int, pronounceOn int) (bool, error)
	IncUserLastID() // remove in future
}

type DBInterface interface {
	UserManager
	// interface for other managers
	// method to connect to db
	// method to query sql
	// method to exec sql
	// ..
}
