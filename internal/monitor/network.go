package monitor

import (
    "github.com/shirou/gopsutil/v3/net"
    "time"
)

// NetworkMetrics holds network-related metrics for a single interface
type NetworkMetrics struct {
    Interface   string
    BytesSent   uint64
    BytesRecv   uint64
    PacketsSent uint64
    PacketsRecv uint64
    BytesSentRate   float64
    BytesRecvRate   float64
}

// NetworkMonitor monitors network usage
type NetworkMonitor struct {
    lastStats map[string]net.IOCountersStat
    lastTime  time.Time
}

// NewNetworkMonitor creates a new network monitor
func NewNetworkMonitor() *NetworkMonitor {
    return &NetworkMonitor{
        lastStats: make(map[string]net.IOCountersStat),
        lastTime:  time.Now(),
    }
}

// GetMetrics returns current network metrics for all interfaces
func (n *NetworkMonitor) GetMetrics() ([]NetworkMetrics, error) {
    stats, err := net.IOCounters(true)
    if err != nil {
        return nil, err
    }

    currentTime := time.Now()
    timeDiff := currentTime.Sub(n.lastTime).Seconds()

    var metrics []NetworkMetrics
    for _, stat := range stats {
        // Skip loopback interface
        if stat.Name == "lo" || stat.Name == "lo0" {
            continue
        }

        metric := NetworkMetrics{
            Interface:   stat.Name,
            BytesSent:   stat.BytesSent,
            BytesRecv:   stat.BytesRecv,
            PacketsSent: stat.PacketsSent,
            PacketsRecv: stat.PacketsRecv,
        }

        // Calculate rates if we have previous data
        if lastStat, exists := n.lastStats[stat.Name]; exists && timeDiff > 0 {
            metric.BytesSentRate = float64(stat.BytesSent-lastStat.BytesSent) / timeDiff
            metric.BytesRecvRate = float64(stat.BytesRecv-lastStat.BytesRecv) / timeDiff
        }

        metrics = append(metrics, metric)
        n.lastStats[stat.Name] = stat
    }

    n.lastTime = currentTime
    return metrics, nil
}