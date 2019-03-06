package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	. "./responses"
	. "./storage"
)

func (server *Server) UsersAPI(w http.ResponseWriter, r *http.Request) {
	head, _ := TypeRequest(r.URL.Path)
	fmt.Println("UsersAPI: ", head)
	if head == "" {
		if r.Method == http.MethodGet {
			server.GetUsers(w, r)
		}
	}
	switch r.Method {
	case http.MethodGet:
		fmt.Println("get")
		server.GetUser(w, r)
	case http.MethodPost:
		fmt.Println("post")
		server.CreateUser(w, r)
	case http.MethodPut:
		fmt.Println("put")
		server.UpdateUser(w, r)
	case http.MethodDelete:
		fmt.Println("delete")
		server.DeleteUser(w, r)
	default:
		fmt.Println("default")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(head)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, find, err := server.Users.GetUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !find {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Response(w, http.StatusOK, result)
}

// curl -d '{"id":1, "username":"tsaanstu", "email":"m@mail.ru", "password":"password", "langID":1, "pronounceOn":1}' -H "Content-Type: application/json" -X POST http://localhost:8090/users/
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	server.Users.UserRegistration(user.Username, user.Email, user.Password, user.LangID, user.PronounceON)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, find, err := server.Users.GetUserByID(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !find {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	server.Users.UpdateUserById(user.ID, user.Username, user.Email, user.Password, user.LangID, user.PronounceON)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(head)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, find, err := server.Users.GetUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !find {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	isDelete, _ := server.Users.DeleteUserById(userID)
	if !isDelete {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
