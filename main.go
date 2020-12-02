package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	httpRequest = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of http request",
	})
)

func serveRequest() {
	http.HandleFunc("/apps", func(w http.ResponseWriter, r *http.Request) {
		httpRequest.Inc()
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":8080", nil)
}

func main() {
	go serveRequest()
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}