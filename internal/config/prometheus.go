package config

import "github.com/prometheus/client_golang/prometheus"

var ResponseTimeHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Histogram of response time for HTTP requests.",
	},
	[]string{"method", "path"},
)
