package wshub

import (
	"github.com/gorilla/websocket"
	"time"
)

const  writeWait = 10 * time.Second

type Client struct {
	Conn *websocket.Conn
	ID int
}

func (cl *Client) SendMes(mes *Message, unsubscribe chan int) {
	_ = cl.Conn.SetWriteDeadline(time.Now().Add(writeWait))
	err := cl.Conn.WriteJSON(mes.Data)
	if err != nil {
		unsubscribe <- cl.ID
	}
}
