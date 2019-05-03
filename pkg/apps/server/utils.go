package server

import (
	"context"
	"net/http"

	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
)

func (server *Server) IsLogined(r *http.Request, secret []byte, cookieField string) bool {
	_, err := server.GetUserIdFromCookie(r)
	return err == nil
}

func (server *Server) GetUserIdFromCookie(r *http.Request) (int, error) {
	cookie, err := r.Cookie(server.CookieField)
	if err != nil {
		return 0, err
	}
	ctx := context.Background()
	StrUserId, err := server.AuthClient.GetIdFromCookie(ctx,
		&authorization.AuthCookie{
			Data:   cookie.Value,
			Secret: server.ServerConfig.Secret,
		})
	if err != nil {
		server.Logger.Log("GetUserIdFromCookie ", err)
		return int(0), err
	}
	return int(StrUserId.UserId), nil
}
