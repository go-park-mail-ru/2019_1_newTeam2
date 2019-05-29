package server

import (
	"github.com/user/2019_1_newTeam2/shared/models"
	"github.com/user/2019_1_newTeam2/shared/pkg/responses"
	"net/http"
)

func (server *Server) GetLangs(w http.ResponseWriter, r *http.Request) {
	langs, found, err := server.DB.GetLangs()
	if err != nil {
		server.Logger.Log(err.Error())
		textError := models.Error{Message: err.Error()}
		responses.WriteToResponse(w, http.StatusInternalServerError, textError)
		return
	}

	if !found {
		responses.WriteToResponse(w, http.StatusNotFound, nil)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, langs)
}
