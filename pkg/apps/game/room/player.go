package room

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/user/2019_1_newTeam2/models"
)

type Player struct {
	ID   string
	Room *Room
	Conn *websocket.Conn
	Data models.PlayerData
}

type IncomingMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
	Player  *Player         `json:"-"`
}

func (p *Player) Listen() {
	p.Room.Logger.Log("start listening messages from player ", p.ID)

	for {
		m := &IncomingMessage{}

		err := p.Conn.ReadJSON(m)
		if websocket.IsUnexpectedCloseError(err) {
			p.Room.Logger.Log("player ", p.ID, " was disconnected")
			p.Room.Unregister <- p
			return
		}

		m.Player = p
		p.Room.Message <- m
	}
}

func (p *Player) Send(s *models.GameMessage) {
	err := p.Conn.WriteJSON(s)
	if err != nil {
		p.Room.Logger.Log("cannot send state to client: ", err)
	}
}