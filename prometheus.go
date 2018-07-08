package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

var (
	latency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "joki_latency",
			Help: "Latency measured by Joki",
		},
		[]string{"type", "probe", "probeset", "src_host", "target_host", "target_name"},
	)
	losspct = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "joki_losspct",
			Help: "Packet loss measured by joki",
		},
		[]string{"probe", "probeset", "src_host", "target_host", "target_name"},
	)
)

func init() {
	prometheus.MustRegister(latency)
	prometheus.MustRegister(losspct)
}

func HandleProm(points []pingresult) {
	hostname, _ := os.Hostname()
	for _, point := range points {
		// Prometheus metric handling
		latency.WithLabelValues("avg", point.probename, point.probeset, hostname, point.host, point.target).Set(point.avg)
		latency.WithLabelValues("max", point.probename, point.probeset, hostname, point.host, point.target).Set(point.max)
		latency.WithLabelValues("min", point.probename, point.probeset, hostname, point.host, point.target).Set(point.min)
		losspct.WithLabelValues(point.probename, point.probeset, hostname, point.host, point.target).Set(float64(point.losspct))
	}
}

func PromExporter() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":65080", nil))
}
