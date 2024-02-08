package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/simpledotorg/alphasms_exporter/alphasms"
)

func main() {
	client := alphasms.Client{APIKey: "yourapikey"}

	exporter := alphasms.NewExporter(&client)
	prometheus.MustRegister(exporter)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
