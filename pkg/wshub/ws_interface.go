package wshub

import (
	"net/http"
)

type IWSCommunicator interface {
	AddClient(w http.ResponseWriter, r *http.Request, id int) error
	SendToClient(mes *Message)
	DeleteClient(ID int)
}
