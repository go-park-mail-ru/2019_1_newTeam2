package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/user/2019_1_newTeam2/filesystem"
	"github.com/user/2019_1_newTeam2/models"
)

func (server *Server) CheckLogin(w http.ResponseWriter, r *http.Request) (bool, int) {
	SECRET := []byte(server.ServerConfig.Secret)
	myCookie, err := r.Cookie("session_id")

	if err != nil {
		return false, -1
	}

	token, err := jwt.Parse(myCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SECRET, nil
	})

	if err != nil {
		server.Logger.Log(err.Error())
		return false, -1
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, int(claims["id"].(float64))
	}
	server.Logger.Log(err)
	return false, -1
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: "logout",
	}
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(1 * time.Microsecond)
	cookie.HttpOnly = false
	cookie.Secure = false
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	server.Logger.Log("successful logout")
}

func (server *Server) IsLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if value, _ := server.CheckLogin(w, r); !value {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{}"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (server *Server) LoginAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("LoginAPI")
	if r.Method == http.MethodOptions {
		textError := models.Error{""}
		WriteToResponse(w, http.StatusOK, textError)
		return
	}
	var user models.UserAuth
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		textError := models.Error{""}
		WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		textError := models.Error{""}
		WriteToResponse(w, http.StatusBadRequest, textError)
		return
	}
	if token, _, err := server.DB.Login(user.Username, user.Password, []byte(server.ServerConfig.Secret)); err != nil {
		textError := models.Error{err.Error()}
		WriteToResponse(w, http.StatusUnauthorized, textError)
		return
	} else {
		cookie := &http.Cookie{
			Name:  "session_id",
			Value: token,
		}
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(5 * time.Hour)
		cookie.HttpOnly = false
		cookie.Secure = false
		http.SetCookie(w, cookie)
		w.Write([]byte(token))
		w.WriteHeader(http.StatusOK)
	}
}

func (server *Server) SignUpAPI(w http.ResponseWriter, r *http.Request) {
	server.Logger.Log("SignUpAPI")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		if value, _ := server.CheckLogin(w, r); value {
			w.WriteHeader(http.StatusOK)
		}
		jsonStr := server.CreateUser(w, r)
		var user models.User
		err := json.Unmarshal(jsonStr, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if token, _, err := server.DB.Login(user.Username, user.Password, []byte(server.ServerConfig.Secret)); err != nil {
			server.Logger.Log(err.Error())
		} else {
			cookie := &http.Cookie{
				Name:  "session_id",
				Value: token,
			}
			cookie.Path = "/"
			cookie.Expires = time.Now().Add(5 * time.Hour)
			cookie.HttpOnly = false
			cookie.Secure = false
			http.SetCookie(w, cookie)
			w.Write([]byte(token))
		}
	}
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if value, user_id := server.CheckLogin(w, r); value {
		result, find, err := server.DB.GetUserByID(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !find {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		WriteToResponse(w, http.StatusOK, result)
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (server *Server) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	function := func(header multipart.FileHeader) error {
		re := regexp.MustCompile(`image/.*`)
		if !re.MatchString(header.Header.Get("Content-Type")) {
			server.Logger.Log(header.Header.Get("Content-Type"))
			return fmt.Errorf("not an image")
		}
		return nil
	}
	_, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pathToAvatar, err := filesystem.UploadFile(w, r, function,
		server.ServerConfig.UploadPath, server.ServerConfig.AvatarsPath)
	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = server.DB.AddImage(pathToAvatar, userID)
	if err != nil {
		server.Logger.Log(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) []byte {
	var user models.User
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteToResponse(w, http.StatusBadRequest, "")
		return jsonStr
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		textError := models.Error{""}
		WriteToResponse(w, http.StatusBadRequest, textError)
		return jsonStr
	}
	if br, err_r := server.DB.UserRegistration(user.Username, user.Email, user.Password, user.LangID, user.PronounceON); br != true {
		server.Logger.Log(err_r.Error())
		textError := models.Error{err_r.Error()}
		WriteToResponse(w, http.StatusBadRequest, textError)
		return jsonStr
	}
	server.DB.IncUserLastID()
	return jsonStr
}

func (server *Server) UsersPaginate(w http.ResponseWriter, r *http.Request) {
	pages, ok := r.URL.Query()["page"]
	if !ok || len(pages[0]) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rows, ok := r.URL.Query()["rows"]
	if !ok || len(rows[0]) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(pages[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	rowsNum, err := strconv.Atoi(rows[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	result, err := server.DB.GetUsers(page, rowsNum)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	WriteToResponse(w, http.StatusOK, result)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if value, user_id := server.CheckLogin(w, r); value {
		var user models.User
		jsonStr, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(jsonStr, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, find, err := server.DB.GetUserByID(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !find {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		server.DB.UpdateUserById(user_id, user.Username, user.Email, user.Password, user.LangID, user.PronounceON)
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if value, user_id := server.CheckLogin(w, r); value {
		_, find, err := server.DB.GetUserByID(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !find {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		isDelete, _ := server.DB.DeleteUserById(user_id)
		if !isDelete {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}