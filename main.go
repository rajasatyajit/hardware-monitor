package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/yourusername/hardware-monitor/internal/monitor"
    "github.com/yourusername/hardware-monitor/internal/ui"
)

var (
    refreshInterval = flag.Duration("interval", 1*time.Second, "Refresh interval")
    colorMode      = flag.Bool("color", true, "Enable colored output")
)

func main() {
    flag.Parse()

    // Initialize the hardware monitor
    hwMonitor, err := monitor.New()
    if err != nil {
        log.Fatalf("Failed to initialize hardware monitor: %v", err)
    }

    // Initialize the display
    display := ui.NewDisplay(*colorMode)
    
    // Set up signal handling for graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

    // Create a ticker for refreshing the display
    ticker := time.NewTicker(*refreshInterval)
    defer ticker.Stop()

    // Initial display
    if err := displayMetrics(hwMonitor, display); err != nil {
        log.Printf("Error displaying metrics: %v", err)
    }

    // Main loop
    for {
        select {
        case <-ticker.C:
            if err := displayMetrics(hwMonitor, display); err != nil {
                log.Printf("Error displaying metrics: %v", err)
            }
        case <-sigChan:
            fmt.Println("\n\nShutting down...")
            return
        }
    }
}

func displayMetrics(hwMonitor *monitor.Monitor, display *ui.Display) error {
    metrics, err := hwMonitor.GetAllMetrics()
    if err != nil {
        return fmt.Errorf("failed to get metrics: %w", err)
    }

    display.Render(metrics)
    return nil
}