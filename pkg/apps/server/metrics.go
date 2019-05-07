package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
)

var (
	ApiMetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "API",
			Help: "The total number of requests on API",
		},
		[]string{"code", "path", "method"},
	)
)

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (srv *Server) metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writter := statusWriter{w, http.StatusOK}
		next.ServeHTTP(&writter, r)
		ApiMetrics.With(prometheus.Labels{"code": strconv.Itoa(writter.statusCode), "path": r.URL.Path, "method": r.Method}).Inc()
	})
}