package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (server *Server) getUser(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = typeRequest(r.URL.Path)
	userID, err := strconv.Atoi(head)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, find, err := server.users.getUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !find {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response(w, http.StatusOK, result)
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

func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "you enter: ", r.FormValue("login"), r.FormValue("password"))
}
