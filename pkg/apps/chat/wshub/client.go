package wshub

import (
	"github.com/gorilla/websocket"
	"github.com/user/2019_1_newTeam2/models"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	maxMessageSize = 512
	pingPeriod     = (pongWait * 9) / 10
)

type Client struct {
	Conn     *websocket.Conn
	ID       int
	hub      *WSHub
	sendChan chan interface{}
}

func (cl *Client) ReadFromInet() {
	defer func() {
		cl.hub.unregister <- cl.ID
		_ = cl.Conn.Close()
	}()

	cl.Conn.SetReadLimit(maxMessageSize)
	_ = cl.Conn.SetReadDeadline(time.Now().Add(pongWait))
	cl.Conn.SetPongHandler(func(appData string) error {
		_ = cl.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		mes := models.Message{}
		err := cl.Conn.ReadJSON(&mes)
		if err != nil {
			break
		}
		cl.hub.broadcast <- &mes
	}
}

func (cl *Client) WriteToInet() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = cl.Conn.Close()
		cl.hub.unregister <- cl.ID
	}()
	for {
		select {
		case mes, ok := <-cl.sendChan:
			_ = cl.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_= cl.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := cl.Conn.WriteJSON(mes)
			if err != nil {
				return
			}
		case <-ticker.C:
			_ = cl.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := cl.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func NewClient(id int, ws *websocket.Conn, hub *WSHub) *Client {
	cl := new(Client)
	cl.ID = id
	cl.hub = hub
	cl.sendChan = make(chan interface{})
	cl.Conn = ws
	return cl
}
