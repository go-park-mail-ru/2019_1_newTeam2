package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/user/2019_1_newTeam2/models"
	"github.com/user/2019_1_newTeam2/pkg/responses"
)

func (server *Server) CreateDictionaryAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("CreateDictionaryAPI")
	userId, _ := GetIdFromCookie(r, []byte(server.ServerConfig.Secret), server.CookieField)
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

	if err = server.DB.DictionaryCreate(userId, dictionary.Name, dictionary.Description, dictionary.Cards); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (server *Server) UpdateDictionaryAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("UpdateDictionaryAPI")
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
	server.Logger.Log("DeleteDictionaryAPI")
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
