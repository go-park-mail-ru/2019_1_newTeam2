package chat

import (
	"github.com/user/2019_1_newTeam2/shared/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"github.com/user/2019_1_newTeam2/pkg/utils"
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
	var page, rowsNum int
	err := utils.ParseParams(w, r, &page, &rowsNum)
	if err != nil {
		return
	}
	result, err := server.DB.GetMessagesBroadcast(page, rowsNum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, result)
}
