package interfaces

import "github.com/user/2019_1_newTeam2/models"

type DBInterface interface {
	UserManager
	LanguageManager
	CardManager
	DictionaryManager
	// interface for other managers
	// method to connect to db
	// method to query sql
	// method to exec sql
	// ..
}

type CardManager interface {
	GetCards(dictId int, page int, rowsNum int) ([]models.Card, bool, error)
	GetCard(cardId int) (models.Card, bool, error)
}

type DictionaryManager interface {
	DictionaryDelete(DictID int) error
	DictionaryCreate(UserID int, Name string, Description string, Cards []models.Card) error
	DictionaryUpdate(DictID int, Name string, Description string) error
	GetDicts(userId int, page int, rowsNum int) ([]models.DictionaryInfo, bool, error)
	GetDict(dictId int) (models.DictionaryInfo, bool, error)
}


type LanguageManager interface {
	GetLangs() (models.Language, bool, error)
}

type UserManager interface {
	Login(username string, password string, secret []byte) (string, string, error)
	GetUserByID(userID int) (models.User, bool, error)
	UserRegistration(username string, email string,
		password string, langid int, pronounceOn int) (bool, error)
	DeleteUserById(userID int) (bool, error)
	GetUsers(page int, rowsNum int) ([]models.UserTableElem, bool, error)
	AddImage(path string, userID int) error
	UpdateUserById(userID int, username string, email string,
		password string, langid int, pronounceOn int) (bool, error)
}