package main

import (
	"fmt"
	"net/http"
	"strings"

	. "./storage"
)

func TypeRequest(sPath string) (string, string) {
	slash := strings.Index(sPath[1:], "/")
	if slash < 0 {
		return sPath[1:], "/"
	}
	slash++
	return sPath[1:slash], sPath[slash:]
}

type Server struct {
	Users UserStorage
}

func InitServer() *Server {
	data := make(map[int]User)
	for i := 0; i < 10; i++ {
		data[i] = User{i, "test_user", "kek@lol.kl", "pass", 0, 1}
	}
	return &Server{UserStorage{data}}
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHTTP")
	var head string
	head, r.URL.Path = TypeRequest(r.URL.Path)

	switch head {
	case "users":
		fmt.Println("users---> ", head)
		server.UsersAPI(w, r)
	default:
		fmt.Println("default")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
