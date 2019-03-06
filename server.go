package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	. "./storage"
)

func TypeRequest(url string) (string, string) {
	separatorPosition := strings.Index(url[1:], "/")
	if separatorPosition < 0 {
		return url[1:], "/"
	}
	separatorPosition++
	return url[1:separatorPosition], url[separatorPosition:]
}

type Server struct {
	Users UserStorage
}

func InitServer() *Server {
	data := make(map[int]User)
	for i := 0; i < 10; i++ {
		data[i] = User{i, "test_user_" + strconv.Itoa(i), "kek@lol.kl", "pass", 0, 1, 0}
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
