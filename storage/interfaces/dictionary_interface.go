package interfaces

import "github.com/user/2019_1_newTeam2/models"

type DictionaryManager interface {
	DictionaryDelete(DictID int) error
	DictionaryCreate(UserID int, Name string, Description string, Cards []models.Card) error
	DictionaryUpdate(DictID int, Name string, Description string) error
	GetDicts(userId int, page int, rowsNum int) ([]models.DictionaryInfo, bool, error)
	GetDict(dictId int) (models.DictionaryInfo, bool, error)
}
