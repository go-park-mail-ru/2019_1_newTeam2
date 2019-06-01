package wshub

import (
	"github.com/user/2019_1_newTeam2/shared/models"
	"net/http"
)

type IWSCommunicator interface {
	AddClient(w http.ResponseWriter, r *http.Request, id int, username string) error
	SendToClient(mes *models.Message)
	DeleteClient(ID int)
}
