package interfaces

import "github.com/user/2019_1_newTeam2/models"

type DBGameInterface interface {
	GameManager
}

type GameManager interface {
	CreateTask() (models.GameQuestion, error)
}
