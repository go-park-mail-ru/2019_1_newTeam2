package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func IsLogined(r *http.Request, secret []byte, cookieField string) bool {
	_, err := GetIdFromCookie(r, secret, cookieField)
	return err == nil
}

func GetIdFromCookie(r *http.Request, secret []byte, cookieField string) (int, error) {
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
