package server

import (
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"net/http"
)

func (server *Server) WSSubscribe(w http.ResponseWriter, r *http.Request) {
	id, _ := GetIdFromCookie(r, []byte(server.ServerConfig.Secret), server.CookieField)
	//id := 0
	err := server.hub.AddClient(w, r, id)
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, models.Error{Message: "cannot subscribe"})
	}
}

func (server *Server) WSUnsubscribe(w http.ResponseWriter, r *http.Request) {
	id, _ := GetIdFromCookie(r, []byte(server.ServerConfig.Secret), server.CookieField)
	server.hub.DeleteClient(id)
	responses.WriteToResponse(w, http.StatusOK, nil)
}
