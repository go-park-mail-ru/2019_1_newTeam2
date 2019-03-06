package responses

import (
	"encoding/json"
	"net/http"

	. "../storage"
)

func Response(w http.ResponseWriter, status int, result User) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	response, _ := json.Marshal(result)
	w.Write(response)
}
