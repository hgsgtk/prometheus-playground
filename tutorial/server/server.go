package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "Number of request handled by Ping Handler",
	},
)

var opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
	Name: "myapp_processed_ops_total",
	Help: "The total number of processed events",
})

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func main() {
	prometheus.MustRegister(pingCounter)

	recordMetrics()

	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8090", nil)
}
