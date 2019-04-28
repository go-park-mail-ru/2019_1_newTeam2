package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"regexp"

	"github.com/user/2019_1_newTeam2/filesystem"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
)

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, find, err := server.DB.GetUserByID(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !find {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	function := func(header multipart.FileHeader) error {
		re := regexp.MustCompile(`image/.*`)
		if !re.MatchString(header.Header.Get("Content-Type")) {
			server.Logger.Log(header.Header.Get("Content-Type"))
			return fmt.Errorf("not an image")
		}
		return nil
	}

	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
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
	err = server.DB.AddImage(pathToAvatar, userId)
	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (server *Server) UsersPaginate(w http.ResponseWriter, r *http.Request) {
	page := 0
	rowsNum := 0
	err := ParseParams(w, r, &page, &rowsNum)
	if err != nil {
		return
	}
	result, found, err := server.DB.GetUsers(page, rowsNum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !found {
		server.Logger.Log("No such a user")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	_, find, err := server.DB.GetUserByID(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !find {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	server.DB.UpdateUserById(userId, user.Username, user.Email /*user.Password, */, user.LangID, user.PronounceON)
	w.WriteHeader(http.StatusOK)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	isDelete, err := server.DB.DeleteUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isDelete {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
