package metrics

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type Metric struct {
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
	Disk   float64 `json:"disk"`
}

// Collect real system metrics
func Collect() Metric {
	// CPU usage
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Println("CPU error:", err)
	}

	// Memory usage
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Memory error:", err)
	}

	// Disk usage (root partition "/")
	diskStat, err := disk.Usage("/")
	if err != nil {
		log.Println("Disk error:", err)
	}

	return Metric{
		CPU:    percent[0],
		Memory: vmStat.UsedPercent,
		Disk:   diskStat.UsedPercent,
	}
}
