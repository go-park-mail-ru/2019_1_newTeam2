package game

import (
	"github.com/user/2019_1_newTeam2/shared/pkg/responses"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/user/2019_1_newTeam2/pkg/apps/game/game"
	"github.com/user/2019_1_newTeam2/shared/models"
)

func (server *GameServer) OpenConnection(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Username, err := server.GetUsernameFromCookie(r)
	if err != nil {
		server.Logger.Log("cannot upgrade connection: %s", err)
	}
	// ws, err := upgrader.Upgrade(w.(http.ResponseWriter), r, nil)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		server.Logger.Log("cannot upgrade connection: %s", err)
	}

	_ = conn.WriteJSON(models.PlayerData{Username: Username, Score: 0})
	server.Game.Register <- &game.GameRegister{Conn: conn, Username: Username}
}

func (server *GameServer) GetDemo(w http.ResponseWriter, r *http.Request) {
	const wordsNum = 20
	cards, err := server.DB.GetWordsForDemo(wordsNum)
	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, cards)
}
