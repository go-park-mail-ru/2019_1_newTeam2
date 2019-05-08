package interfaces

import "github.com/user/2019_1_newTeam2/models"

type DBMGRInterface interface {
	MGRManager
}

type MGRManager interface {
	UpdateScore(username string, added_points int) (models.UserScore, error)
}
