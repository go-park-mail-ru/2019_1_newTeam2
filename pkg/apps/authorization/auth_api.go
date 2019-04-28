package authorization

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
)

func (server *AuthServer) Logout(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("Logout")
	server.CreateCookie("logout", -1, w, r)
	w.WriteHeader(http.StatusOK)
	server.Logger.Log("successful logout")
}

func (server *AuthServer) IsLogin(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("IsLogin")
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

func (server *AuthServer) LoginAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("LoginAPI")
	var user models.UserAuth
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.Logger.Log(err)
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		server.Logger.Log(err)
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	if token, _, err := server.DB.Login(user.Username, user.Password, []byte(server.ServerConfig.Secret)); err != nil {
		textError := models.Error{err.Error()}
		responses.WriteToResponse(w, http.StatusUnauthorized, textError)
		return
	} else {
		server.CreateCookie(token, 60, w, r)
		w.Write([]byte(token))
		w.WriteHeader(http.StatusOK)
	}
}

func (server *AuthServer) SignUpAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("SignUpAPI")
	jsonStr := server.CreateUser(w, r)
	var user models.User
	err := json.Unmarshal(jsonStr, &user)
	if err != nil {
		server.Logger.Log(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token, _, err := server.DB.Login(user.Username, user.Password, []byte(server.ServerConfig.Secret)); err != nil {
		server.Logger.Log(err.Error())
	} else {
		server.CreateCookie(token, 60, w, r)
		w.Write([]byte(token))
	}
	w.WriteHeader(http.StatusOK)
}
