package room

import (
	"encoding/json"
	"log"
	"time"

	"github.com/satori/go.uuid"
	"github.com/user/2019_1_newTeam2/models"
)

type NewPlayer struct {
	Username string `json:"username"`
}

type Room struct {
	ID         string
	Ticker     *time.Ticker
	Players    map[string]*Player
	MaxPlayers int
	Register   chan *Player
	Unregister chan *Player
	Message    chan *IncomingMessage
	Broadcast  chan *models.GameMessage
}

func New() *Room {
	id := uuid.NewV4().String()

	return &Room{
		ID:         id,
		MaxPlayers: 2,
		Players:    make(map[string]*Player),
		Register:   make(chan *Player),
		Unregister: make(chan *Player),
		Broadcast:  make(chan *models.GameMessage),
		Message:    make(chan *IncomingMessage),
	}
}

func (r *Room) Run() {
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

func (r *Room) ListenToPlayers() {
	for {
		select {
		case m := <-r.Message:
			log.Printf("message from player %s: %v", m.Player.ID, string(m.Payload))

			switch m.Type {
			case "newPlayer":
				np := &NewPlayer{}
				json.Unmarshal(m.Payload, np)
				m.Player.Data.Username = np.Username
			}

		case p := <-r.Unregister:
			delete(r.Players, p.ID)
			log.Printf("player was deleted from room %s", r.ID)
		}

	}
}

func (r *Room) RunBroadcast() {
	for {
		m := <-r.Broadcast
		for _, p := range r.Players {
			p.Send(m)
		}
	}
}
