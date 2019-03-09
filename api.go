package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	. "./requests"
	. "./responses"
	. "./storage"
)

func (server *Server) LoginAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginAPI")
	if r.Method == http.MethodPost {
		var user UserAuth
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
		if token, err := server.Users.Login(user.Username, user.Password); err != nil {
			fmt.Println(err)
		} else {
			cookie := &http.Cookie{
				Name:  "session_id",
				Value: token,
			}
			http.SetCookie(w, cookie)
			w.Write([]byte(token))
		}
	}
}

func (server *Server) SignUpAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignUpAPI")
	if r.Method == http.MethodPost {
		server.CreateUser(w, r)
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
		if token, err := server.Users.Login(user.Username, user.Password); err != nil {
			fmt.Println(err)
		} else {
			cookie := &http.Cookie{
				Name:  "session_id",
				Value: token,
			}
			http.SetCookie(w, cookie)
			w.Write([]byte(token))
		}
	}
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(r.URL.Path[1:])
	fmt.Println(head, r.URL.Path, userID)
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
	UserResponse(w, http.StatusOK, result)
}

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
	server.Users.LastId++
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := server.Users.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	UserTableResponse(w, http.StatusOK, result)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	_, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	server.Users.UpdateUserById(userID, user.Username, user.Email, user.Password, user.LangID, user.PronounceON)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(r.URL.Path[1:])

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
