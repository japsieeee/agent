package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/japsieeee/agent/internal/collector"
	"github.com/japsieeee/agent/internal/config"
	"github.com/japsieeee/agent/internal/sender"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Create sender with injected config
	s := sender.NewSender(cfg)

	// Use interval from config
	interval := time.Duration(cfg.Metrics.IntervalSeconds) * time.Second
	if interval <= 0 {
		interval = 30 * time.Second
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Graceful shutdown handling
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Println("agent started")

	for {
		select {
		case <-ctx.Done():
			log.Println("shutting down agent...")
			return

		case <-ticker.C:
			runCycle(s)
		}
	}
}

func runCycle(s *sender.Sender) {
	metrics, err := collector.Collect()
	if err != nil {
		log.Println("collect error:", err)
		return
	}

	log.Printf(
		"CPU: %.2f%% | Mem: %d/%d | Disk: %d/%d | Net: %d/%d",
		metrics.CPUPercent,
		metrics.MemoryUsed,
		metrics.MemoryTotal,
		metrics.DiskUsed,
		metrics.DiskTotal,
		metrics.NetworkSent,
		metrics.NetworkRecv,
	)

	if err := s.Send(metrics); err != nil {
		log.Println("send error:", err)
	}
}