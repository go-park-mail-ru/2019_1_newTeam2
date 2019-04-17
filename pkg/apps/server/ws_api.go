package server

import (
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"net/http"
)

func (server *Server) WSSubscribe(w http.ResponseWriter, r *http.Request) {
	id, _ := GetIdFromCookie(r, []byte(server.ServerConfig.Secret), server.CookieField)
	err := server.Hub.AddClient(w, r, id)
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, models.Error{Message: "cannot subscribe"})
	}
	responses.WriteToResponse(w, http.StatusOK, nil)
}

