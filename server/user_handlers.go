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
	fmt.Println("checklogin")
	// SECRET := []byte("kekusmaxima")
	SECRET := []byte(server.serverConfig.Secret)
	myCookie, err := r.Cookie("session_id")

	fmt.Println("в CheckLogin пришло cookie:", myCookie)

	if err != nil {
		return false, -1
	}
	token, err := jwt.Parse(myCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("NOT OK")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		fmt.Println("qwertyuiop")
		return SECRET, nil
	})
	fmt.Println("END")
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println([]byte("hello" + claims["username"].(string) + claims["id"].(string)))
		// id, _ := strconv.Atoi(claims["id"].(string))
		return true, 1
	}
	fmt.Println([]byte("not authorized"))
	fmt.Println(err)
	return false, -1
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		fmt.Println("OPTIONS LOGOUT")
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodDelete {

		fmt.Println("LOGOUT")

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
		fmt.Println("successful logout")
	}
}

func (server *Server) IsLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("->>>>>>>>>>>>>>>> Is login")

	myCookie, _ := r.Cookie("session_id")
	fmt.Println("в IsLogin пришло cookie:", myCookie)

	if r.Method == http.MethodOptions {
		fmt.Println("Is Login in options")
		w.WriteHeader(http.StatusOK)
		return
	}

	if value, _ := server.CheckLogin(w, r); !value {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("{}"))
		fmt.Println("StatusNoContent")
		return
	}
	fmt.Println("StatusOK")
	w.WriteHeader(http.StatusOK)
}

func (server *Server) LoginAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginAPI")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {

		var user models.UserAuth
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

		test, _ := r.Cookie("session_id")
		fmt.Println("пришло cookie:", test)

		if token, _, err := server.db.Login(user.Username, user.Password, []byte(server.serverConfig.Secret)); err != nil {
			fmt.Println("Login error")
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			cookie := &http.Cookie{
				Name:  "session_id",
				Value: token,

			}

			fmt.Println("должно внутри быть token =", token)

			cookie.Path = "/"
			cookie.Expires = time.Now().Add(5 * time.Hour)
			cookie.HttpOnly = false
			cookie.Secure = false
			http.SetCookie(w, cookie)
			w.Write([]byte(token))
			w.WriteHeader(http.StatusOK)

			fmt.Println("successful authorization")
		}
	}
}

func (server *Server) SignUpAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignUpAPI")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {

		if value, _ := server.CheckLogin(w, r); value {
			w.WriteHeader(http.StatusOK)
			fmt.Println("successful authorization")
		}

		jsonStr := server.CreateUser(w, r)

		var user models.User
		fmt.Println("json: ", jsonStr)
		err := json.Unmarshal(jsonStr, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if token, _, err := server.db.Login(user.Username, user.Password, []byte(server.serverConfig.Secret)); err != nil {
			fmt.Println("Login error")
			fmt.Println(err)
		} else {
			cookie := &http.Cookie{
				Name:  "session_id",
				Value: token,
			}
			// id_cookie := &http.Cookie{
			// 	Name:  "user_id",
			// 	Value: userId,
			// }
			cookie.Path = "/"
			cookie.Expires = time.Now().Add(5 * time.Hour)
			cookie.HttpOnly = false
			cookie.Secure = false
			http.SetCookie(w, cookie)
			// http.SetCookie(w, id_cookie)

			w.Write([]byte(token))
			fmt.Println("successful authorization")
		}
	}
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if value, user_id := server.CheckLogin(w, r); value {
		fmt.Println("VALUE IN ID", user_id)
		result, find, err := server.db.GetUserByID(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !find {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println(result)
		WriteToResponse(w, http.StatusOK, result)
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (server *Server) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	function := func(header multipart.FileHeader) error {
		re := regexp.MustCompile(`image/.*`)
		if !re.MatchString(header.Header.Get("Content-Type")) {
			fmt.Println(header.Header.Get("Content-Type"))
			return fmt.Errorf("not an image")
		}
		return nil
	}
	_, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pathToAvatar, err := filesystem.UploadFile(w, r, function, 
		server.serverConfig.UploadPath, server.serverConfig.AvatarsPath)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = server.db.AddImage(pathToAvatar, userID)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) []byte {
	var user models.User
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return jsonStr
	}
	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return jsonStr
	}
	if br, err_r := server.db.UserRegistration(user.Username, user.Email, user.Password, user.LangID, user.PronounceON); br != true {
		fmt.Println(err_r.Error())
		w.WriteHeader(http.StatusConflict)
		return jsonStr
	}
	server.db.LastId++
	return jsonStr
}

func (server *Server) UsersPaginate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("debug")
	pages, ok := r.URL.Query()["page"]
	if !ok || len(pages[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rows, ok := r.URL.Query()["rows"]
	if !ok || len(rows[0]) < 1 {
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
	result, err := server.db.GetUsers(page, rowsNum)
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
		_, find, err := server.db.GetUserByID(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !find {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		server.db.UpdateUserById(user_id, user.Username, user.Email, user.Password, user.LangID, user.PronounceON)
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if value, user_id := server.CheckLogin(w, r); value {
		_, find, err := server.db.GetUserByID(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !find {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		isDelete, _ := server.db.DeleteUserById(user_id)
		if !isDelete {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
