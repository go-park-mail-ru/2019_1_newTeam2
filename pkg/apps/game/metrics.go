package game

import "github.com/prometheus/client_golang/prometheus"

var (
	MultiplayerHitsMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "Multiplayer",
			Help: "The total number of requests on multiplayer",
		},
		[]string{"code", "method"},
	)
)
