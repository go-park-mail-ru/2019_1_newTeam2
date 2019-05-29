package server

import (
	"github.com/mailru/easyjson"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/user/2019_1_newTeam2/shared/models"
	"github.com/user/2019_1_newTeam2/shared/pkg/responses"
	"github.com/user/2019_1_newTeam2/shared/pkg/utils"
)

func (server *Server) CreateDictionaryAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("CreateDictionaryAPI")
	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var dictionary models.CreateDictionary
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = easyjson.Unmarshal(jsonStr, &dictionary)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	result, err := server.DB.DictionaryCreate(userId, dictionary.Name, dictionary.Description, dictionary.Cards)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responses.WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) UpdateDictionaryAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("UpdateDictionaryAPI")
	var dictionary models.CreateDictionary
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = easyjson.Unmarshal(jsonStr, &dictionary)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	if err = server.DB.DictionaryUpdate(dictionary.ID, dictionary.Name, dictionary.Description); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, _, err := server.DB.GetDict(dictionary.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responses.WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) DeleteDictionaryAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("DeleteDictionaryAPI")
	// var delete models.ParametersId
	/*jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = json.Unmarshal(jsonStr, &delete)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}*/
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = server.DB.DictionaryDelete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (server *Server) GetDictionaryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if result.ID == userId {
		result.Privilege = true
	} else {
		result.Privilege = false
	}
	responses.WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) DictsPaginate(w http.ResponseWriter, r *http.Request) {
	page := 0
	rowsNum := 0
	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = utils.ParseParams(w, r, &page, &rowsNum)
	if err != nil {
		return
	}
	result, found, err := server.DB.GetDicts(userId, page, rowsNum)
	if err != nil {
		server.Logger.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !found {
		server.Logger.Log(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	responses.WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) BorrowDictById(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	idStr := vars["id"]
	dictId, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, _ := common.GetIdFromCookie(r, []byte(server.ServerConfig.Secret), server.CookieField)
	//ownerId, createdDict, err := server.DB.BorrowDictById(dictId, userId)
	if err == storage.ErrNotFound {
		responses.WriteToResponse(w, http.StatusNotFound, nil)
	}

	if err != nil {
		responses.WriteToResponse(w, http.StatusInternalServerError, nil)
	}
	responses.WriteToResponse(w, http.StatusOK, createdDict)
	//mes := wshub.Message{ID: ownerId, Data: models.DictionaryNote{BorrowerId: userId, DictionaryName: createdDict.Name}}
	//server.Hub.SendToClient(&mes)*/
}
