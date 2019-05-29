package responses

import (
	"github.com/mailru/easyjson"
	"net/http"
)

func WriteToResponse(w http.ResponseWriter, status int, v easyjson.Marshaler) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	response, _ := easyjson.Marshal(v)
	_, _ = w.Write(response)
}
