package server

import (
	"encoding/json"
	"fmt"
	"github.com/user/2019_1_newTeam2/shared/models"
	"github.com/user/2019_1_newTeam2/shared/pkg/responses"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) GetSingleGame(w http.ResponseWriter, r *http.Request) {
	dictId, ok := r.URL.Query()["dict"]
	if !ok {
		responses.WriteToResponse(w, http.StatusBadRequest, fmt.Errorf("no dict id"))
	}
	dict, err := strconv.Atoi(dictId[0])
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, fmt.Errorf("dict id incorrect"))
	}

	wordsNum, ok := r.URL.Query()["words"]
	if !ok {
		responses.WriteToResponse(w, http.StatusBadRequest, fmt.Errorf("no words num"))
	}
	num, err := strconv.Atoi(wordsNum[0])
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, fmt.Errorf("words num incorrect"))
	}
	cards, found, err := server.DB.GetCardsForGame(dict, num)
	if err != nil {
		responses.WriteToResponse(w, http.StatusInternalServerError, fmt.Errorf("db error"))
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responses.WriteToResponse(w, http.StatusOK, cards)
}

func (server *Server) SetGameResults(w http.ResponseWriter, r *http.Request) {
	results := models.GameResults{}

	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	err = json.Unmarshal(jsonStr, &results)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	err, found := server.DB.UpdateFrequencies(results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
