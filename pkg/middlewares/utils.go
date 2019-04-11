package middlewares

import (
	"net/http"
)

type logResWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLogResWriter(w http.ResponseWriter) *logResWriter {
	return &logResWriter{w, http.StatusOK}
}

func (w *logResWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
