package room

import (
	uuid "github.com/satori/go.uuid"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/logger"
	"github.com/user/2019_1_newTeam2/storage"
	"github.com/user/2019_1_newTeam2/storage/interfaces"
	//"encoding/json"
	"log"
	"os"
	"time"
)

type NewPlayer struct {
	Username string `json:"username"`
}

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
	Answer	   string
}

func New(DBUser string, DBPassUser string) *Room {
	id := uuid.NewV4().String()

	logger := new(logger.GoLogger)
	logger.SetOutput(os.Stderr)
	logger.SetPrefix(id + " LOG: ")

	newDB, err := storage.NewDataBase(DBUser, DBPassUser)
	if err != nil {
		logger.Log("new room: ", err)
		return nil
	}

	logger.Log("AuthDB:", DBUser, DBPassUser)

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

//TODO(tsaanstu): в комнате хранится JSON, начальная генерация при создании комнаты, отдавать юзеру при подключении

func (r *Room) ListenToPlayers() {
	for {
		select {
		case m := <-r.Message:
			log.Printf("message from player %s: %v", m.Player.ID, string(m.Payload))
			switch m.Type {
				//  {"type":"ANSWER","payload":"бык"}
				case "ANSWER":
					answer := string(m.Payload)[1:len(string(m.Payload))-1]
					if answer == r.Answer {
						log.Printf("Right!")
						m.Player.Data.Score += 1
						NewTask := r.CreateTask()
						r.Answer = NewTask.Answer
						NewTask.Answer = "LolKek4eburek)"
						r.Broadcast <- &models.GameMessage{Type: "Task", Payload: NewTask}
					}
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