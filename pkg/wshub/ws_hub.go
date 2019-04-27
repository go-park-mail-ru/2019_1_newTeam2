package wshub

type WSHub struct {
	clients    map[int]*Client
	register   chan *Client
	unregister chan int
	sendTo     chan *Message
	broadcast  chan *Message
}

func (h *WSHub) SendToCl(mes *Message) {
	cl, ok := h.clients[mes.ID]
	if ok {
		cl.sendChan <- mes
	}
}

func (h *WSHub) SendAll(mes *Message) {
	for clientID := range h.clients {
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
			_, ok := h.clients[client.ID]
			if !ok {
				h.clients[client.ID] = client	// possible bugs
			}
			h.clients[client.ID] = client

		case clID := <-h.unregister:
			_, ok := h.clients[clID]
			if ok {
				delete(h.clients, clID)
			}

		case mes := <-h.sendTo:
			h.SendToCl(mes)

		case mes := <-h.broadcast:
			h.SendAll(mes)
		}
	}
}

func NewWSHub() *WSHub {
	hub := new(WSHub)
	hub.unregister = make(chan int)
	hub.register = make(chan *Client)
	hub.sendTo = make(chan *Message)
	hub.clients = make(map[int]*Client)
	hub.broadcast = make(chan *Message)
	return hub
}
