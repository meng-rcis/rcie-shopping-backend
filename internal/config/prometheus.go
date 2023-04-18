package config

import "github.com/prometheus/client_golang/prometheus"

// Get TPS & Average Response Time
var HttpDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_response_time_seconds",
		Help:    "Histogram of response time for HTTP requests.",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"method", "path"},
)

// Get Response Status
var HttpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_response_time_seconds_count",
		Help: "Total number of HTTP requests",
	},
	[]string{"status", "method", "path"},
)
