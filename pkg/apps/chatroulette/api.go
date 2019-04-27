package chatroulette

import (
	"log"
	"net/http"
)

func (server *ChatServer) CreateChat(w http.ResponseWriter, r *http.Request) {
	log.Println("create chat")
}

func (server *ChatServer) GetHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("get history")
}