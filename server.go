package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/user/2019_1_newTeam2/storage"
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
	Users  storage.UserStorage
	Router *mux.Router
}

func InitServer() *Server {
	data := make(map[int]storage.User)
	LastId := 10
	for i := 0; i < LastId; i++ {

		data[i] = storage.User{i, "test_user_" + strconv.Itoa(i), "kek@lol.kl", "pass", 0, 1, 0}
	}
	router := mux.NewRouter()

	server := new(Server)

	router.HandleFunc("/users/", server.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users/{[0-9]+}", server.GetUser).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/signup/", server.SignUpAPI).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/users/{[0-9]+}", server.UpdateUser).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/users/{[0-9]+}", server.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/login/", server.LoginAPI).Methods(http.MethodPost, http.MethodOptions)

	server.Router = router
	server.Users = storage.UserStorage{data, LastId}

	return server
}
