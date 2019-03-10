package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"

	"github.com/user/2019_1_newTeam2/newfs1"
	"github.com/user/2019_1_newTeam2/requests"
	"github.com/user/2019_1_newTeam2/responses"
	"github.com/user/2019_1_newTeam2/storage"
)

func (server *Server) LoginAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginAPI")
	if r.Method == http.MethodPost {
		var user requests.UserAuth
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

		// fmt.Println(w.Cookie["session_id"])
		// fmt.Println(server.GetMyCoookieMan(w, r))

		// server.Users.IsLogin(w, r, user.Username, user.Password)

		test, _ := r.Cookie("session_id")
		fmt.Println("cookie:", test) //  TODO (tsaanstu): kekekekekekekekekekekekekek

		if token, err := server.Users.Login(user.Username, user.Password); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			cookie := &http.Cookie{
				Name:  "session_id",
				Value: token,
			}
			http.SetCookie(w, cookie)
			// w.Write([]byte(token))
			w.WriteHeader(http.StatusOK)
		}
		// fmt.Println("qwerty")
	}
}

func (server *Server) SignUpAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignUpAPI")
	if r.Method == http.MethodPost {
		jsonStr := server.CreateUser(w, r)
		var user storage.User
		fmt.Println("json: ", jsonStr)
		err := json.Unmarshal(jsonStr, &user)
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
	responses.UserResponse(w, http.StatusOK, result)
}

func (server *Server) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	function := func(header multipart.FileHeader) error {
		re := regexp.MustCompile(`image/.*`)
		if !re.MatchString(header.Header.Get("Content-Type")) {
			fmt.Println(header.Header.Get("Content-Type"))
			return fmt.Errorf("not an image")
		}
		return nil
	}
	_, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pathToAvatar, err := newfs1.UploadFile(w, r, function, "avatars/")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = server.Users.AddImage(pathToAvatar, userID)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) []byte {
	var user storage.User
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return jsonStr
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return jsonStr
	}
	server.Users.UserRegistration(user.Username, user.Email, user.Password, user.LangID, user.PronounceON)
	server.Users.LastId++
	return jsonStr
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := server.Users.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	responses.UserTableResponse(w, http.StatusOK, result)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	_, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user storage.User
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
