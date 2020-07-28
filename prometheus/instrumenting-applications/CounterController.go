package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func CountRequests(counter *prometheus.CounterVec) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// counter.Inc()
		code := 200
		defer func() {
			// httpDuration := time.Since(start)
			counter.WithLabelValues(fmt.Sprintf("%d", code)).Inc()
		}()

		code = http.StatusBadRequest
		if r.Method == "GET" {
			code = http.StatusOK
			w.Write([]byte("Hello World"))
		} else {
			w.WriteHeader(code)
		}
	}
}
