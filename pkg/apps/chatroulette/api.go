package chatroulette

import (
	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"log"
	"net/http"
	"strconv"
)

func (server *ChatServer) CreateChat(w http.ResponseWriter, r *http.Request) {
	/*id, err := common.GetIdFromCookie(r, []byte(server.ServerConfig.Secret), server.CookieField)
	if err != nil {
		responses.WriteToResponse(w, http.StatusInternalServerError, models.Error{Message: "cannot subscribe"})
	}*/
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = server.Hub.AddClient(w, r, id)
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, models.Error{Message: "cannot subscribe"})
	}
	//responses.WriteToResponse(w, http.StatusOK, nil)
}

func (server *ChatServer) GetHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("get history")
}