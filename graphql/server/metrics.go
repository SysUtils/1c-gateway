package server

import (
	"github.com/prometheus/client_golang/prometheus"
)

func NewHttpMetric() *prometheus.HistogramVec {
	metrics := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "The duration of HTTP requests processed by an ASP.NET Core application.", // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.ExponentialBuckets(0.001, 2.0, 16),                             // 5 buckets, each 5 centigrade wide.
	}, []string{"code", "resolver", "type"})
	prometheus.MustRegister(metrics)
	return metrics
}
