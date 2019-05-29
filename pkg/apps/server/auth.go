package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/user/2019_1_newTeam2/shared/models"
	"github.com/user/2019_1_newTeam2/shared/pkg/responses"
)

func (server *Server) CreateCookie(token string, minutes int, w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  server.CookieField,
		Value: token,
	}
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Duration(minutes) * time.Minute)
	cookie.HttpOnly = true
	cookie.Secure = false
	http.SetCookie(w, cookie)
}

//  создание пользователя и возвращение данных в функцию регистрации для дальнейшей авторизации
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) []byte {
	var user models.User
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, "")
		return jsonStr
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return jsonStr
	}
	if br, err_r := server.DB.UserRegistration(user.Username, user.Email, user.Password, user.LangID, user.PronounceON); !br {
		server.Logger.Log(err_r.Error())
		textError := models.Error{Message: err_r.Error()}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return jsonStr
	}
	return jsonStr
}
