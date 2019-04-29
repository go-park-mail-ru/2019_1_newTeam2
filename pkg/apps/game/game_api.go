package game

import (
	"net/http"
	"context"
	"strconv"

	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/apps/authorization"
	"github.com/gorilla/websocket"
)

func (server *GameServer) OpenConnection(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	UserId, err := server.GetUserIdFromCookie(r)
	conn, err := upgrader.Upgrade(w, r, http.Header{"Upgrade": []string{"websocket"}})
	if err != nil {
		server.Logger.Log("cannot upgrade connection: %s", err)
	}

	conn.WriteJSON(models.GameMessage{"CONNECTED" + strconv.Itoa(UserId), nil})
	// g.Register <- conn
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