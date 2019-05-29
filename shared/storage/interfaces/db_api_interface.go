package interfaces

import "github.com/user/2019_1_newTeam2/shared/models"

type DBInterface interface {
	UserManager
	LanguageManager
	CardManager
	DictionaryManager
	SinglePlayerGameManager
}

type CardManager interface {
	GetCards(dictId int, page int, rowsNum int) (models.Cards, bool, error)
	GetCard(cardId int) (models.Card, bool, error)
	SetCardToDictionary(userId int, dictID int, card models.Card) error
	DeleteCardInDictionary(userId int, cardID int, dictionaryID int) error
	GetCardsForGame(dictId int, cardsNum int) (models.GameWords, bool, error)
}

type DictionaryManager interface {
	DictionaryDelete(DictID int) error
	DictionaryCreate(UserID int, Name string, Description string, Cards []models.Card) (models.DictionaryInfoPrivilege, error)
	DictionaryUpdate(DictID int, Name string, Description string) error
	GetDicts(userId int, page int, rowsNum int) (models.DictInfos, bool, error)
	GetDict(dictId int) (models.DictionaryInfoPrivilege, bool, error)
	BorrowDictById(dictId int, thiefId int) (int, models.DictionaryInfo, error)
	FillDictionaryFromXLSX(userId int, dictId int, pathToFile string) error
}

type LanguageManager interface {
	GetLangs() (models.Langs, bool, error)
	GetLangByName(LangName string) (models.Language, error)
}

type UserManager interface {
	Login(username string, password string, secret []byte) (string, string, error)
	GetUserByID(userID int) (models.User, bool, error)
	UserRegistration(username string, email string,
		password string, langid int, pronounceOn int) (bool, error)
	DeleteUserById(userID int) (bool, error)
	GetUsers(page int, rowsNum int) (models.TableUsers, bool, error)
	AddImage(path string, userID int) error
	UpdateUserById(userID int, username string, email string,
		langid int, pronounceOn int) (bool, error)
}

type SinglePlayerGameManager interface {
	UpdateFrequencies(results models.GameResults) (error, bool)
}
