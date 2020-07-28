package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func Push(queue *prometheus.GaugeVec) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// counter.Inc()
		code := 200
		defer func() {
			// httpDuration := time.Since(start)
			queue.WithLabelValues(fmt.Sprintf("%d", code)).Inc()
		}()

		code = http.StatusBadRequest
		if r.Method == "GET" {
			code = http.StatusOK
			w.Write([]byte("Pushed an item to Queue"))
		} else {
			w.WriteHeader(code)
		}
	}
}

func Pop(queue *prometheus.GaugeVec) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		code := 200
		defer func() {
			// httpDuration := time.Since(start)
			queue.WithLabelValues(fmt.Sprintf("%d", code)).Dec()
		}()

		code = http.StatusBadRequest
		if r.Method == "GET" {
			code = http.StatusOK
			w.Write([]byte("Popped an item from Queue"))
		} else {
			w.WriteHeader(code)
		}
	}
}
