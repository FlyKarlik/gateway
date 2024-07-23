package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sync"
)

// RunPrometheus start prometheus service
func RunPrometheus(wg *sync.WaitGroup, prometheusHost string) {
	defer wg.Done()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(prometheusHost, nil)

	if err != nil {
		log.Fatalf("prometheus failed: %s\n", err)
	}
}
