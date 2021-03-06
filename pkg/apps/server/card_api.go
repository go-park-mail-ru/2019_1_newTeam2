package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/user/2019_1_newTeam2/shared/filesystem"
	"github.com/user/2019_1_newTeam2/shared/models"
	"github.com/user/2019_1_newTeam2/shared/pkg/responses"
	"github.com/user/2019_1_newTeam2/shared/pkg/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func (server *Server) UploadWordsFileAPI(w http.ResponseWriter, r *http.Request) {
	function := func(header multipart.FileHeader) error {
		re := regexp.MustCompile(`application/.*`)
		if !re.MatchString(header.Header.Get("Content-Type")) {
			server.Logger.Log(header.Header.Get("Content-Type"))
			return fmt.Errorf("not a file")
		}
		return nil
	}

	server.Logger.Log("UploadWordsFileAPI")
	dictionaryIdString, parseErr := r.URL.Query()["dictionaryId"]
	if !parseErr {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dictionaryId, ConvErr := strconv.Atoi(dictionaryIdString[0])
	if ConvErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_ = os.Mkdir(server.ServerConfig.UploadPath+"temp_docs/"+strconv.Itoa(userId), 0777)
	pathToFile, err := filesystem.UploadFile(w, r, function,
		server.ServerConfig.UploadPath, "temp_docs/"+strconv.Itoa(userId))

	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, pathToFile = utils.TypeRequest(pathToFile)
	pathToFile = server.ServerConfig.UploadPath[:len(server.ServerConfig.UploadPath)-1] + pathToFile
	err = server.DB.FillDictionaryFromXLSX(userId, dictionaryId, pathToFile)
	_ = os.RemoveAll(server.ServerConfig.UploadPath + "temp_docs/" + strconv.Itoa(userId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (server *Server) CardsPaginate(w http.ResponseWriter, r *http.Request) {
	page := 0
	rowsNum := 0
	dictId, ok := r.URL.Query()["dict"]
	if !ok {
		responses.WriteToResponse(w, http.StatusBadRequest, models.Error{Message: "no dict"})
	}
	dict, err := strconv.Atoi(dictId[0])
	if err != nil {
		responses.WriteToResponse(w, http.StatusBadRequest, models.Error{Message: "incorrect dict"})
	}
	err = utils.ParseParams(w, r, &page, &rowsNum)
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

func (server *Server) DeleteCardInDictionary(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("DeleteCardInDictionaryAPI")
	fmt.Println(r.URL.Query())
	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var card models.CardDelete

	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	err = easyjson.Unmarshal(jsonStr, &card)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	if err = server.DB.DeleteCardInDictionary(userId, card.DictionaryId, card.CardId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (server *Server) CreateCardInDictionary(w http.ResponseWriter, r *http.Request) {
	userId, err := server.GetUserIdFromCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	server.Logger.Log("CreateCardInDictionaryAPI")
	dictionaryIdString, parseErr := r.URL.Query()["dictionaryId"]
	if !parseErr {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	dictionaryId, ConvErr := strconv.Atoi(dictionaryIdString[0])
	if ConvErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var card models.Card
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = easyjson.Unmarshal(jsonStr, &card)
	if err != nil {
		textError := models.Error{Message: ""}
		responses.WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}

	if err = server.DB.SetCardToDictionary(userId, dictionaryId, card); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
