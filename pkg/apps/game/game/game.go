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
	Rooms    	map[string]*room.Room
	MaxRooms 	int
	Register 	chan *websocket.Conn
	DBUser   	string
	DBPassUser	string
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

func NewGame(DBUser string, DBPassUser string) *Game{
	return &Game{
		Rooms:    make(map[string]*room.Room),
		MaxRooms: 2,
		Register: make(chan *websocket.Conn),
		DBUser: DBUser,
		DBPassUser: DBPassUser,
	}
}

func (game *Game) Run() {
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

	room := room.New(game.DBUser, game.DBPassUser)
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
	log.Printf("player %s joined room %s", player.ID, room.ID)

	go player.Listen()

	if len(room.Players) == 1 {
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

	NewTask := r.CreateTask()
	r.Answer = NewTask
	log.Printf("Answer: %s\n", r.Answer)
	r.Broadcast <- &models.GameMessage{Type: "Task", Payload: NewTask}

	for {
		<-r.Ticker.C
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
		r.Broadcast <- &models.GameMessage{Type: "Leaderboard", Payload: state}
	}
}