package game

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/user/2019_1_newTeam2/shared/models"
	"github.com/user/2019_1_newTeam2/pkg/apps/game/game"
)

func (server *GameServer) OpenConnection(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Username, err := server.GetUsernameFromCookie(r)
	// ws, err := upgrader.Upgrade(w.(http.ResponseWriter), r, nil)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		server.Logger.Log("cannot upgrade connection: %s", err)
	}

	_ = conn.WriteJSON(models.PlayerData{Username, 0})
	server.Game.Register <- &game.GameRegister{conn, Username}
}
