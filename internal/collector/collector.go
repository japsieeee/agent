package collector

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"

	"github.com/japsieeee/agent/internal/models"
)

func Collect() (*models.Metrics, error) {
	now := time.Now()

	// CPU (1 second sample for accurate utilization)
	cpuPercents, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	// Memory
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	// Disk (root partition)
	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	// Network (aggregate all interfaces)
	netStats, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}

	metrics := &models.Metrics{
		Timestamp:   now,
		CPUPercent:  cpuPercents[0],
		MemoryUsed:  vmStat.Used,
		MemoryTotal: vmStat.Total,
		DiskUsed:    diskStat.Used,
		DiskTotal:   diskStat.Total,
		NetworkSent: netStats[0].BytesSent,
		NetworkRecv: netStats[0].BytesRecv,
	}

	return metrics, nil
}