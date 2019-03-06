package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	. "./storage"
)

func TypeRequest(sPath string) (string, string) {
	sPath = path.Clean("/" + sPath)
	slash := strings.Index(sPath[1:], "/")
	if slash < 0 {
		return sPath[1:], "/"
	}
	slash++
	return sPath[1:slash], sPath[slash:]
}

type Server struct {
	Users UserStorage
	// users userStorage
}

func InitServer() *Server {
	data := make([]User, 10)
	for i := 0; i < 10; i++ {
		data[i].ID = i + 1
		data[i].Username = "test_user_1"
		data[i].Email = "kek@lol.kl"
		data[i].Password = "pass"
		data[i].LangID = 0
		data[i].PronounceON = 1
	}
	return &Server{UserStorage{data}}
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHTTP")
	var head string
	head, r.URL.Path = TypeRequest(r.URL.Path)

	switch head {
	case "user":
		fmt.Println("user---> ", head)
		server.UserAPI(w, r)
	default:
		fmt.Println("default")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (server *Server) UserAPI(w http.ResponseWriter, r *http.Request) {
	head, _ := TypeRequest(r.URL.Path)
	fmt.Println("userAPI: ", head)

	switch r.Method {
	case http.MethodGet:
		fmt.Println("get")
		server.GetUser(w, r)
	case http.MethodPost:
		fmt.Println("post")
		server.CreateUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
