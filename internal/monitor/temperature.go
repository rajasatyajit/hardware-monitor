package monitor

import (
    "github.com/shirou/gopsutil/v3/host"
)

// TemperatureMetrics holds temperature-related metrics
type TemperatureMetrics struct {
    CPUTemp     float64
    IsAvailable bool
}

// TemperatureMonitor monitors system temperatures
type TemperatureMonitor struct{}

// NewTemperatureMonitor creates a new temperature monitor
func NewTemperatureMonitor() *TemperatureMonitor {
    return &TemperatureMonitor{}
}

// GetMetrics returns current temperature metrics
func (t *TemperatureMonitor) GetMetrics() (*TemperatureMetrics, error) {
    temps, err := host.SensorsTemperatures()
    if err != nil {
        return &TemperatureMetrics{IsAvailable: false}, nil
    }

    // Look for CPU temperature
    for _, temp := range temps {
        // Different systems report CPU temperature with different keys
        if temp.SensorKey == "coretemp" || temp.SensorKey == "cpu" || 
           temp.SensorKey == "CPU" || temp.SensorKey == "Package id 0" {
            return &TemperatureMetrics{
                CPUTemp:     temp.Temperature,
                IsAvailable: true,
            }, nil
        }
    }

    // If no specific CPU temp found, try to get average of all temps
    if len(temps) > 0 {
        sum := 0.0
        for _, temp := range temps {
            sum += temp.Temperature
        }
        return &TemperatureMetrics{
            CPUTemp:     sum / float64(len(temps)),
            IsAvailable: true,
        }, nil
    }

    return &TemperatureMetrics{IsAvailable: false}, nil
}