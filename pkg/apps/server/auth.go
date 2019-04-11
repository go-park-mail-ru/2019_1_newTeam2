package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
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
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return jsonStr
	}
	if br, err_r := server.DB.UserRegistration(user.Username, user.Email, user.Password, user.LangID, user.PronounceON); br != true {
		server.Logger.Log(err_r.Error())
		textError := models.Error{err_r.Error()}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return jsonStr
	}
	return jsonStr
}


func IsLogined(r *http.Request, secret []byte, cookieField string) bool {
	_, err := GetIdFromCookie(r, secret, cookieField)
	return err == nil
}

func GetIdFromCookie(r *http.Request, secret []byte, cookieField string) (int, error){
	cookie, err := r.Cookie(cookieField)

	if err != nil {
		return 0, err
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["id"].(float64)), nil
	}
	return 0, fmt.Errorf("token invalid")
}
