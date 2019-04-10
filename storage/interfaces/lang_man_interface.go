package interfaces

import "github.com/user/2019_1_newTeam2/models"

type LangManager interface {
	GetLangs() (models.Language, bool, error)
}


