package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	// Prometheus: Counter to collect required metrics
	requests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of different HTTP requests",
	}, []string{"code"}) // this will be partitioned by the HTTP code.

	// Prometheus: Gauge to collect required metrics
	queue := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "items_in_queue",
		Help: "Count of items in Queue",
	}, []string{"code"}) // this will be partitioned by the HTTP code.

	// Prometheus: Histogram to collect required metrics
	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "greeting_seconds",
		Help: "Time take to greet someone",
		// Buckets: []float64{1, 2, 5, 6, 10}, //defining small buckets as this app should not take more than 1 sec to respond
	}, []string{"code"}) // this will be partitioned by the HTTP code.

	router := mux.NewRouter()

	// Counter
	router.Handle("/hello", countRequests(requests))

	// Gauge
	router.Handle("/push", Push(queue))
	router.Handle("/pop", Pop(queue))

	// Histogram
	router.Handle("/wait", Sayhello(histogram))

	//Metrics endpoint for scrapping
	router.Handle("/metrics", promhttp.Handler())

	//Registering the defined metric with Prometheus
	prometheus.Register(requests)
	prometheus.Register(queue)
	prometheus.Register(histogram)

	log.Fatal(http.ListenAndServe(":8009", router))
}
