package interfaces

import "github.com/user/2019_1_newTeam2/shared/models"

type DBChatInterface interface {
	ChatManager
}

type ChatManager interface {
	GetMessagesBroadcast(page int, rowsNum int) (models.Messages, error)
	AddMessage(UserID int, message string) error
}
