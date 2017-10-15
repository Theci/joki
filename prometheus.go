package main

import (
	"github.com/prometheus/client_golang/prometheus"
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
