package game

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/apps/game/room"
	"github.com/user/2019_1_newTeam2/pkg/apps/mgr"
)

type GameRegister struct {
	Conn     *websocket.Conn
	Username string
}

type Game struct {
	Rooms       map[string]*room.Room
	MaxRooms    int
	Register    chan *GameRegister
	DBUser      string
	DBPassUser  string
	ScoreClient mgr.UserScoreUpdaterClient
}

func (game *Game) DeleteEmptyRoom(ERoom *room.Room) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	for name, room := range game.Rooms {
		if room == ERoom {
			delete(game.Rooms, name)
			RoomCountMetric.Dec()
			ERoom.Logger.Log("room removed")
		}
	}
	mutex.Unlock()
}

func NewGame(DBUser string, DBPassUser string, scoreClient mgr.UserScoreUpdaterClient) *Game {
	return &Game{
		Rooms:       make(map[string]*room.Room),
		MaxRooms:    2,
		Register:    make(chan *GameRegister),
		DBUser:      DBUser,
		DBPassUser:  DBPassUser,
		ScoreClient: scoreClient,
	}
}

func (game *Game) Run() {
	for {
		conn := <-game.Register
		game.ProcessConn(conn)
	}
}

func (game *Game) ProcessConn(conn *GameRegister) error {
	//id := uuid.NewV4().String()
	player := &room.Player{
		Conn: conn.Conn,
		ID:   conn.Username,
		Data: models.PlayerData{conn.Username, 0},
	}
	r := game.FindRoom(player)
	if r == nil {
		return fmt.Errorf("Can`t create and found room")
	}
	r.Logger.Log("player ", player.Data.Username, " joined room ", r.ID)
	room.PlayerCountMetric.Inc()
	go player.Listen()
	if len(r.Players) <= 1 {
		go game.RoomRun(r)
	} else {
		player.Send(&models.GameMessage{Type: "Task", Payload: player.Room.Answer})
	}
	return nil
}

func (game *Game) FindRoom(player *room.Player) *room.Room {
	for _, room := range game.Rooms {
		if len(room.Players) < room.MaxPlayers {
			if room.FindPlayer(player) == true {
				return nil
			}
			room.Players[player.ID] = player
			player.Room = room
			return room
		}
	}
	if len(game.Rooms) >= game.MaxRooms {
		return nil
	}
	room := room.New(game.DBUser, game.DBPassUser, game.ScoreClient)
	room.Players[player.ID] = player
	player.Room = room
	game.Rooms[room.ID] = room
	go room.ListenToPlayers()
	RoomCountMetric.Inc()
	room.Logger.Log("room created")
	return room
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
