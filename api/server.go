package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func typeRequest(sPath string) (string, string) {
	sPath = path.Clean("/" + sPath)
	slash := strings.Index(sPath[1:], "/")
	if slash < 0 {
		return sPath[1:], "/"
	}
	slash++
	return sPath[1:slash], sPath[slash:]
}

type Server struct {
	// users storage.userStorage
	users userStorage
}

func initServer() *Server {
	data := make([]User, 10)
	for i := 0; i < 10; i++ {
		data[i].ID = i + 1
		data[i].Username = "test_user_1"
		data[i].Email = "kek@lol.kl"
		data[i].Password = "pass"
		data[i].langID = 0
		data[i].pronounceON = 1
	}
	return &Server{userStorage{data}}
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHTTP")
	var head string
	head, r.URL.Path = typeRequest(r.URL.Path)

	// w.Write(loginFormTmpl)

	// switch r.Method {
	// case http.MethodGet:
	// 	fmt.Println("get")
	// 	server.getUser(w, r)
	// case http.MethodPost:
	// 	fmt.Println("post")
	// 	server.createUser(w, r)
	// default:
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// }

	switch head {
	case "user":
		fmt.Println("user---> ", head)
		server.userAPI(w, r)
	default:
		fmt.Println("default")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (server *Server) userAPI(w http.ResponseWriter, r *http.Request) {
	head, _ := typeRequest(r.URL.Path)
	fmt.Println("userAPI: ", head)

	switch r.Method {
	case http.MethodGet:
		fmt.Println("get")
		server.getUser(w, r)
	case http.MethodPost:
		fmt.Println("post")
		server.createUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
