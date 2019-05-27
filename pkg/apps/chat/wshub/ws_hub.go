package wshub

import (
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/storage"
	"github.com/user/2019_1_newTeam2/storage/interfaces"
	"log"
)

type WSHub struct {
	clients    map[int]*Client
	register   chan *Client
	unregister chan int
	sendTo     chan *models.Message
	broadcast  chan *models.Message
	DB         interfaces.DBChatInterface
}

func (h *WSHub) SendToCl(mes *models.Message) {
	cl, ok := h.clients[mes.ID]
	if ok {
		cl.sendChan <- mes
	}
}

func (h *WSHub) SendAll(mes *models.Message) {
	for clientID := range h.clients {
		log.Printf("ClientID=%v\nmes.ID=%v\n", clientID, mes.ID)
		if clientID == mes.ID {
			continue
		}
		client := h.clients[clientID]
		select {
		case client.sendChan <- mes:
		default:
			close(client.sendChan)
			delete(h.clients, clientID)
		}
	}
}

func (h *WSHub) Run() {
	for {
		select {
		case client := <-h.register:
			/*_, ok := h.clients[client.ID]
			if !ok {
				h.clients[client.ID] = client	// possible bugs
			}*/
			h.clients[client.ID] = client
			h.SendToCl(&models.Message{client.ID, "Welcome to Word chat!)"})

		case clID := <-h.unregister:
			_, ok := h.clients[clID]
			if ok {
				delete(h.clients, clID)
			}

		case mes := <-h.sendTo:
			h.SendToCl(mes)

		case mes := <-h.broadcast:
			err := h.DB.AddMessage(mes.ID, mes.Data)
			if err != nil {
				log.Println(err)
			}
			log.Println(mes.Data)
			h.SendAll(mes)
		}
	}
}

func NewWSHub(host string, username string, pass string) *WSHub {
	hub := new(WSHub)
	hub.unregister = make(chan int)
	hub.register = make(chan *Client)
	hub.sendTo = make(chan *models.Message)
	hub.clients = make(map[int]*Client)
	hub.broadcast = make(chan *models.Message)
	newDB, err := storage.NewDataBase(host, username, pass)
	if err != nil {
		return nil
	}
	hub.DB = newDB
	return hub
}
