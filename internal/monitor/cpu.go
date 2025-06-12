package monitor

import (
    "github.com/shirou/gopsutil/v3/cpu"
)

// CPUMetrics holds CPU-related metrics
type CPUMetrics struct {
    UsagePerCore []float64
    TotalUsage   float64
    CoreCount    int
}

// CPUMonitor monitors CPU usage
type CPUMonitor struct{}

// NewCPUMonitor creates a new CPU monitor
func NewCPUMonitor() *CPUMonitor {
    return &CPUMonitor{}
}

// GetMetrics returns current CPU metrics
func (c *CPUMonitor) GetMetrics() (*CPUMetrics, error) {
    // Get per-core CPU usage
    perCore, err := cpu.Percent(0, true)
    if err != nil {
        return nil, err
    }

    // Get total CPU usage
    total, err := cpu.Percent(0, false)
    if err != nil {
        return nil, err
    }

    totalUsage := float64(0)
    if len(total) > 0 {
        totalUsage = total[0]
    }

    return &CPUMetrics{
        UsagePerCore: perCore,
        TotalUsage:   totalUsage,
        CoreCount:    len(perCore),
    }, nil
}