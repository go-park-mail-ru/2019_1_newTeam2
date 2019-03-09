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

	router.HandleFunc("/users/", server.GetUsers).Methods("GET")
	router.HandleFunc("/users/{[0-9]+}", server.GetUser).Methods("GET") //  id:[0-9]+
	router.HandleFunc("/signup/", server.SignUpAPI).Methods("POST")
	router.HandleFunc("/users/{[0-9]+}", server.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{[0-9]+}", server.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/signup/", server.SignupAPI).Methods("POST")
	router.HandleFunc("/login/", server.LoginAPI).Methods("POST")

	server.Router = router
	server.Users = UserStorage{data, LastId}
	// server.LastId = LastId

	return server
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("ServeHTTP")
	// var head, url string
	// head, url = TypeRequest(r.URL.Path)

	// router := mux.NewRouter()

	// switch head {
	// case "users":
	// 	r.URL.Path = url
	// 	fmt.Println("users---> ", head)
	// 	server.UsersAPI(w, r)
	// case "login":
	// 	fallthrough
	// case "signup":
	// 	fmt.Println("auth module---> ", head)
	// 	server.AuthAPI(w, r)
	// default:
	// 	fmt.Println("default")
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// }
}
