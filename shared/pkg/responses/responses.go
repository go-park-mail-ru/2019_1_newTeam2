package responses

import (
	"encoding/json"
	"net/http"
)

func WriteToResponse(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	response, _ := json.Marshal(v)
	_, _ = w.Write(response)
}
