package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func (server *Server) CheckLogin(w http.ResponseWriter, r *http.Request) (bool, int) {
	SECRET := []byte(server.ServerConfig.Secret)
	myCookie, err := r.Cookie(server.CookieField)

	if err != nil {
		return false, -1
	}

	token, err := jwt.Parse(myCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SECRET, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return false, -1
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, int(claims["id"].(float64))
	}
	fmt.Println(err)
	return false, -1
}

func TypeRequest(url string) (string, string) {
	separatorPosition := strings.Index(url[1:], "/")
	if separatorPosition < 0 {
		return url[1:], "/"
	}
	separatorPosition++
	return url[1:separatorPosition], url[separatorPosition:]
}

func WriteToResponse(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	response, _ := json.Marshal(v)
	w.Write(response)
}
