package monitor

import (
    "fmt"
    "time"
)

// Metrics holds all hardware metrics
type Metrics struct {
    Timestamp   time.Time
    CPU         CPUMetrics
    Memory      MemoryMetrics
    Disk        []DiskMetrics
    Battery     BatteryMetrics
    Network     []NetworkMetrics
    Temperature TemperatureMetrics
}

// Monitor is the main hardware monitor
type Monitor struct {
    cpuMonitor     *CPUMonitor
    memoryMonitor  *MemoryMonitor
    diskMonitor    *DiskMonitor
    batteryMonitor *BatteryMonitor
    networkMonitor *NetworkMonitor
    tempMonitor    *TemperatureMonitor
}

// New creates a new hardware monitor instance
func New() (*Monitor, error) {
    return &Monitor{
        cpuMonitor:     NewCPUMonitor(),
        memoryMonitor:  NewMemoryMonitor(),
        diskMonitor:    NewDiskMonitor(),
        batteryMonitor: NewBatteryMonitor(),
        networkMonitor: NewNetworkMonitor(),
        tempMonitor:    NewTemperatureMonitor(),
    }, nil
}

// GetAllMetrics fetches all hardware metrics
func (m *Monitor) GetAllMetrics() (*Metrics, error) {
    metrics := &Metrics{
        Timestamp: time.Now(),
    }

    // Get CPU metrics
    cpuMetrics, err := m.cpuMonitor.GetMetrics()
    if err != nil {
        return nil, fmt.Errorf("failed to get CPU metrics: %w", err)
    }
    metrics.CPU = *cpuMetrics

    // Get Memory metrics
    memMetrics, err := m.memoryMonitor.GetMetrics()
    if err != nil {
        return nil, fmt.Errorf("failed to get memory metrics: %w", err)
    }
    metrics.Memory = *memMetrics

    // Get Disk metrics
    diskMetrics, err := m.diskMonitor.GetMetrics()
    if err != nil {
        return nil, fmt.Errorf("failed to get disk metrics: %w", err)
    }
    metrics.Disk = diskMetrics

    // Get Battery metrics
    batteryMetrics, err := m.batteryMonitor.GetMetrics()
    if err == nil { // Battery might not be available on all systems
        metrics.Battery = *batteryMetrics
    }

    // Get Network metrics
    netMetrics, err := m.networkMonitor.GetMetrics()
    if err != nil {
        return nil, fmt.Errorf("failed to get network metrics: %w", err)
    }
    metrics.Network = netMetrics

    // Get Temperature metrics
    tempMetrics, err := m.tempMonitor.GetMetrics()
    if err == nil { // Temperature might not be available on all systems
        metrics.Temperature = *tempMetrics
    }

    return metrics, nil
}