package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/user/2019_1_newTeam2/filesystem"
	"github.com/user/2019_1_newTeam2/models"
)

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	cookie := &http.Cookie{
		Name:  server.CookieField,
		Value: "logout",
	}
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(-1 * time.Microsecond)
	cookie.HttpOnly = true
	cookie.Secure = false
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	server.Logger.Log("successful logout")
}

func (server *Server) IsLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if value, _ := server.CheckLogin(w, r); !value {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{}"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (server *Server) LoginAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("LoginAPI")
	if r.Method == http.MethodOptions {
		textError := models.Error{""}
		WriteToResponse(w, http.StatusOK, textError)
		return
	}
	var user models.UserAuth
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{""}
		WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		textError := models.Error{""}
		WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	if token, _, err := server.DB.Login(user.Username, user.Password, []byte(server.ServerConfig.Secret)); err != nil {
		textError := models.Error{err.Error()}
		WriteToResponse(w, http.StatusUnauthorized, textError)
		return
	} else {
		cookie := &http.Cookie{
			Name:  server.CookieField,
			Value: token,
		}
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(5 * time.Hour)
		cookie.HttpOnly = true
		cookie.Secure = false
		http.SetCookie(w, cookie)
		w.Write([]byte(token))
		w.WriteHeader(http.StatusOK)
	}
}

func (server *Server) SignUpAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("SignUpAPI")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if value, _ := server.CheckLogin(w, r); value {
		w.WriteHeader(http.StatusOK)
	}
	jsonStr := server.CreateUser(w, r)
	var user models.User
	err := json.Unmarshal(jsonStr, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token, _, err := server.DB.Login(user.Username, user.Password, []byte(server.ServerConfig.Secret)); err != nil {
		server.Logger.Log(err.Error())
	} else {
		cookie := &http.Cookie{
			Name:  server.CookieField,
			Value: token,
		}
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(5 * time.Hour)
		cookie.HttpOnly = true
		cookie.Secure = false
		http.SetCookie(w, cookie)
		w.Write([]byte(token))
	}
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	value, user_id := server.CheckLogin(w, r)
	if !value {
		w.WriteHeader(http.StatusUnauthorized)
	}
	result, find, err := server.DB.GetUserByID(user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !find {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	function := func(header multipart.FileHeader) error {
		re := regexp.MustCompile(`image/.*`)
		if !re.MatchString(header.Header.Get("Content-Type")) {
			server.Logger.Log(header.Header.Get("Content-Type"))
			return fmt.Errorf("not an image")
		}
		return nil
	}
	_, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pathToAvatar, err := filesystem.UploadFile(w, r, function,
		server.ServerConfig.UploadPath, server.ServerConfig.AvatarsPath)
	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = server.DB.AddImage(pathToAvatar, userID)
	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (server *Server) UsersPaginate(w http.ResponseWriter, r *http.Request) {
	pages, ok := r.URL.Query()["page"]
	if !ok || len(pages[0]) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rows, ok := r.URL.Query()["rows"]
	if !ok || len(rows[0]) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(pages[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	rowsNum, err := strconv.Atoi(rows[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	result, err := server.DB.GetUsers(page, rowsNum)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if value, user_id := server.CheckLogin(w, r); value {
		var user models.User
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
		_, find, err := server.DB.GetUserByID(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !find {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		server.DB.UpdateUserById(user_id, user.Username, user.Email, user.Password, user.LangID, user.PronounceON)
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if value, user_id := server.CheckLogin(w, r); value {
		_, find, err := server.DB.GetUserByID(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !find {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		isDelete, _ := server.DB.DeleteUserById(user_id)
		if !isDelete {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
