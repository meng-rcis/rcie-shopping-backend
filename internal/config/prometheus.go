package config

import "github.com/prometheus/client_golang/prometheus"

// Get TPS & Average Response Time
var HttpDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "HTTP response time in seconds",
	},
	[]string{"method", "path"},
)

// Get Response Status
var HttpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"status", "method", "path"},
)
