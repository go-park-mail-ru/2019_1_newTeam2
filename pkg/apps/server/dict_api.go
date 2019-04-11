package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
	"io/ioutil"
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

func (server *Server) CreateDictionaryAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("LoginAPI")
	var dictionary models.CreateDictionary
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = json.Unmarshal(jsonStr, &dictionary)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	if err = server.DB.DictionaryCreate(2, dictionary.Name, dictionary.Description, dictionary.Cards); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (server *Server) UpdateDictionaryAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("LoginAPI")
	var dictionary models.CreateDictionary
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = json.Unmarshal(jsonStr, &dictionary)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	if err = server.DB.DictionaryUpdate(dictionary.ID, dictionary.Name, dictionary.Description); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (server *Server) DeleteDictionaryAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("LoginAPI")
	var dictionary models.CreateDictionary
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = json.Unmarshal(jsonStr, &dictionary)
	if err != nil {
		textError := models.Error{""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	// if err = server.DB.DictionaryUpdate(dictionary.ID, dictionary.Name, dictionary.Description); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
}
