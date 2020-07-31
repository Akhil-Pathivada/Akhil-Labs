package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_pushgateway_example_last_exection_timestamp",
		Help: "The timestamp of the last successful execution of Go program.",
	})

	completionTime.SetToCurrentTime()

	if err := push.New("http://localhost:9091", "push_gateway_sample_job_1").
		Collector(completionTime).
		Grouping("db", "customers").
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
