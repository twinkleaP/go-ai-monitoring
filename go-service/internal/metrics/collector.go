package metrics

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type Metric struct {
	CPU float64 `json:"cpu"`
	RAM float64 `json:"ram"`
}

func Collect() Metric {
	// CPU usage
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Printf("error collecting CPU: %v", err)
	}

	// RAM usage
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("error collecting RAM: %v", err)
	}

	return Metric{
		CPU: percent[0],         // first CPU core usage
		RAM: vmStat.UsedPercent, // RAM usage in %
	}
}
