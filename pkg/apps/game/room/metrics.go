package room

import "github.com/prometheus/client_golang/prometheus"

var (
	PlayerCountMetric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "players_count",
			Help: "Players online",
		},
	)
)
