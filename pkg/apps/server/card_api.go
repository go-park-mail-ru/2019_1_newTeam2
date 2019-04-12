package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/pkg/responses"
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
	fmt.Println("id: ", idStr)
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

func (server *Server) CreateCardInDictionary(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("CreateCardInDictionaryAPI")
	// dictionaryIdString, parseErr := r.URL.Query()["dictionaryId"]
	// if !parseErr {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	// dictionaryId, ConvErr := strconv.Atoi(dictionaryIdString[0])
	// if ConvErr != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// }
	// var card models.Card
	// jsonStr, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	textError := models.Error{""}
	// 	responses.WriteToResponse(w, http.StatusBadRequest, textError)
	// 	return
	// }
	// err = json.Unmarshal(jsonStr, &card)
	// if err != nil {
	// 	textError := models.Error{""}
	// 	responses.WriteToResponse(w, http.StatusBadRequest, textError)
	// 	return
	// }

	// if err = server.DB.SetCardToDictionary(dictionaryId, card); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
}
