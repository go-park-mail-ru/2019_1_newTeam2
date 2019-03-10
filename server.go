package main

import (
	"net/http"
	"strconv"
	"strings"

	. "./storage"
	"github.com/gorilla/mux"
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
	Users  UserStorage
	Router *mux.Router
}

func InitServer() *Server {
	data := make(map[int]User)
	LastId := 10
	for i := 0; i < LastId; i++ {

		data[i] = User{i, "test_user_" + strconv.Itoa(i), "kek@lol.kl", "pass", 0, 1, 0}
	}
	router := mux.NewRouter()

	server := new(Server)

	router.HandleFunc("/users/", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{[0-9]+}", server.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/signup/", server.SignUpAPI).Methods(http.MethodPost)
	router.HandleFunc("/users/{[0-9]+}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{[0-9]+}", server.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/login/", server.LoginAPI).Methods(http.MethodPost)

	server.Router = router
	server.Users = UserStorage{data, LastId}

	return server
}
