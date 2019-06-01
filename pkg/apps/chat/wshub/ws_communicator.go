package wshub

import (
	"github.com/gorilla/websocket"
	"github.com/user/2019_1_newTeam2/shared/models"
	"net/http"
)

const (
	maxWSMessageSize = 1024 * 1024
)

type WSCommunicator struct {
	hub *WSHub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxWSMessageSize,
	WriteBufferSize: maxWSMessageSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (com *WSCommunicator) AddClient(w http.ResponseWriter, r *http.Request, id int, username string) error {
	ws, err := upgrader.Upgrade(w.(http.ResponseWriter), r, nil)

	if err != nil {
		return err
	}
	//cl := &Client{ID: id, Conn: ws, sendChan: make(chan interface{}), hub: com.hub}
	cl := NewClient(id, username, ws, com.hub)
	com.hub.register <- cl
	go cl.ReadFromInet()
	go cl.WriteToInet()
	return nil
}

func (com *WSCommunicator) SendToClient(mes *models.Message) {
	com.hub.sendTo <- mes
}
func (com *WSCommunicator) DeleteClient(ID int) {
	com.hub.unregister <- ID
}

func NewWSCommunicator(host string, username string, pass string) *WSCommunicator {
	com := new(WSCommunicator)
	com.hub = NewWSHub(host, username, pass)
	go com.hub.Run()
	return com
}
