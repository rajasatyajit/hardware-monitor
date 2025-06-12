
package monitor

import (
    "github.com/shirou/gopsutil/v3/disk"
)

// DiskMetrics holds disk-related metrics for a single disk
type DiskMetrics struct {
    MountPoint  string
    Total       uint64
    Used        uint64
    Free        uint64
    UsedPercent float64
}

// DiskMonitor monitors disk usage
type DiskMonitor struct{}

// NewDiskMonitor creates a new disk monitor
func NewDiskMonitor() *DiskMonitor {
    return &DiskMonitor{}
}

// GetMetrics returns current disk metrics for all mounted disks
func (d *DiskMonitor) GetMetrics() ([]DiskMetrics, error) {
    partitions, err := disk.Partitions(false)
    if err != nil {
        return nil, err
    }

    var metrics []DiskMetrics
    for _, partition := range partitions {
        usage, err := disk.Usage(partition.Mountpoint)
        if err != nil {
            continue // Skip this partition if we can't get usage
        }

        metrics = append(metrics, DiskMetrics{
            MountPoint:  partition.Mountpoint,
            Total:       usage.Total,
            Used:        usage.Used,
            Free:        usage.Free,
            UsedPercent: usage.UsedPercent,
        })
    }

    return metrics, nil
}