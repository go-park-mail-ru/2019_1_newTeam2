package wshub

import (
	"net/http"
	"github.com/user/2019_1_newTeam2/models"
)

type IWSCommunicator interface {
	AddClient(w http.ResponseWriter, r *http.Request, id int) error
	SendToClient(mes *models.Message)
	DeleteClient(ID int)
}
