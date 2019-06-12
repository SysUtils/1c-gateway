package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

func NewHttpMetric() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_metrics",
		Help:    "http requests metrics",            // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(0, 5, 20), // 20 buckets, each 5 centigrade wide.
	}, []string{"code", "resolver", "type"})
}

func MetricCollectMiddleware(metrics *prometheus.HistogramVec, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startRequest := time.Now()
		f(w, r)
		end := time.Now()
		duration := end.Sub(startRequest)
		metrics.With(prometheus.Labels{"code": "200", "type": "query", "resolver": "ProductInManufacturingDocuments"}).Observe(duration.Seconds())
	}
}
