package game

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/apps/game/room"
)

type Game struct {
	Rooms    map[string]*room.Room
	MaxRooms int
	Register chan *websocket.Conn
}

func (game *Game) DeleteEmptyRoom(ERoom *room.Room) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	for name, room := range game.Rooms {
		if room == ERoom {
			delete(game.Rooms, name)
			log.Printf("room %s removed", ERoom.ID)
		}
	}
	mutex.Unlock()
}

func NewGame() *Game{
	return &Game{
		Rooms:    make(map[string]*room.Room),
		MaxRooms: 2,
		Register: make(chan *websocket.Conn),
	}
}

func (game *Game) Run() {
	//go game.DeleteEmptyRoom()
	for {
		conn := <-game.Register
		log.Printf("got new connection")
		game.ProcessConn(conn)
	}
}

func (game *Game) FindRoom(player *room.Player) *room.Room {
	for _, room := range game.Rooms {
		if len(room.Players) < room.MaxPlayers {
			room.Players[player.ID] = player
			player.Room = room
			return room
		}
	}

	if len(game.Rooms) >= game.MaxRooms {
		return nil
	}

	room := room.New()
	room.Players[player.ID] = player
	player.Room = room
	game.Rooms[room.ID] = room
	go room.ListenToPlayers()
	log.Printf("room %s created", room.ID)

	return room
}

func (game *Game) ProcessConn(conn *websocket.Conn) {
	id := uuid.NewV4().String()
	player := &room.Player{
		Conn: conn,
		ID:   id,
	}

	room := game.FindRoom(player)
	if room == nil {
		return
	}
	//room.Players[player.ID] = player
	//player.Room = room
	log.Printf("player %s joined room %s", player.ID, room.ID)
	go player.Listen()

	if len(room.Players) == room.MaxPlayers {
		go game.RoomRun(room)
	}

}

func (game *Game) RoomRun(r *room.Room) {
	r.Ticker = time.NewTicker(time.Second)
	go r.RunBroadcast()

	players := []models.PlayerData{}
	for _, p := range r.Players {
		players = append(players, p.Data)
	}
	state := &models.State{
		Players: players,
	}

	r.Broadcast <- &models.GameMessage{Type: "SIGNAL_START_THE_GAME", Payload: state}

	for {
		<-r.Ticker.C
		log.Printf("room %s tick with %d players", r.ID, len(r.Players))
		if len(r.Players) == 0 {
			game.DeleteEmptyRoom(r)
			return
		}
		players := []models.PlayerData{}
		for _, p := range r.Players {
			players = append(players, p.Data)
		}
		state := &models.State{
			Players: players,
		}
		r.Broadcast <- &models.GameMessage{Type: "SIGNAL_NEW_GAME_STATE", Payload: state}
	}
}