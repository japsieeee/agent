package main

import (
	"log"
	"time"
	"fmt"

	"github.com/japsieeee/agent/internal/collector"
	"github.com/japsieeee/agent/internal/sender"
)

func main() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			metrics, err := collector.Collect()
			if err != nil {
				log.Println("collect error:", err)
				continue
			}

			fmt.Printf("Time: %v\n", metrics.Timestamp)
			fmt.Printf("CPU: %.2f%%\n", metrics.CPUPercent)
			fmt.Printf("Memory: %d / %d\n", metrics.MemoryUsed, metrics.MemoryTotal)
			fmt.Printf("Disk: %d / %d\n", metrics.DiskUsed, metrics.DiskTotal)
			fmt.Printf("Network Sent: %d\n", metrics.NetworkSent)
			fmt.Printf("Network Recv: %d\n", metrics.NetworkRecv)
			fmt.Println("---------------")

			err = sender.Send(metrics)
			if err != nil {
				log.Println("send error:", err)
			}
		}
	}
}