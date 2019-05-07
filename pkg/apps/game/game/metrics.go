package game

import "github.com/prometheus/client_golang/prometheus"

var (
	RoomCountMetric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "rooms_count",
			Help: "Rooms online",
		},
	)
)