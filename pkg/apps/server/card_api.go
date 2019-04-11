package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"net/http"
	"strconv"
)

func (server *Server) CardsPaginate(w http.ResponseWriter, r *http.Request) {
	page := 0
	rowsNum := 0
	/*vars := mux.Vars(r)
	dictIdStr := vars["dictId"]
	dictId, err := strconv.Atoi(dictIdStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/
	dictId, ok := r.URL.Query()["dict"]
	if !ok {
		responses.WriteToResponse(w, http.StatusBadRequest, fmt.Errorf("no dict id"))
	}
	dict, err := strconv.Atoi(dictId[0])
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, fmt.Errorf("dict idincorrect"))
	}
	err = ParseParams(w, r, &page, &rowsNum)
	if err != nil {
		return
	}
	result, found, err := server.DB.GetCards(dict, page, rowsNum)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !found {
		server.Logger.Log("No suitable cards")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) GetCardById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, found, err := server.DB.GetCard(id)
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
