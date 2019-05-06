package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	ApiMetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "API_requests",
			Help: "The total number of requests on API",
		},
		[]string{"code", "path", "method"},
	)
)

func (server *Server) metricsMiddleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ApiMetrics.With(prometheus.Labels{"code": strconv.Itoa(http.StatusOK), "path": r.URL.Path, "method": r.Method}).Inc()
		})
	}
}