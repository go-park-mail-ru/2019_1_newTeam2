package chat

import (
	"context"
	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
	"net/http"
)

func (server *ChatServer) GetUserIdFromCookie(r *http.Request) (int, error) {
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

func (server *ChatServer) GetUserNameFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(server.CookieField)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	StrUsername, err := server.AuthClient.GetUsernameFromCookie(ctx,
		&authorization.AuthCookie{
			Data:   cookie.Value,
			Secret: server.ServerConfig.Secret,
		})
	if err != nil {
		server.Logger.Log("GetUserIdFromCookie ", err)
		return "", err
	}
	return string(StrUsername.Data), nil
}
