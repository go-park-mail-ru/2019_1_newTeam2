package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"

	"github.com/user/2019_1_newTeam2/filesystem"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
)

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	server.CreateCookie("logout", -1, w, r)
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
		server.Logger.Log("User not logined")
		return
	}
	server.Logger.Log("User is logined")
	w.WriteHeader(http.StatusOK)
}

func (server *Server) LoginAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("LoginAPI")
	if r.Method == http.MethodOptions {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusOK, textError)
		return
	}
	var user models.UserAuth
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	if token, _, err := server.DB.Login(user.Username, user.Password, []byte(server.ServerConfig.Secret)); err != nil {
		textError := models.Error{err.Error()}
		responses.WriteToResponse(w, http.StatusUnauthorized, textError)
		return
	} else {
		server.CreateCookie(token, 20, w, r)
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
		server.CreateCookie(token, 20, w, r)
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
		return
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
	responses.WriteToResponse(w, http.StatusOK, result)
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

	errBool, userID := server.CheckLogin(w, r)
	if !errBool {
		server.Logger.Log("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
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
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	pages, ok := r.URL.Query()["page"]
	if !ok || len(pages[0]) < 1 {
		server.Logger.Log("No pages in query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rows, ok := r.URL.Query()["rows"]
	if !ok || len(rows[0]) < 1 {
		server.Logger.Log("No rows in query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(pages[0])
	if err != nil {
		server.Logger.Log("Incorrect pages")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rowsNum, err := strconv.Atoi(rows[0])
	if err != nil {
		server.Logger.Log("Incorrect rows")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := server.DB.GetUsers(page, rowsNum)
	if err != nil {
		server.Logger.Log("No such a user")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	value, user_id := server.CheckLogin(w, r)
	if !value {
		w.WriteHeader(http.StatusUnauthorized)
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

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
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
