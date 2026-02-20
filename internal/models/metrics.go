package models

import "time"

type Metrics struct {
	Timestamp     time.Time
	CPUPercent    float64
	MemoryUsed    uint64
	MemoryTotal   uint64
	DiskUsed      uint64
	DiskTotal     uint64
	NetworkSent   uint64
	NetworkRecv   uint64
}