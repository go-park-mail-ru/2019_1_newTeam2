package chatroulette

import (
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/apps/common"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"log"
	"net/http"
)

func (server *ChatServer) CreateChat(w http.ResponseWriter, r *http.Request) {
	id, _ := common.GetIdFromCookie(r, []byte(server.ServerConfig.Secret), server.CookieField)
	id = 1
	err := server.Hub.AddClient(w, r, id)
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, models.Error{Message: "cannot subscribe"})
	}
	//responses.WriteToResponse(w, http.StatusOK, nil)
}

func (server *ChatServer) GetHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("get history")
}