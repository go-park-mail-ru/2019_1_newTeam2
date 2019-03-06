package main

import (
	"fmt"
	"net/http"
	"strconv"

	. "./responses"
)

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = TypeRequest(r.URL.Path)
	userID, err := strconv.Atoi(head)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, find, err := server.Users.GetUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !find {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Response(w, http.StatusOK, result)
}

// var loginFormTmpl = []byte(`
// <html>
// 	<body>
// 	<form action="/" method="post">
// 		Login: <input type="text" name="login">
// 		Password: <input type="password" name="password">
// 		<input type="submit" value="Login">
// 	</form>
// 	</body>
// </html>
// `)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "you enter: ", r.FormValue("login"), r.FormValue("password"))

}
