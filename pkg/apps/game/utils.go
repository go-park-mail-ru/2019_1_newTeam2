package game

import (
	"context"
	"net/http"

	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
)

func (server *GameServer) IsLogined(r *http.Request, secret []byte, cookieField string) bool {
	_, err := server.GetUsernameFromCookie(r)
	return err == nil
}

func (server *GameServer) GetUsernameFromCookie(r *http.Request) (string, error){
	cookie, err := r.Cookie(server.CookieField)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	StrUsername, err := server.AuthClient.GetUsernameFromCookie(ctx,
		&authorization.AuthCookie {
			Data: cookie.Value,
			Secret: server.ServerConfig.Secret,
		})
	if err != nil {
		server.Logger.Log("GetUserIdFromCookie ", err)
		return "", err
	}
	return StrUsername.Data, nil
}