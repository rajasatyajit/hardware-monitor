package monitor

import (
    "github.com/shirou/gopsutil/v3/mem"
)

// MemoryMetrics holds memory-related metrics
type MemoryMetrics struct {
    Total       uint64
    Used        uint64
    Free        uint64
    UsedPercent float64
}

// MemoryMonitor monitors memory usage
type MemoryMonitor struct{}

// NewMemoryMonitor creates a new memory monitor
func NewMemoryMonitor() *MemoryMonitor {
    return &MemoryMonitor{}
}

// GetMetrics returns current memory metrics
func (m *MemoryMonitor) GetMetrics() (*MemoryMetrics, error) {
    vmStat, err := mem.VirtualMemory()
    if err != nil {
        return nil, err
    }

    return &MemoryMetrics{
        Total:       vmStat.Total,
        Used:        vmStat.Used,
        Free:        vmStat.Free,
        UsedPercent: vmStat.UsedPercent,
    }, nil
}