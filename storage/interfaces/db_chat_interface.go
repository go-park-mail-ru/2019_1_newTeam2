package interfaces

import "github.com/user/2019_1_newTeam2/models"

type DBChatInterface interface {
	ChatManager
}

type ChatManager interface {
	GetMessagesBroadcast(page int, rowsNum int) ([]models.Message, bool, error)
	AddMessage(UserID int, message string) error
}
