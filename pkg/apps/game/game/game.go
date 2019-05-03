package game

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"github.com/user/2019_1_newTeam2/pkg/apps/game/room"
)

type Game struct {
	Rooms    map[string]*room.Room
	MaxRooms int
	Register chan *websocket.Conn
}

func NewGame() *Game {
	return &Game{
		Rooms:    make(map[string]*room.Room),
		MaxRooms: 2,
		Register: make(chan *websocket.Conn),
	}
}

func (game *Game) Run() {
	for {
		conn := <-game.Register
		log.Printf("got new connection")
		game.ProcessConn(conn)
	}
}

func (game *Game) FindRoom() *room.Room {
	for _, room := range game.Rooms {
		if len(room.Players) < room.MaxPlayers {
			return room
		}
	}

	if len(game.Rooms) >= game.MaxRooms {
		return nil
	}

	room := room.New()
	go room.ListenToPlayers()
	game.Rooms[room.ID] = room
	log.Printf("room %s created", room.ID)

	return room
}

func (game *Game) ProcessConn(conn *websocket.Conn) {
	id := uuid.NewV4().String()
	player := &room.Player{
		Conn: conn,
		ID:   id,
	}

	room := game.FindRoom()
	if room == nil {
		return
	}
	room.Players[player.ID] = player
	player.Room = room
	log.Printf("player %s joined room %s", player.ID, room.ID)
	go player.Listen()

	if len(room.Players) == room.MaxPlayers {
		go room.Run()
	}

}
