package interfaces

// import "github.com/user/2019_1_newTeam2/models"

type DBChatInterface interface {
	ChatManager
}

type ChatManager interface {
	GetMessages(dictId int, page int, rowsNum int) ([]models.Card, bool, error)
	AddMessage(cardId int) (models.Card, bool, error)
}
