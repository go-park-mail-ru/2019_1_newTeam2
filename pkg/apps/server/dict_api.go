package server

import (
	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"net/http"
	"strconv"
)

func (server *Server) GetDictionaryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// TODO(sergeychur): have a look some shit with ids
	result, found, err := server.DB.GetDict(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !found {
		server.Logger.Log("No such a card")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) DictsPaginate(w http.ResponseWriter, r *http.Request) {
	page := 0
	rowsNum := 0
	userId, _ := GetIdFromCookie(r, []byte(server.ServerConfig.Secret), server.CookieField)
	err := ParseParams(w, r, &page, &rowsNum)
	if err != nil {
		return
	}
	result, found, err := server.DB.GetDicts(userId, page, rowsNum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !found {
		server.Logger.Log("No suitable dicts")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, result)
}
