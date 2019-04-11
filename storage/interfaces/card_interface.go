package interfaces

import "github.com/user/2019_1_newTeam2/models"

type CardManager interface {
	GetCards(dictId int, page int, rowsNum int) ([]models.Card, bool, error)
	GetCard(cardId int) (models.Card, bool, error)
}
