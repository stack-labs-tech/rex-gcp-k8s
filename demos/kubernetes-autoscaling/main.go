package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := flag.Int64("port", 8080, "port to expose metrics on")
	flag.Parse()

	router := http.NewServeMux()
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		w.Write([]byte(`OK`))
	})
	router.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
	h := promhttp.InstrumentMetricHandler(prometheus.DefaultRegisterer, router)

	http.Handle("/", h)
	log.Printf("starting to listen on :%d", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("failed to start serving metrics: %v", err)
	}
}
