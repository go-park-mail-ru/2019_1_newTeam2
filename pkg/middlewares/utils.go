package middlewares

import (
	"bufio"
	"fmt"
	"net"
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

func (w *logResWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {

	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("cannot Highjack")
	}
	return h.Hijack()
}
