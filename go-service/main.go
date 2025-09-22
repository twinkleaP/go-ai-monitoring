package main

import (
	"fmt"
	"time"

	"github.com/you/go-service/client"
	"github.com/you/go-service/metrics"
)

func main() {

	for {
		m := metrics.Collect()
		fmt.Println(m)
		client.SendMetricstoAI(m)
		time.Sleep(10 * time.Second)
	}
}
