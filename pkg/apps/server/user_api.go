package server

import (
	"encoding/json"
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/user/2019_1_newTeam2/shared/filesystem"
	"github.com/user/2019_1_newTeam2/shared/models"
	"github.com/user/2019_1_newTeam2/shared/pkg/responses"
	"github.com/user/2019_1_newTeam2/shared/pkg/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"regexp"
)

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("GetUser")
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
	//pathToAvatar, err := filesystem.UploadFile(w, r, function,
	//	server.ServerConfig.UploadPath, server.ServerConfig.AvatarsPath)
	pathToAvatar, err := filesystem.UploadFileToCloud(w, r, function, server.svc)

	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//pathToAvatar = strings.TrimPrefix(pathToAvatar, "files/")
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
	err := utils.ParseParams(w, r, &page, &rowsNum)
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
	server.Logger.Log("UpdateUser")
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

	_, _ = server.DB.UpdateUserById(userId, user.Username, user.Email /*user.Password, */, user.LangID, user.PronounceON)
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

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	server.CreateCookie("logout", -1, w, r)
	w.WriteHeader(http.StatusOK)
	server.Logger.Log("successful logout")
}

func (server *Server) IsLogin(w http.ResponseWriter, r *http.Request) {
	if value := server.IsLogined(r, []byte(server.ServerConfig.Secret), server.CookieField); !value {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent)
		_, _ = w.Write([]byte("{}"))
		server.Logger.Log("User not logined")
		return
	}
	server.Logger.Log("User is logined")
	w.WriteHeader(http.StatusOK)
}

func (server *Server) LoginAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("LoginAPI")
	var user models.UserAuth
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	if token, _, err := server.DB.Login(user.Username, user.Password, []byte(server.ServerConfig.Secret)); err != nil {
		textError := models.Error{Message: err.Error()}
		responses.WriteToResponse(w, http.StatusUnauthorized, textError)
		return
	} else {
		server.CreateCookie(token, 60, w, r)
		_, _ = w.Write([]byte(token))
		w.WriteHeader(http.StatusOK)
	}
}

func (server *Server) SignUpAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("SignUpAPI")
	jsonStr := server.CreateUser(w, r)
	var user models.User
	err := easyjson.Unmarshal(jsonStr, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token, _, err := server.DB.Login(user.Username, user.Password, []byte(server.ServerConfig.Secret)); err != nil {
		server.Logger.Log(err.Error())
	} else {
		server.CreateCookie(token, 60, w, r)
		_, _ = w.Write([]byte(token))
	}
	w.WriteHeader(http.StatusOK)
}
