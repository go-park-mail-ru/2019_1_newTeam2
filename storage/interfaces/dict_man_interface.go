package interfaces

import "github.com/user/2019_1_newTeam2/models"

type DictManager interface {
	GetDicts(userId int, page int, rowsNum int) ([]models.DictReduced, bool, error)
	GetDict(dictId int) (models.DictReduced, bool, error)
}
