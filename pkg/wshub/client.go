package wshub

import (
	"github.com/gorilla/websocket"
	"github.com/user/2019_1_newTeam2/models"
	"time"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	maxMessageSize = 512
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	Conn *websocket.Conn
	ID   int
	hub WSHub
	sendChan chan interface{}
}

/*func (cl *Client) SendMes(mes *Message, unsubscribe chan int) {
	_ = cl.Conn.SetWriteDeadline(time.Now().Add(writeWait))
	err := cl.Conn.WriteJSON(mes.Data)
	if err != nil {
		unsubscribe <- cl.ID
	}
}*/

func (cl *Client) ReadFromInet() {
	defer func() {
		cl.hub.unregister <- cl.ID
		_  = cl.Conn.Close()
	}()
	cl.Conn.SetReadLimit(maxMessageSize)
	_ = cl.Conn.SetReadDeadline(time.Now().Add(pongWait))
	cl.Conn.SetPongHandler(func(appData string) error {
		_ = cl.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, _, err := cl.Conn.ReadMessage()
		// here we can work with message like we want(second arg, add in future)
		if err != nil {
			break
		}
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
		case mes, ok := <- cl.sendChan:
			_ = cl.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				err := cl.Conn.WriteJSON(models.Error{"connection was closed"})
				if err != nil {
					return
				}
			}
			err := cl.Conn.WriteJSON(mes)
			if err != nil {
				return
			}
		case <- ticker.C:
			_ = cl.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := cl.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}