package main

import (
	"fmt"
	"go-service/client"
	"go-service/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
	Disk   float64 `json:"disk"`
}

var (
	cpuGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_cpu_percent",
		Help: "CPU usage percent",
	})
	memoryGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_memory_percent",
		Help: "Memory usage percent",
	})
	diskGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_disk_percent",
		Help: "Disk usage percent",
	})
	anomalyCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "anomalies_detected_total",
		Help: "Total anomalies detected",
	})
)

func init() {
	prometheus.MustRegister(cpuGauge, memoryGauge, diskGauge, anomalyCounter)
}

func main() {

	for {
		m := metrics.Collect()
		fmt.Println(m)
		client.SendMetricstoAI(m)
		time.Sleep(10 * time.Second)
	}
}
