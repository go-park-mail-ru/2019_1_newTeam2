package wshub

import (
	"sync"
)

type ClientsMap struct {
	clients map[int]*Client
	mx sync.RWMutex
}

func (m *ClientsMap) Load(key int) (*Client, bool){
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.clients[key]
	return val, ok
}

func (m * ClientsMap) Store(client *Client) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.clients[client.ID] = client
}

func (m *ClientsMap) Delete(key int) {
	m.mx.Lock()
	defer m.mx.Unlock()
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