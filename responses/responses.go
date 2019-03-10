package responses

import (
	"encoding/json"
	"net/http"

	"github.com/user/2019_1_newTeam2/requests"
	"github.com/user/2019_1_newTeam2/storage"
)

func UserResponse(w http.ResponseWriter, status int, result storage.User) {
	result.Password = ""
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	response, _ := json.Marshal(result)
	w.Write(response)
}

func UserTableResponse(w http.ResponseWriter, status int, result []storage.User) {
	table := make([]requests.UserTableElem, 0)
	for _, i := range result {
		table = append(table, requests.UserTableElem{i.Username, i.Score})
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	response, _ := json.Marshal(table)
	w.Write(response)
}
