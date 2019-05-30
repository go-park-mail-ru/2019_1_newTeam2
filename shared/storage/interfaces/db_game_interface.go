package interfaces

import "github.com/user/2019_1_newTeam2/shared/models"

type DBGameInterface interface {
	MultiPlayerGameManager
}

type MultiPlayerGameManager interface {
	CreateTask() (models.GameQuestion, error)
	GetWordsForDemo(wordsNum int) (models.GameQuestions, error)
}
