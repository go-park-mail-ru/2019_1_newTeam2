package authorization

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type AuthManager struct {}

func NewAuthManager() *AuthManager {
	return &AuthManager{}
}

func (am *AuthManager) GetIdFromCookie(ctx context.Context, in *AuthCookie) (*Id, error) {
	token, err := jwt.Parse(in.Data, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(in.Secret), nil
	})
	if err != nil {
		return &Id{UserId: 0}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &Id{UserId: int64(claims["id"].(float64))}, nil
	}
	return &Id{UserId: 0}, fmt.Errorf("token invalid")
}