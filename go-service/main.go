package main

import (
	"fmt"
	"go-service/client"
	"go-service/metrics"
	"time"
)

type Metrics struct {
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
	Disk   float64 `json:"disk"`
}

func main() {

	for {
		m := metrics.Collect()
		fmt.Println(m)
		client.SendMetricstoAI(m)
		time.Sleep(10 * time.Second)
	}
}
