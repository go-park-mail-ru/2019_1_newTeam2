package wshub

type ClientsMap struct {
	clients map[int]*Client
}

func (m *ClientsMap) Load(key int) (*Client, bool) {
	val, ok := m.clients[key]
	return val, ok
}

func (m *ClientsMap) Store(client *Client) {
	m.clients[client.ID] = client
}

func (m *ClientsMap) Delete(key int) {
	_, ok := m.clients[key]
	if ok {
		delete(m.clients, key)
	}
}

func NewClientsMap() *ClientsMap {
	clMap := new(ClientsMap)
	clMap.clients = make(map[int]*Client)
	return clMap
}
