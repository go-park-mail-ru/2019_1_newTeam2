package chat

import (
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"net/http"
)

func (server *ChatServer) CreateChat(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("CreateChatAPI")
	id, err := server.GetUserIdFromCookie(r)
	if err != nil {
		responses.WriteToResponse(w, http.StatusInternalServerError, models.Error{Message: "cannot subscribe"})
		return
	}
	err = server.Hub.AddClient(w, r, id)
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, models.Error{Message: "cannot subscribe"})
	}
	//responses.WriteToResponse(w, http.StatusOK, nil)
}

func (server *ChatServer) GetHistory(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("GetHistoryAPI")

	responses.WriteToResponse(w, http.StatusOK, nil)
}