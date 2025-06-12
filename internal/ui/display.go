package ui

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strings"
    "time"

    "github.com/fatih/color"
    "github.com/yourusername/hardware-monitor/internal/monitor"
)

// Display handles the CLI output
type Display struct {
    colorEnabled bool
    titleColor   *color.Color
    headerColor  *color.Color
    normalColor  *color.Color
    warningColor *color.Color
    criticalColor *color.Color
    goodColor    *color.Color
}

// NewDisplay creates a new display instance
func NewDisplay(colorEnabled bool) *Display {
    return &Display{
        colorEnabled:  colorEnabled,
        titleColor:    color.New(color.FgCyan, color.Bold),
        headerColor:   color.New(color.FgYellow, color.Bold),
        normalColor:   color.New(color.FgWhite),
        warningColor:  color.New(color.FgYellow),
        criticalColor: color.New(color.FgRed, color.Bold),
        goodColor:     color.New(color.FgGreen),
    }
}

// Render displays all metrics
func (d *Display) Render(metrics *monitor.Metrics) {
    d.clearScreen()
    
    // Title
    d.printTitle("System Monitor", metrics.Timestamp)
    fmt.Println()

    // CPU Section
    d.printCPU(&metrics.CPU)
    fmt.Println()

    // Memory Section
    d.printMemory(&metrics.Memory)
    fmt.Println()

    // Disk Section
    d.printDisks(metrics.Disk)
    fmt.Println()

    // Network Section
    d.printNetwork(metrics.Network)
    fmt.Println()

    // Battery Section
    if metrics.Battery.IsAvailable {
        d.printBattery(&metrics.Battery)
        fmt.Println()
    }

    // Temperature Section
    if metrics.Temperature.IsAvailable {
        d.printTemperature(&metrics.Temperature)
        fmt.Println()
    }
}

// clearScreen clears the terminal screen
func (d *Display) clearScreen() {
    switch runtime.GOOS {
    case "windows":
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    default:
        fmt.Print("\033[H\033[2J")
    }
}

// printTitle prints the main title
func (d *Display) printTitle(title string, timestamp time.Time) {
    d.titleColor.Printf("=== %s === %s ===\n", title, timestamp.Format("15:04:05"))
}

// printCPU prints CPU metrics
func (d *Display) printCPU(cpu *monitor.CPUMetrics) {
    d.headerColor.Println("CPU Usage:")
    
    // Total CPU usage with color coding
    totalColor := d.getColorForPercent(cpu.TotalUsage)
    fmt.Print("  Total: ")
    totalColor.Printf("%.1f%%", cpu.TotalUsage)
    fmt.Print(" ")
    d.printBar(cpu.TotalUsage, 30)
    fmt.Println()

    // Per-core usage
    for i, usage := range cpu.UsagePerCore {
        coreColor := d.getColorForPercent(usage)
        fmt.Printf("  Core %d: ", i)
        coreColor.Printf("%5.1f%%", usage)
        fmt.Print(" ")
        d.printBar(usage, 25)
        fmt.Println()
    }
}

// printMemory prints memory metrics
func (d *Display) printMemory(mem *monitor.MemoryMetrics) {
    d.headerColor.Println("Memory Usage:")
    
    memColor := d.getColorForPercent(mem.UsedPercent)
    fmt.Print("  RAM: ")
    memColor.Printf("%s / %s (%.1f%%)", 
        d.formatBytes(mem.Used), 
        d.formatBytes(mem.Total), 
        mem.UsedPercent)
    fmt.Print(" ")
    d.printBar(mem.UsedPercent, 30)
    fmt.Println()
}

// printDisks prints disk metrics
func (d *Display) printDisks(disks []monitor.DiskMetrics) {
    d.headerColor.Println("Disk Usage:")
    
    for _, disk := range disks {
        diskColor := d.getColorForPercent(disk.UsedPercent)
        fmt.Printf("  %s: ", disk.MountPoint)
        diskColor.Printf("%s / %s (%.1f%%)", 
            d.formatBytes(disk.Used), 
            d.formatBytes(disk.Total), 
            disk.UsedPercent)
        fmt.Print(" ")
        d.printBar(disk.UsedPercent, 20)
        fmt.Println()
    }
}

// printNetwork prints network metrics
func (d *Display) printNetwork(networks []monitor.NetworkMetrics) {
    d.headerColor.Println("Network Usage:")
    
    for _, net := range networks {
        fmt.Printf("  %s:\n", net.Interface)
        d.normalColor.Printf("    ↓ %s/s", d.formatBytes(uint64(net.BytesRecvRate)))
        fmt.Print("  ")
        d.normalColor.Printf("↑ %s/s", d.formatBytes(uint64(net.BytesSentRate)))
        fmt.Printf("  (Total: ↓ %s ↑ %s)\n", 
            d.formatBytes(net.BytesRecv), 
            d.formatBytes(net.BytesSent))
    }
}

// printBattery prints battery metrics
func (d *Display) printBattery(battery *monitor.BatteryMetrics) {
    d.headerColor.Println("Battery:")
    
    batteryColor := d.goodColor
    if battery.ChargePercent < 20 {
        batteryColor = d.criticalColor
    } else if battery.ChargePercent < 50 {
        batteryColor = d.warningColor
    }
    
    status := "Discharging"
    if battery.IsCharging {
        status = "Charging"
    }
    
    fmt.Print("  ")
    batteryColor.Printf("%.1f%% (%s)", battery.ChargePercent, status)
    fmt.Print(" ")
    d.printBar(battery.ChargePercent, 30)
    fmt.Println()
}

// printTemperature prints temperature metrics
func (d *Display) printTemperature(temp *monitor.TemperatureMetrics) {
    d.headerColor.Println("Temperature:")
    
    tempColor := d.normalColor
    if temp.CPUTemp > 80 {
        tempColor = d.criticalColor
    } else if temp.CPUTemp > 70 {
        tempColor = d.warningColor
    } else if temp.CPUTemp < 50 {
        tempColor = d.goodColor
    }
    
    fmt.Print("  CPU: ")
    tempColor.Printf("%.1f°C", temp.CPUTemp)
    fmt.Println()
}

// printBar prints a progress bar
func (d *Display) printBar(percent float64, width int) {
    filled := int(percent * float64(width) / 100)
    empty := width - filled
    
    barColor := d.getColorForPercent(percent)
    
    fmt.Print("[")
    if d.colorEnabled {
        barColor.Print(strings.Repeat("█", filled))
    } else {
        fmt.Print(strings.Repeat("=", filled))
    }
    fmt.Print(strings.Repeat(" ", empty))
    fmt.Print("]")
}

// getColorForPercent returns appropriate color based on percentage
func (d *Display) getColorForPercent(percent float64) *color.Color {
    if !d.colorEnabled {
        return d.normalColor
    }
    
    if percent >= 90 {
        return d.criticalColor
    } else if percent >= 75 {
        return d.warningColor
    } else if percent < 50 {
        return d.goodColor
    }
    return d.normalColor
}

// formatBytes formats bytes to human-readable format
func (d *Display) formatBytes(bytes uint64) string {
    const unit = 1024
    if bytes < unit {
        return fmt.Sprintf("%d B", bytes)
    }
    
    div, exp := uint64(unit), 0
    for n := bytes / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    
    return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}