package server

import (
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"net/http"
)

func (server *Server) GetLangs(w http.ResponseWriter, r *http.Request) {
	langs, found, err := server.DB.GetLangs()
	if err != nil {
		server.Logger.Log(err.Error())
		textError := models.Error{err.Error()}
		responses.WriteToResponse(w, http.StatusInternalServerError, textError)
		return
	}

	if !found {
		responses.WriteToResponse(w, http.StatusNotFound, nil)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, langs)
}