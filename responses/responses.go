package responses

import (
	"encoding/json"
	"net/http"

	. "../requests"
	. "../storage"
)

func UserResponse(w http.ResponseWriter, status int, result User) {
	result.Password = ""
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	response, _ := json.Marshal(result)
	w.Write(response)
}

func UserTableResponse(w http.ResponseWriter, status int, result []User) {
	table := make([]UserTableElem, 0)
	for _, i := range result {
		table = append(table, UserTableElem{i.Username, i.Score})
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	response, _ := json.Marshal(table)
	w.Write(response)
}
