package wshub

import (
	"github.com/gorilla/websocket"
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

func (com *WSCommunicator) AddClient(w http.ResponseWriter, r *http.Request, id int) error {
	ws, err := upgrader.Upgrade(w.(http.ResponseWriter), r, nil)

	if err != nil {
		return err
	}
	cl := &Client{ID: id, Conn: ws}
	com.hub.register <- cl
	go cl.ReadFromInet()
	go cl.WriteToInet()
	return nil
}

func (com *WSCommunicator) SendToClient(mes *Message) {
	com.hub.sendTo <- mes
}
func (com *WSCommunicator) DeleteClient(ID int) {
	com.hub.unregister <- ID
}

func NewWSCommunicator() *WSCommunicator {
	com := new(WSCommunicator)
	com.hub = NewWSHub()
	go com.hub.Run()
	return com
}
