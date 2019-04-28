package authorization

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
)

func (server *AuthServer) CreateUser(w http.ResponseWriter, r *http.Request) []byte {
	var user models.User
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.Logger.Log("CreateUser ", err)
		responses.WriteToResponse(w, http.StatusBadRequest, "")
		return jsonStr
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		server.Logger.Log("CreateUser ", err)
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return jsonStr
	}
	if br, err_r := server.DB.UserRegistration(user.Username, user.Email, user.Password, user.LangID, user.PronounceON); br != true {
		server.Logger.Log("CreateUser ", err_r.Error())
		textError := models.Error{err_r.Error()}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return jsonStr
	}
	return jsonStr
}

func (server *AuthServer) IsLogined(r *http.Request) bool {
	cookie, err := r.Cookie(server.CookieField)
	if err != nil {
		return false
	}
	ctx := context.Background()
	_, err = server.AuthClient.GetIdFromCookie(ctx, &AuthCookie{
		Data: cookie.Value,
		Secret: server.ServerConfig.Secret,
	})
	return err == nil
}

func (server *AuthServer) CreateCookie(token string, minutes int, w http.ResponseWriter, r *http.Request) {
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