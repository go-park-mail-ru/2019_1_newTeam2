package wshub

type WSHub struct {
	clients    *ClientsMap
	register   chan *Client
	unregister chan int
	sendTo     chan *Message
}

func (h *WSHub) SendToCl(mes *Message) {
	cl, ok := h.clients.Load(mes.ID)
	if ok {
		go cl.SendMes(mes, h.unregister)
	}
}

func (h *WSHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients.Store(client)
			break

		case clID := <-h.unregister:
			h.clients.Delete(clID)
			break

		case mes := <-h.sendTo:
			h.SendToCl(mes)
			break
		}
	}
}

func NewWSHub() *WSHub {
	hub := new(WSHub)
	hub.unregister = make(chan int)
	hub.register = make(chan *Client)
	hub.sendTo = make(chan *Message)
	hub.clients = NewClientsMap()
	return hub
}
