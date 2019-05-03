package game

import (
	"context"
	"net/http"
	//"strconv"

	"github.com/gorilla/websocket"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
	"github.com/user/2019_1_newTeam2/pkg/apps/game/game"
)

func (server *GameServer) OpenConnection(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Username, err := server.GetUsernameFromCookie(r)
	conn, err := upgrader.Upgrade(w, r, http.Header{"Upgrade": []string{"websocket"}})
	if err != nil {
		server.Logger.Log("cannot upgrade connection: %s", err)
	}

	conn.WriteJSON(models.PlayerData{Username, 0})
	server.Game.Register <- &game.GameRegister{conn, Username}
}

func (server *GameServer) GetUserIdFromCookie(r *http.Request) (int, error){
	cookie, err := r.Cookie(server.CookieField)
	if err != nil {
		return 0, err
	}
	ctx := context.Background()
	StrUserId, err := server.AuthClient.GetIdFromCookie(ctx,
		&authorization.AuthCookie {
		Data: cookie.Value,
		Secret: server.ServerConfig.Secret,
	})
	if err != nil {
		server.Logger.Log("GetUserIdFromCookie ", err)
		return int(0), err
	}
	return int(StrUserId.UserId), nil
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