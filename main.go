package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	message := map[string]string{"message": "pong"}
	json.NewEncoder(w).Encode(message)
}

func main() {
	prometheus.MustRegister(pingCounter)
	port := os.Getenv("PORT")
	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+port, nil)
}
