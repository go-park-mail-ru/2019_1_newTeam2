package room

import (
	uuid "github.com/satori/go.uuid"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"github.com/user/2019_1_newTeam2/storage"
	"github.com/user/2019_1_newTeam2/storage/interfaces"
	"os"
	"time"
)

type Room struct {
	ID         string
	Ticker     *time.Ticker
	Logger     logger.LoggerInterface
	DB         interfaces.DBGameInterface
	Players    map[string]*Player
	MaxPlayers int
	Register   chan *Player
	Unregister chan *Player
	Message    chan *IncomingMessage
	Broadcast  chan *models.GameMessage
	Answer	   models.GameQuestion
}

func New(DBUser string, DBPassUser string) *Room {
	id := uuid.NewV4().String()

	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix("ROOM ("+ id + ") LOG: ")

	newDB, err := storage.NewDataBase(DBUser, DBPassUser)
	if err != nil {
		logger.Log("new room: ", err)
		return nil
	}

	return &Room{
		ID:         id,
		MaxPlayers: 3,
		Players:    make(map[string]*Player),
		Logger:		logger,
		DB:			newDB,
		Register:   make(chan *Player),
		Unregister: make(chan *Player),
		Broadcast:  make(chan *models.GameMessage),
		Message:    make(chan *IncomingMessage),
	}
}

func (r *Room) ListenToPlayers() {
	for {
		select {
		case m := <-r.Message:
			switch m.Type {
				case "ANSWER":
					answer := string(m.Payload)[1:len(string(m.Payload))-1]
					if answer == r.Answer.Answer {
						m.Player.Data.Score += 1
						NewTask := r.CreateTask()
						r.Answer = NewTask
						r.Broadcast <- &models.GameMessage{Type: "Task", Payload: NewTask}
					}
			}
		case p := <-r.Unregister:
			delete(r.Players, p.ID)
			r.Logger.Log("player was deleted from room ", r.ID)
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
