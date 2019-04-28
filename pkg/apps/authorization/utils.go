package authorization

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
)

func (server *AuthServer) CreateUser(w http.ResponseWriter, r *http.Request) []byte {
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

func (server *AuthServer) IsLogined(r *http.Request, secret []byte, cookieField string) bool {
	cookie, err := r.Cookie(cookieField)
	if err != nil {
		return false
	}
	ctx := context.Background()
	_, err = server.GetIdFromCookie(ctx, &AuthCookie{
		Data: cookie.Value,
	})
	return err == nil
}

//func GetIdFromCookie(r *http.Request, secret []byte, cookieField string) (int, error) {
func (server *AuthServer) GetIdFromCookie(ctx context.Context, in *AuthCookie) (*Id, error) {
	token, err := jwt.Parse(in.Data, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return server.ServerConfig.Secret, nil
	})

	if err != nil {
		return &Id{UserId: 0}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &Id{UserId: int64(claims["id"].(float64))}, nil
	}
	return &Id{UserId: 0}, fmt.Errorf("token invalid")
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